package telegram

import (
	"context"
	"fmt"
	"strings"
	"sync"
	customerhttp "telegram/internal/clients/customer"
	customer "telegram/internal/model/customer"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type State string

const (
	StateNone                State = ""
	StateAwaitingName              = "awaiting_name"
	StateAwaitingEmail             = "awaiting_email"
	StateAwaitingCity              = "awaiting_city"
	StateAwaitingFullAdress        = "awaiting_adress"
	StateAwaitingPostalCode        = "awaiting_postal_code"
	StateAwaitingPhoneNumber       = "awaiting_phone_number"
)

type CustomerDraft struct {
	Name        string
	Email       string
	City        string
	FullAddress string
	PostalCode  string
	PhoneNumber string
}

type CustomerHandler struct {
	states map[int64]State
	drafts map[int64]*CustomerDraft
	mu     sync.RWMutex
	client *customerhttp.CustomerHTTPClient
}

func InitCustomerHandler(client *customerhttp.CustomerHTTPClient, states map[int64]State) *CustomerHandler {
	return &CustomerHandler{
		client: client,
		states: make(map[int64]State),
		drafts: make(map[int64]*CustomerDraft),
	}
}

func (g *CustomerHandler) StartAddCustomer(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	g.mu.Lock()
	g.states[chatID] = StateAwaitingName
	g.drafts[chatID] = &CustomerDraft{}
	g.mu.Unlock()

	bot.Send(tgbotapi.NewMessage(chatID, "Введите имя покупателя:"))
}

func (g *CustomerHandler) HandleMessage(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var custom customer.Customer

	if update.Message != nil && update.Message.Text != "" {
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

	case StateAwaitingName:
		draft.Name = text
		g.mu.Lock()
		g.states[chatID] = StateAwaitingEmail
		g.mu.Unlock()
		bot.Send(tgbotapi.NewMessage(chatID, "Введите email покупателя:"))

	case StateAwaitingEmail:
		draft.Email = text
		g.mu.Lock()
		g.states[chatID] = StateAwaitingCity
		g.mu.Unlock()
		bot.Send(tgbotapi.NewMessage(chatID, "Введите город:"))

	case StateAwaitingCity:
		draft.City = text
		g.mu.Lock()
		g.states[chatID] = StateAwaitingFullAdress
		g.mu.Unlock()
		bot.Send(tgbotapi.NewMessage(chatID, "Введите адресс"))

	case StateAwaitingFullAdress:
		draft.FullAddress = text
		g.mu.Lock()
		g.states[chatID] = StateAwaitingPostalCode
		g.mu.Unlock()
		bot.Send(tgbotapi.NewMessage(chatID, "Введите почтовый код"))

	case StateAwaitingPhoneNumber:
		draft.PhoneNumber = text

		cust, err := g.client.AddCustomer(
			ctx, custom,
		)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(chatID, "Ошибка при создании покупателя: "))
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Покупатель создан:\nИмя: %s\nEmail: %s", cust.Name, cust.Email)))
		}
	}
}

func (g *CustomerHandler) AddCustomer(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	var cust customer.Customer

	var customer *customerhttp.CustomerHTTPClient

	args := strings.Fields(update.Message.CommandArguments())
	if len(args) != 1 {
		bot.Send(tgbotapi.NewMessage(chatID, "используйте - /addcustomer"))
		return
	}

	add, err := customer.AddCustomer(ctx, cust)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Error: %d", err)))
		return
	}

	bot.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Customer %v", add)))
}
