package telegram

import (
	"context"
	"fmt"
	"strings"

	orderhttp "telegram/internal/clients/order"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type OrderHandler struct {
	client *orderhttp.OrderHTTPClient
}

func InitOrderHandler(client *orderhttp.OrderHTTPClient) *OrderHandler {
	return &OrderHandler{
		client: client,
	}
}

func (g *OrderHandler) AddOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var order *orderhttp.OrderHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /isavaliable <id>"))
		return
	}

	add, err := order.AddOrder(ctx, args[0])
	if err != nil {
		bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Order %s added", add)))
}

func (g *OrderHandler) DeleteOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var order *orderhttp.OrderHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /isavaliable <id>"))
		return
	}

	err := order.DeleteOrder(ctx, args[0])
	if err != nil {
		bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Order deleted")))
}

func (g *OrderHandler) GetAllOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var order *orderhttp.OrderHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /isavaliable <id>"))
		return
	}

	add, err := order.GetAllOrders(ctx)
	if err != nil {
		bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Orders %s", add)))
}

func (g *OrderHandler) GetByIdOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var order *orderhttp.OrderHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /isavaliable <id>"))
		return
	}

	add, err := order.GetByIdOrder(ctx, args[0])
	if err != nil {
		bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Order %s", add)))
}

func (g *OrderHandler) UpdateOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var order *orderhttp.OrderHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /isavaliable <id>"))
		return
	}

	err := order.UpdateOrder(ctx, args[0])
	if err != nil {
		bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Order updated")))
}
