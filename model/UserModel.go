package model

import "database/sql"

type UserModel struct {
	ID        sql.NullInt64
	Username  sql.NullString
	Password  sql.NullString
	FirstName sql.NullString
	LastName  sql.NullString
	Gender    sql.NullString
	Telephone sql.NullString
	Email     sql.NullString
	Address   sql.NullString
	CreatedAt sql.NullString
	UpdatedAt sql.NullString
}
