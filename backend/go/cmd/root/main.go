package main

import (
	"github.com/lenismtho/area/cmd/core"
	"github.com/lenismtho/area/cmd/discord"
	"github.com/lenismtho/area/cmd/dropbox"
	"github.com/lenismtho/area/cmd/ethereum"
	"github.com/lenismtho/area/cmd/github"
	"github.com/lenismtho/area/cmd/gmail"
	"github.com/lenismtho/area/cmd/linkedin"
	"github.com/lenismtho/area/cmd/miro"
	"github.com/lenismtho/area/cmd/notion"
	"github.com/lenismtho/area/cmd/spotify"
	"github.com/lenismtho/area/cmd/twitch"
	"github.com/lenismtho/area/cmd/twitter"
)

func main() {
	go discord.Main()
	go dropbox.Main()
	go ethereum.Main()
	go gmail.Main()
	go github.Main()
	go linkedin.Main()
	go miro.Main()
	go notion.Main()
	go spotify.Main()
	go twitch.Main()
	go twitter.Main()

	core.Main()
}
