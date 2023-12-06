package gormModels

import "gorm.io/gorm"

type Area struct {
	gorm.Model
	Name        string
	ActService  string
	RouteAction string
	ReaService  string
	Route       string
	Base        string

	UserID uint
}
