package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	start *StartHandler
	order *OrderHandler
}

func (r *Router) Route(msg *tgbotapi.Message) {
	switch msg.Command() || bot.Command() {
	case "start":
		r.start.Handle(msg)
	case "order":
		r.order.GetByIdOrder(ctx context.Context, bot, update)
	default:
		r.start.Unknown(msg)
	}

}
