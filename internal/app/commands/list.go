package commands

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(message *tgbotapi.Message) {
	lastIdx := 5
	outputMsgText := fmt.Sprintf("Here %d ... %d products:\n", 1, lastIdx)
	products := c.productService.GetRange(0, lastIdx)
	for _, p := range products {
		outputMsgText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, outputMsgText)

	data, _ := json.Marshal(CommandData{
		Offset:      5,
		CommandName: "next_page",
		LastIdx:     lastIdx,
	})

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", string(data)),
		),
	)

	c.bot.Send(msg)
}
