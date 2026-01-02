package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	start *StartHandler
	order *OrderHandler
}

func (r *Router) Route(msg *tgbotapi.Message) {
	switch msg.Command() {
	case "start":
		r.start.Handle(msg)
	case "order":
		r.order.List(msg)
	default:
		r.start.Unknown(msg)
	}

}
