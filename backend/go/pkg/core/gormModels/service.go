package gormModels

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Name         string
	Identifier   string
	Token        string
	RefreshToken string

	UserID uint
}
