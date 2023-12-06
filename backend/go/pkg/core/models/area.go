package models

import (
	"github.com/lenismtho/area/pkg/protogen"
)

type Area struct {
	ID          uint
	Name        string
	ActService  string
	RouteAction string
	ReaServices string
	Route       string
	Base        *protogen.Base

	UserID uint
}

type ActService string

const (
	HttpService     ActService = "http"
	DiscordService  ActService = "discord"
	EthereumService ActService = "ethereum"
	SpotifyService  ActService = "spotify"
	GmailService    ActService = "google_gmail"
)
