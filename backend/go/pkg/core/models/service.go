package models

type Service struct {
	Name         string
	Identifier   string
	AccessToken  string
	RefreshToken string

	UserID uint
}
