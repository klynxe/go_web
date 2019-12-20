package pg

var tables []table

type table struct {
	name   string
	create string
}

func init() {
	tables = []table{
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
const t_auth_f_activation_key = "activation_key"

/*const t_session = "t_session"
const t_session_f_id = "id"
const t_session_f_login = "id_user"
const t_auth_f_email = "email"
const t_auth_f_password = "password"
const t_auth_f_status = "status"
const t_auth_f_activation_key = "activation_key"*/

const t_auth_create = `
CREATE TABLE ` + t_auth + `
(
   	` + t_auth_f_id + ` bigserial NOT NULL,
    ` + t_auth_f_login + ` character varying NOT NULL,
	` + t_auth_f_email + ` character varying NOT NULL,
    ` + t_auth_f_password + ` character varying NOT NULL,    
	` + t_auth_f_status + ` integer NOT NULL DEFAULT 0,    
	` + t_auth_f_activation_key + ` character varying NOT NULL,    
    PRIMARY KEY ( ` + t_auth_f_id + `),
	CONSTRAINT u_login UNIQUE (` + t_auth_f_login + `),
	CONSTRAINT u_email UNIQUE ( ` + t_auth_f_email + `)
)
`

/*const t_session_create = `
CREATE TABLE ` + t_auth + `
(
   	` + t_auth_f_id + ` bigserial NOT NULL,
    ` + t_auth_f_login + ` character varying NOT NULL,
	` + t_auth_f_email + ` character varying NOT NULL,
    ` + t_auth_f_password + ` character varying NOT NULL,
	` + t_auth_f_status + ` integer NOT NULL DEFAULT 0,
	` + t_auth_f_activation_key + ` character varying NOT NULL,
    PRIMARY KEY ( ` + t_auth_f_id + `),
	CONSTRAINT u_login UNIQUE (` + t_auth_f_login + `),
	CONSTRAINT u_email UNIQUE ( ` + t_auth_f_email + `)
)
`*/
