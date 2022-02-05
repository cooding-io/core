package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func Send(ID int64, message string) {
	if ID != 0 {
		msg := tgbotapi.NewMessage(ID, message)
		bot.Send(msg)
	}
}
