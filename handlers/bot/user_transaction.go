package bot

import (
	"context"
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// func (h *Handler) getSumTransactionDebtorUser(b *gotgbot.Bot, ctx *ext.Context) error {
// 	id := c.Param("uuid")
// 	if id == "" {
// 		ctx.EffectiveMessage.Reply(b, "you must sign in", nil)
// 		return err
// 		newErrorResponse(c, http.StatusBadRequest, "uuid is null")
// 		return nil
// 	}

// 	userUUID, err := getUserUUID(c)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}
// 	value, err := h.service.User.GetSumTransactionDebtor(c, userUUID, id)
// 	if err != nil {
// 		newErrorResponse(c, http.StatusBadRequest, err.Error())
// 		return nil
// 	}

// 	c.JSON(http.StatusOK, value)
// 	return nil
// }

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

	ctx.EffectiveChat.SendMessage(b, fmt.Sprintln(value), nil)
	return nil
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

	ctx.EffectiveChat.SendMessage(b, fmt.Sprintln(value), nil)
	return nil
}
