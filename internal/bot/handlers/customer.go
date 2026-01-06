package telegram

import (
	"context"
	"fmt"
	"strings"
	customerhttp "telegram/internal/clients/customer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CustomerHandler struct {
	client *customerhttp.CustomerHTTPClient
}

func InitCustomerHandler(client *customerhttp.CustomerHTTPClient) *CustomerHandler {
	return &CustomerHandler{
		client: client,
	}
}

func (g *GoodHandler) AddCustomer(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var customer *customerhttp.CustomerHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /isavaliable <id>"))
		return
	}

	add, err := customer.AddCustomer(ctx, args[0])
	if err != nil {
		bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgborapi.NewMessage(chatID, fmt.Sprintf("Customer %s", add)))
}
