package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (h *Handler) getMyDebtors(b *gotgbot.Bot, ctx *ext.Context) error {
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
				{Text: v.UserUUID, CallbackData: fmt.Sprintf("debtor %s", v.UserUUID)},
			})
	}
	_, err = ctx.EffectiveMessage.Reply(b,
		"Click debtor to view more detail",
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

func (h *Handler) getMyDebtor(b *gotgbot.Bot, ctx *ext.Context) error {
	userUUID, err := h.getUserUUIDByTelegramId(ctx.EffectiveUser.Id)
	if userUUID == "" || err != nil {
		ctx.EffectiveMessage.Reply(b, "you must sign in", nil)
		return err
	}

	cb := ctx.Update.CallbackQuery
	userSplit := strings.Split(cb.Data, " ")
	if len(userSplit) < 2 {
		err = fmt.Errorf("uncorect data")
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		return err
	}

	debtorUUID := userSplit[1]
	cntx := context.Background()

	value, err := h.service.User.GetSumTransactionDebtor(cntx, userUUID, debtorUUID)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, err.Error(), nil)
		return err
	}
	_, err = ctx.EffectiveChat.SendMessage(b, fmt.Sprintf("%s : %d", debtorUUID, value), nil)
	return err
}

func (h *Handler) getSumTransactionDebtorsUser(b *gotgbot.Bot, ctx *ext.Context) error {
	userUUID, err := h.getUserUUIDByTelegramId(ctx.EffectiveUser.Id)
	if userUUID == "" || err != nil {
		ctx.EffectiveMessage.Reply(b, "you must sign in", nil)
		return err
	}

	cntx := context.Background()
	value, err := h.service.User.GetSumTransactionDebtors(cntx, userUUID)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, err.Error(), nil)
		return err
	}

	return h.sendMapMessage(b, ctx, value)
}

func (h *Handler) getSumTransactionMy(b *gotgbot.Bot, ctx *ext.Context) error {
	userUUID, err := h.getUserUUIDByTelegramId(ctx.EffectiveUser.Id)
	if userUUID == "" || err != nil {
		ctx.EffectiveMessage.Reply(b, "you must sign in", nil)
		return err
	}

	cntx := context.Background()
	value, err := h.service.User.GetSumMy(cntx, userUUID)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, err.Error(), nil)
		return err
	}

	return h.sendMapMessage(b, ctx, value)
}

func (h *Handler) sendMapMessage(b *gotgbot.Bot, ctx *ext.Context, message map[string]int) error {
	builder := strings.Builder{}

	for key, value := range message {
		builder.WriteString(fmt.Sprintf("%s : %d\n", key, value))
	}

	_, err := ctx.EffectiveChat.SendMessage(b, builder.String(), nil)
	return err
}
