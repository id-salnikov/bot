package commands

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) NextPage(callbackMessage *tgbotapi.Message, parsedData *CommandData) {
	lastIdx := parsedData.LastIdx + parsedData.Offset
	outputMsgText := fmt.Sprintf("Here %d ... %d products:\n", parsedData.LastIdx+1, lastIdx)

	products := c.productService.GetRange(parsedData.LastIdx, lastIdx)
	for _, p := range products {
		outputMsgText += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(callbackMessage.Chat.ID, outputMsgText)

	data, _ := json.Marshal(CommandData{
		Offset:      parsedData.Offset,
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
