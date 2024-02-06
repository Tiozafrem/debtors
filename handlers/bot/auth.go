package bot

import (
	"context"
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func (h *Handler) login(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "write email",
		nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	h.usersEmail[ctx.EffectiveUser.Id] = &userInput{endpoint: h.signIn}
	return handlers.NextConversationState(email)

}

func (h *Handler) register(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "write email",
		nil)
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	h.usersEmail[ctx.EffectiveUser.Id] = &userInput{endpoint: h.signUp}
	return handlers.NextConversationState(email)

}

func (h *Handler) signUp(b *gotgbot.Bot, ctx *ext.Context) error {
	user := h.usersEmail[ctx.EffectiveUser.Id]
	cntx := context.Background()

	token, err := h.service.Authorization.SignUp(cntx, user.email, user.password)
	if err != nil {
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		ctx.EffectiveChat.SendMessage(b, "write email", nil)
		return handlers.NextConversationState(email)
	}

	uuid, err := h.service.ParseTokenToUserUUID(cntx, token.AccessToken)
	if err != nil {
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		ctx.EffectiveChat.SendMessage(b, "write email", nil)
		return handlers.NextConversationState(email)
	}

	err = h.service.User.PinTelegramId(cntx, uuid, fmt.Sprint(ctx.EffectiveUser.Id))
	if err != nil {
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		ctx.EffectiveChat.SendMessage(b, "write email", nil)
		return handlers.NextConversationState(email)
	}

	return handlers.EndConversation()
}

func (h *Handler) signIn(b *gotgbot.Bot, ctx *ext.Context) error {
	user := h.usersEmail[ctx.EffectiveUser.Id]
	cntx := context.Background()

	token, err := h.service.Authorization.SignIn(user.email, user.password)
	if err != nil {
		fmt.Println(err)
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		ctx.EffectiveChat.SendMessage(b, "write email", nil)
		return handlers.NextConversationState(email)
	}

	uuid, err := h.service.ParseTokenToUserUUID(cntx, token.AccessToken)
	if err != nil {
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		ctx.EffectiveChat.SendMessage(b, "write email", nil)
		return handlers.NextConversationState(email)
	}

	err = h.service.User.PinTelegramId(cntx, uuid, fmt.Sprint(ctx.EffectiveUser.Id))
	if err != nil {
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		ctx.EffectiveChat.SendMessage(b, "write email", nil)
		return handlers.NextConversationState(email)
	}

	return handlers.EndConversation()
}

func (h *Handler) getEmail(b *gotgbot.Bot, ctx *ext.Context) error {
	email := ctx.EffectiveMessage.Text
	user := h.usersEmail[ctx.EffectiveUser.Id]
	user.email = email

	_, err := ctx.EffectiveMessage.Reply(b, "write password",
		nil)
	if err != nil {
		return fmt.Errorf("failed to send name message: %w", err)
	}

	return handlers.NextConversationState(password)
}

func (h *Handler) getPassword(b *gotgbot.Bot, ctx *ext.Context) error {
	password := ctx.EffectiveMessage.Text
	user := h.usersEmail[ctx.EffectiveUser.Id]
	user.password = password
	return user.endpoint(b, ctx)
}
