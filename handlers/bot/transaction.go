package bot

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func (h *Handler) printAddTransaction(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, "write value transaction",
		nil)
	if err != nil {
		return fmt.Errorf("failed to send name message: %w", err)
	}
	return handlers.NextConversationState(value)
}

func (h *Handler) addTransaction(b *gotgbot.Bot, ctx *ext.Context) error {
	valueStr := ctx.EffectiveMessage.Text
	_, err := strconv.Atoi(valueStr)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, err.Error(), nil)
		ctx.EffectiveMessage.Reply(b, "write value", nil)
		return err
	}
	userUUID, err := h.getUserUUIDByTelegramId(ctx.EffectiveUser.Id)
	if userUUID == "" || err != nil {
		ctx.EffectiveMessage.Reply(b, "you must sign in", nil)
		return err
	}

	cntx := context.Background()
	users, err := h.service.User.GetUsersMy(cntx,
		userUUID)
	if err != nil {
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		return err
	}
	users_answer := [][]gotgbot.InlineKeyboardButton{{}}
	for _, v := range users {
		users_answer = append(users_answer,
			[]gotgbot.InlineKeyboardButton{
				{Text: v.UserUUID, CallbackData: fmt.Sprintf("%s %s %s", value, v.UserUUID, valueStr)},
			})
	}
	_, err = ctx.EffectiveMessage.Reply(b,
		"Click debtor to add transaction",
		&gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: users_answer,
			},
		})
	if err != nil {
		return fmt.Errorf("failed to send start message: %w", err)
	}
	return nil
}

func (h *Handler) addTransactionCallback(b *gotgbot.Bot, ctx *ext.Context) error {
	userUUID, err := h.getUserUUIDByTelegramId(ctx.EffectiveUser.Id)
	if userUUID == "" || err != nil {
		ctx.EffectiveMessage.Reply(b, "you must sign in", nil)
		return err
	}

	cb := ctx.Update.CallbackQuery
	userSplit := strings.Split(cb.Data, " ")
	if len(userSplit) < 3 {
		err = fmt.Errorf("uncorect data")
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		return err
	}

	debtorUUID := userSplit[1]
	valueStr := userSplit[2]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, err.Error(), nil)
		ctx.EffectiveMessage.Reply(b, "write value", nil)
		return err
	}
	cntx := context.Background()

	err = h.service.User.AddTransaction(cntx, userUUID, debtorUUID, value)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, err.Error(), nil)
		return err
	}
	_, _, err = cb.Message.EditText(b, fmt.Sprintf("%s successfully add %s", debtorUUID, valueStr), nil)
	return err
}
