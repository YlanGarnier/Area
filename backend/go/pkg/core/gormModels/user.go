package gormModels

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	FirstName sql.NullString
	LastName  sql.NullString
	Password  sql.NullString
	Username  sql.NullString `gorm:"unique"`
	Token     sql.NullString `gorm:"unique"`
	Kind      string

	Services []*Service
	Areas    []*Area
}
