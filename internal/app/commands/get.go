package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	idx, err := strconv.Atoi(inputMessage.CommandArguments())
	outMsg := ""
	if err != nil {
		outMsg = `wrong args, follow "/get 0" template`
		log.Println(outMsg)
	} else {
		product, err := c.productService.Get(idx)
		if err != nil {
			outMsg = fmt.Sprintf("Fail to get product %d: %v", idx, err)
			log.Println(outMsg)
		}
		outMsg = product.Title
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		outMsg,
	)

	c.bot.Send(msg)
}
