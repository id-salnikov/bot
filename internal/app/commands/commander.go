package commands

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/id-salnikov/bot/internal/service/product"
	"log"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

type CommandData struct {
	CommandName string `json:"command_name"`
	Offset      int    `json:"offset"`
	LastIdx     int    `json:"last_idx"`
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		parsedData := CommandData{}
		err := json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		if err != nil {
			log.Println(err)
			return
		}

		switch parsedData.CommandName {
		case "next_page":
			c.NextPage(update.CallbackQuery.Message, &parsedData)
		default:
			c.Default(update.CallbackQuery.Message)
		}
		return
	}

	if update.Message == nil {
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
