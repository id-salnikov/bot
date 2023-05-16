package main

import (
	"github.com/id-salnikov/bot/internal/app/commands"
	"github.com/id-salnikov/bot/internal/service/product"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
