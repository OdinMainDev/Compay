package user

import (
	"context"
	"efw"
)

func ResetSequence() {
	efw.DBTransaction[any]().Exec(context.Background(), `SELECT setval(pg_get_serial_sequence('users', 'id'), (SELECT MAX(id) FROM users))`)
}
