package telegram

import (
	"context"
	"fmt"
	"strings"

	goodhttp "telegram/internal/clients/good"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type GoodHandler struct {
	client *goodhttp.GoodHTTPClient
}

func InitGoodHandler(client *goodhttp.GoodHTTPClient) *GoodHandler {
	return &GoodHandler{
		client: client,
	}
}

func (g *GoodHandler) RestGood(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var good *goodhttp.GoodHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /rest <id>"))
		return
	}

	rest, err := good.RestOfGood(ctx, args[0])
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}
	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Rest of good %s: %d", args[0], rest)))
}

func (g *GoodHandler) IsGoodAvaliableForOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var good *goodhttp.GoodHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /isavaliable <id>"))
		return
	}

	avaliable, err := good.IsAvaliableForOrder(ctx, args[0])
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	status := "Avaliable!"
	if !avaliable {
		status = "Is not avaliable"
	}
	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Good %s: %d", args[0], status)))
}
