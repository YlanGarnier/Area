package models

type Authorization struct {
	ClientID    string
	Scope       string
	State       string
	RedirectUri string
}
