package api

import (
	"browser-mmo-backend/internal/data"
	"browser-mmo-backend/internal/gamecontent"
)

type Application struct {
	Models      data.Models
	GameContent gamecontent.GameContent
}
