package bot

import (
	"context"
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/conversation"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/callbackquery"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/tiozafrem/debtors/services"
)

const (
	email    = "email"
	password = "password"
	value    = "value"
)

type userInput struct {
	email    string
	password string
	endpoint func(b *gotgbot.Bot, ctx *ext.Context) error
}

type Handler struct {
	service    *services.Service
	dispatcher *ext.Dispatcher
	usersEmail map[int64]*userInput
}

func NewHandler(service *services.Service, dispatcher *ext.Dispatcher) *Handler {
	handler := &Handler{service: service, dispatcher: dispatcher, usersEmail: map[int64]*userInput{}}
	handler.InitRoutes()
	return handler
}

func (h *Handler) InitRoutes() {
	h.dispatcher.AddHandler(handlers.NewCommand("start", start))

	handlerAuth := map[string][]ext.Handler{
		email:    {handlers.NewMessage(noCommands, h.getEmail)},
		password: {handlers.NewMessage(noCommands, h.getPassword)},
	}
	handlerOpts := &handlers.ConversationOpts{
		StateStorage: conversation.NewInMemoryStorage(conversation.KeyStrategySenderAndChat),
		AllowReEntry: true,
	}

	h.dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCommand("login", h.login)},
		handlerAuth,
		handlerOpts,
	))
	h.dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCommand("register", h.register)},
		handlerAuth,
		handlerOpts,
	))

	h.dispatcher.AddHandler(handlers.NewCommand("pin_print", h.getUsers))
	h.dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix("pin_user"), h.pinUserToUser))

	h.dispatcher.AddHandler(handlers.NewCommand("my_debtors", h.getSumTransactionDebtorsUser))
	h.dispatcher.AddHandler(handlers.NewCommand("i_debtor", h.getSumTransactionMy))

	h.dispatcher.AddHandler(handlers.NewCommand("debtors_print", h.getMyDebtors))
	h.dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix("debtor"), h.getMyDebtor))

	h.dispatcher.AddHandler(handlers.NewConversation(
		[]ext.Handler{handlers.NewCommand("add_value", h.printAddTransaction)},
		map[string][]ext.Handler{
			value: {handlers.NewMessage(noCommands, h.addTransaction)},
		},
		handlerOpts,
	))
	h.dispatcher.AddHandler(handlers.NewCallback(callbackquery.Prefix(value), h.addTransactionCallback))

}

func noCommands(msg *gotgbot.Message) bool {
	return message.Text(msg) && !message.Command(msg)
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "write /register or /login to continue work",
		nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil

}

func (h *Handler) getUserUUIDByTelegramId(id int64) (string, error) {
	return h.service.GetUUIDByTelegramId(context.Background(), fmt.Sprint(id))
}
