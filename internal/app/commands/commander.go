package commands

import (
	"log"

	"github.com/b0gochort/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) DefaultBehavior(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	// msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil { // If we got a message

		switch update.Message.Command() {
		case "help":
			c.HelpCommand(update.Message)
		case "list":
			c.ListCommand(update.Message)
		case "get":
			c.GetCommand(update.Message)
		default:
			c.DefaultBehavior(update.Message)
		}

	}
}