package bot

import (
	"log"
	"strconv"

	"github.com/Ensena/core/env-global"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var SendID int64
var bot *tgbotapi.BotAPI

func init() {
	token := env.Check("telegram_token", "Missing telegram_token")
	groupID := env.Check("telegram_group", "Missing telegram_group")
	var err error
	SendID, _ = strconv.ParseInt(groupID, 10, 64)
	if err != nil {
		log.Println("Fail to parse ID")
	}
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Warning Failed to connect to telegram")
	}

}
