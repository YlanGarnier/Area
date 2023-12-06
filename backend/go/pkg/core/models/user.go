package models

import "database/sql"

type User struct {
	ID        uint
	Email     string
	FirstName sql.NullString
	LastName  sql.NullString
	Password  sql.NullString
	Username  sql.NullString
	Token     sql.NullString
	Kind      string

	Services []Service
	Areas    []Area
}
