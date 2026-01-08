package telegram

import (
	"context"
	"fmt"

	//"log"
	"strings"
	"sync"

	orderhttp "telegram/internal/clients/order"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type StateOrder string

const (
	StateNoneOrder        State = ""
	StateAwaitingIdOfGood       = "awaiting_IDofGood"
)

type OrderDraft struct {
	IdOfGood string
}

type OrderHandler struct {
	states map[int64]State
	drafts map[int64]*OrderDraft
	mu     sync.RWMutex
	client *orderhttp.OrderHTTPClient
}

func InitOrderHandler(client *orderhttp.OrderHTTPClient, states map[int64]State) *OrderHandler {
	return &OrderHandler{
		client: client,
		states: make(map[int64]State),
		drafts: make(map[int64]*OrderDraft),
	}
}

func (g *OrderHandler) StartAddOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	g.mu.Lock()
	g.states[chatID] = StateAwaitingIdOfGood
	g.drafts[chatID] = &OrderDraft{}
	g.mu.Unlock()

	bot.Send(tgbotapi.NewMessage(chatID, "Введите ид желаемых товаров:"))
}

func (g *OrderHandler) HandleMessage(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil && update.Message.Text == "" {
		return
	}

	chatID := update.Message.Chat.ID

	g.mu.RLock()
	state := g.states[chatID]
	draft := g.drafts[chatID]
	g.mu.RUnlock()

	if state == StateNone {
		return
	}

	text := update.Message.Text

	switch state {

	case StateAwaitingIdOfGood:
		draft.IdOfGood = text

		cust, err := g.client.AddOrder(
			ctx, draft.IdOfGood,
		)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при создании заказа"))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Заказ %s создан", cust)))
		}
	}
}

func (g *OrderHandler) AddOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var order *orderhttp.OrderHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /addorder <id>"))
		return
	}

	add, err := order.AddOrder(ctx, args[0])
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Order %s added", add)))
}

/*func (g *OrderHandler) DeleteOrder(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "do you want to delete this order?"))
		return
	}

	orderID := args[0]

	deleteButton := tgbotapi.NewInlineKeyboardButtonData("Удалить", "delete_order:"+orderID)
	row := tgbotapi.NewInlineKeyboardRow(deleteButton)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row)

	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Вы уверены, что хотите удалить заказ %s?", orderID))
	msg.ReplyMarkup = keyboard

	bot.Send(msg)
}

func (g *OrderHandler) HandleCallback(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := update.CallbackQuery
	chatID := callback.Message.Chat.ID
	data := callback.Data

	defer func() {
		err := bot.AnswerCallbackQuery(tgbotapi.CallbackConfig{
			CallbackQueryID: callback.ID,
			Text:            "",
			ShowAlert:       false,
		})
		if err != nil {
			log.Printf("Ошибка подтверждения callback: %v", err)
		}
	}()

	if strings.HasPrefix(data, "delete_order:") {
		orderID := strings.TrimPrefix(data, "delete_order:")

		err := g.client.DeleteOrder(ctx, orderID)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Ошибка удаления: "+err.Error()))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Заказ %s удалён", orderID)))
		}

	} else if data == "cancel" {
		bot.Send(tgbotapi.NewMessage(chatID, "Операция отменена."))
	}
}*/

func (g *OrderHandler) GetAllOrders(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	orders, err := g.client.GetAllOrders(ctx)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Ошибка получения заказов: "+err.Error()))
		return
	}

	if len(orders) == 0 {
		bot.Send(tgbotapi.NewMessage(chatID, "У вас нет заказов."))
		return
	}

	var msgText strings.Builder
	msgText.WriteString("Ваши заказы:\n")
	for _, order := range orders {
		msgText.WriteString(fmt.Sprintf("- Заказ %s\n", order.Id))
	}

	bot.Send(tgbotapi.NewMessage(chatID, msgText.String()))
}

func (g *OrderHandler) GetOrderById(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "Используйте: /getorder <id>"))
		return
	}

	orderID := args[0]
	order, err := g.client.GetByIdOrder(ctx, orderID)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, "Ошибка: "+err.Error()))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Заказ %s", order.Id)))
}
