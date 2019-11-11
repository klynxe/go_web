package postgresql

var tables []index

type table struct {
	name   string
	create string
}

func init() {
	indexes = []index{
		{
			t_auth,
			t_auth_create,
		},
	}
}

const t_auth = "t_auth"
const t_auth_f_id = "id"
const t_auth_f_login = "login"
const t_auth_f_email = "email"
const t_auth_f_password = "password"
const t_auth_f_status = "status"

const t_auth_create = `
CREATE TABLE ` + t_auth + `
(
   	` + t_auth_f_id + ` bigserial NOT NULL,
    ` + t_auth_f_login + ` character varying NOT NULL,
	` + t_auth_f_email + ` character varying NOT NULL,
    ` + t_auth_f_password + ` character varying NOT NULL,    
	` + t_auth_f_status + ` integer NOT NULL,    
    PRIMARY KEY ( ` + t_auth_f_id + `),
	CONSTRAINT u_login UNIQUE (` + t_auth_f_login + `)
	CONSTRAINT u_login UNIQUE ( ` + t_auth_f_email + `)
)
`
