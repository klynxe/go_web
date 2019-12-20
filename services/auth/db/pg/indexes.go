package pg

var indexes []index

type index struct {
	name   string
	create string
}

func init() {
	indexes = []index{
		{
			i_auth_login,
			i_auth_login_create,
		},
		{
			i_auth_email,
			i_auth_email_create,
		},
	}
}

const i_auth_login = "i_auth_login"
const i_auth_email = "i_auth_email"

const i_auth_login_create = `
CREATE UNIQUE INDEX ` + i_auth_login + `
    ON t_auth USING btree
    (login ASC NULLS LAST);
`

const i_auth_email_create = `
CREATE UNIQUE INDEX ` + i_auth_email + `
    ON t_auth USING btree
    (login ASC NULLS LAST);
`
