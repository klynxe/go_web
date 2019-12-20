package pg

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	_ "github.com/lib/pq"
	"main/ExtError"
	"main/rand"
	"os"
)

var saltForPassword = os.Getenv("DATABASE_PASSWORD_SALT")

type DBPostgresql struct {
	*sql.DB
}

func (pg *DBPostgresql) IsOpen() (open bool) {
	return pg.DB != nil && pg.DB.Ping() == nil
}

func (pg *DBPostgresql) Open(connect string) (extErr *ExtError.Error) {
	var err error
	if pg.DB, err = sql.Open("postgres", connect); err != nil {
		return ExtError.Resend("Error open DB", 1, err)
	}

	if err = pg.DB.Ping(); err != nil {
		return ExtError.Resend("Error ping DB", 2, err)
	}

	for _, table := range tables {
		tableExist := false
		tableExist, err = pg.isTableExist(table.name)
		if err = pg.DB.Ping(); err != nil {
			return ExtError.Resend("Check exist table", 3, err)
		}
		if !tableExist {
			if err := pg.create(table.create); err != nil {
				return ExtError.Resend("Create table", 4, err)
			}
		}
	}

	for _, index := range indexes {
		indexExist := false
		indexExist, err = pg.isIndexExist(index.name)
		if err = pg.DB.Ping(); err != nil {
			return ExtError.Resend("Check exist index", 3, err)
		}
		if !indexExist {
			if err := pg.create(index.create); err != nil {
				return ExtError.Resend("Create index", 4, err)
			}
		}
	}

	return
}

func (pg *DBPostgresql) isTableExist(tableName string) (exist bool, extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	var count int
	err := pg.QueryRow(
		`SELECT COUNT(*) FROM information_schema.tables
			WHERE table_schema = 'public'
			AND  TABLE_NAME = $1`,
		tableName).
		Scan(&count)
	if err != nil {
		extErr = ExtError.Resend("Error check exist table "+tableName, 0, err)
		return
	}
	exist = count > 0
	return
}

func (pg *DBPostgresql) isIndexExist(indexName string) (exist bool, extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	var count int
	err := pg.QueryRow(
		`SELECT COUNT(*) FROM pg_class
			WHERE relname = $1`,
		indexName).
		Scan(&count)
	if err != nil {
		extErr = ExtError.Resend("Error check exist index "+indexName, 1, err)
		return
	}
	exist = count > 0
	return
}

func (pg *DBPostgresql) create(create string) (extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	_, err := pg.Exec(create)
	if err != nil {
		extErr = ExtError.Resend("Error create "+create, 1, err)
		return
	}
	return
}

func (pg *DBPostgresql) CheckExistLogin(login string) (exist bool, extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	var count int
	err := pg.QueryRow(
		`SELECT COUNT(*) FROM `+t_auth+`
			WHERE `+t_auth_f_login+` = $1`,
		login).
		Scan(&count)
	if err != nil {
		extErr = ExtError.Resend("Error check exist login "+login, 1, err)
		return
	}
	exist = count > 0
	return
}

func (pg *DBPostgresql) checkExistEmail(email string) (exist bool, extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	var count int
	err := pg.QueryRow(
		`SELECT COUNT(*) FROM `+t_auth+`
			WHERE `+t_auth_f_email+` = $1`,
		email).
		Scan(&count)
	if err != nil {
		extErr = ExtError.Resend("Error check exist email "+email, 1, err)
		return
	}
	exist = count > 0
	return
}

func (pg *DBPostgresql) SignUp(login, email, password string) (activationKey string, extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	exist, err := pg.CheckExistLogin(login)
	if err != nil {
		extErr = ExtError.Resend("Error Sing Up ", 1, err)
		return
	}
	if exist {
		extErr = ExtError.New("Error Sing Up, Auth already exist ", 2)
		return
	}

	exist, err = pg.checkExistEmail(email)
	if err != nil {
		extErr = ExtError.Resend("Error Sing Up ", 3, err)
		return
	}
	if exist {
		extErr = ExtError.New("Error Sing Up, Email already exist ", 4)
		return
	}

	md := md5.Sum([]byte(saltForPassword + password))
	hexPass := hex.EncodeToString(md[:])

	tx, dbErr := pg.Begin()
	if dbErr != nil {
		tx.Rollback()
		extErr = ExtError.New("Error Sing Up start transaction ", 5)
		return
	}

	activationKey = rand.String(64)
	mdActivationKey := md5.Sum([]byte(activationKey))
	hashActivationKey := hex.EncodeToString(mdActivationKey[:])

	_, dbErr = tx.Exec(`
	INSERT INTO `+t_auth+`(
		`+t_auth_f_login+`, `+t_auth_f_email+`, `+t_auth_f_password+`, `+t_auth_f_activation_key+`)
		VALUES ($1, $2, $3, $4);
	`,
		login, email, hexPass, hashActivationKey)
	if dbErr != nil {
		tx.Rollback()
		extErr = ExtError.Resend("Error Sing Up Error create auth row "+login+", "+email, 6, dbErr)
		return
	}

	dbErr = tx.Commit()
	if dbErr != nil {
		tx.Rollback()
		extErr = ExtError.Resend("Error Sing Up Error create auth row "+login+", "+email+" (commit)", 7, dbErr)
		return
	}

	return
}

func (pg *DBPostgresql) Login(loginOrEmail, password string) (status int, extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	md := md5.Sum([]byte(saltForPassword + password))
	hexPass := hex.EncodeToString(md[:])
	err := pg.QueryRow(
		`SELECT `+t_auth_f_status+` FROM `+t_auth+`
			WHERE (`+t_auth_f_login+` = $1 OR `+t_auth_f_email+` = $1) AND `+t_auth_f_password+` = $2`,
		loginOrEmail, hexPass).
		Scan(&status)
	if err != nil {
		extErr = ExtError.Resend("Error login "+loginOrEmail, 1, err)
		return
	}
	return
}

func (pg *DBPostgresql) RemoveAccount(email string) (extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}
	_, err := pg.Exec(
		`DELETE FROM `+t_auth+`
			WHERE `+t_auth_f_email+` = $1`,
		email)
	if err != nil {
		extErr = ExtError.Resend("Error delete email: "+email, 1, err)
		return
	}
	return
}

func (pg *DBPostgresql) ConfirmAccount(login, activationKey string, status int) (statusOld int, extErr *ExtError.Error) {
	if pg.DB == nil {
		extErr = ExtError.New("DB is not opening ", 0)
		return
	}

	err := pg.QueryRow(
		`SELECT `+t_auth_f_status+` FROM `+t_auth+`
			WHERE `+t_auth_f_login+` = $1`,
		login).
		Scan(&statusOld)
	if err != nil {
		extErr = ExtError.Resend("Error select activation login "+login, 1, err)
		return
	}

	md := md5.Sum([]byte(activationKey))
	hexActivationKey := hex.EncodeToString(md[:])
	res, err := pg.Exec(
		`UPDATE `+t_auth+` SET `+t_auth_f_status+` = $1 
			WHERE `+t_auth_f_login+` = $2 AND `+t_auth_f_activation_key+` = $3`,
		status, login, hexActivationKey)
	if err != nil {
		extErr = ExtError.Resend("Error activate login: "+login, 1, err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		extErr = ExtError.Resend("Error activate login: "+login+" on get row affected", 1, err)
		return
	}
	if count < 1 {
		extErr = ExtError.New("Activation not valid: "+login+" key: "+activationKey, 1)
		return
	}
	return
}
