package main

import (
	"context"
	"log"
	"os"

	telegram "telegram/internal/bot/handlers"
	customerhttp "telegram/internal/clients/customer"
	goodhttp "telegram/internal/clients/good"
	orderhttp "telegram/internal/clients/order"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN does not exist")
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	state := make(map[int64]telegram.State)

	customer := customerhttp.InitCustomerHTTPClient("http://localhost:8080")
	good := goodhttp.InitgoodHTTPClient("http://localhost:8080")
	order := orderhttp.InitOrderHTTPClient("http://localhost:8080")

	customerhandler := telegram.InitCustomerHandler(customer, state)
	goodhandler := telegram.InitGoodHandler(good)
	orderhandler := telegram.InitOrderHandler(order, state)

	router := telegram.InitRouter(customerhandler, orderhandler, goodhandler)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	ctx := context.Background()

	for update := range updates {
		router.HandleUpdate(ctx, bot, update)
	}

}
