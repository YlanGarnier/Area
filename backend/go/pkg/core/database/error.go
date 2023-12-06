package database

import "strings"

type dbError string

const (
	ErrUniqueConstraintFailed dbError = "user already exists: %!w(<nil>)"
)

func DbIsError(err error, dbError dbError) bool {
	return strings.Contains(err.Error(), string(dbError))
}
