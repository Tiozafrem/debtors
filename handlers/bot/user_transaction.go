package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

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
