package serviceDb

import (
	"main/ExtError"
	"main/db/postgresql"
	"main/interface/interfaceDb"
)

var service dbService

type dbService struct {
	connect string
	db      interfaceDb.DB
}

func StartService(connect string) (extErr *ExtError.Error) {
	service.db = &postgresql.DBPostgresql{}
	service.connect = connect
	return service.db.OpenDB(connect)
}

func GetDb() (db interfaceDb.DB, extErr *ExtError.Error) {
	if service.connect == "" {
		extErr = ExtError.New("Service not connected", 1)
		return
	}
	if !service.db.IsOpen() {
		extErr = ExtError.New("DB not open", 1)
		return
	}
	return service.db, nil
}
