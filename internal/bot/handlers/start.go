package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type StartHandler struct {
	api *tgbotapi.BotAPI
}

func InitStartHandler(api *tgbotapi.BotAPI) *StartHandler {
	return &StartHandler{
		api: api,
	}
}

func (s *StartHandler) Handle(msg *tgbotapi.Message) {

}

func (s *StartHandler) Unknown(msg *tgbotapi.Message) {

}
