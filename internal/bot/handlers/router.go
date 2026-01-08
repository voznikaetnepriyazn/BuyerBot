package telegram

import (
	"context"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Router struct {
	customerHandler *CustomerHandler
	orderHandler    *OrderHandler
	goodHandler     *GoodHandler
}

func NewRouter(customerHandler *CustomerHandler, orderHandler *OrderHandler, goodHandler *GoodHandler) *Router {
	return &Router{
		customerHandler: customerHandler,
		orderHandler:    orderHandler,
		goodHandler:     goodHandler,
	}
}

func (r *Router) HandleUpdate(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		r.handleCommand(ctx, bot, update)
		return
	}

	r.handleMessage(ctx, bot, update)
}

func (r *Router) handleCommand(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	cmd := strings.ToLower(update.Message.Command())

	switch cmd {
	case "start":
		bot.Send(tgbotapi.NewMessage(chatID, "Привет! Используйте /addcustomer для добавления покупателя."))

	case "rest":
		r.goodHandler.RestGood(ctx, bot, update)

	case "isavaliable":
		r.goodHandler.IsGoodAvaliableForOrder(ctx, bot, update)

	case "addcustomer":
		r.customerHandler.StartAddCustomer(ctx, bot, update)
		r.customerHandler.HandleMessage(ctx, bot, update)
		r.customerHandler.AddCustomer(ctx, bot, update)

	case "addorder":
		r.orderHandler.StartAddOrder(ctx, bot, update)
		r.orderHandler.AddOrder(ctx, bot, update)

	/*case "deleteorder":
	r.orderHandler.DeleteOrder(ctx, bot, update)
	r.orderHandler.HandleCallback(ctx, bot, update)*/

	case "getallorders":
		r.orderHandler.GetAllOrders(ctx, bot, update)

	case "gatbyidorder":
		r.orderHandler.GetOrderById(ctx, bot, update)

	default:
		bot.Send(tgbotapi.NewMessage(chatID, "Неизвестная команда: /"+cmd))
	}
}

func (r *Router) handleMessage(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// Передаём сообщение в обработчик, который отвечает за диалоговые состояния
	r.customerHandler.HandleMessage(ctx, bot, update)
}
