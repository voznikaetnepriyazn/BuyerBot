package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type OrderHandler struct {
	api *tgbotapi.BotAPI
}

func InitOrderHandler(api *tgbotapi.BotAPI) *OrderHandler {
	return &OrderHandler{
		api: api,
	}
}

func (o *OrderHandler) List(msg *tgbotapi.Message) {

}
