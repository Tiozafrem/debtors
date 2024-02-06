package bot

import (
	"context"
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func (h *Handler) getUsers(b *gotgbot.Bot, ctx *ext.Context) error {
	cntx := context.Background()
	userUUID, err := h.getUserUUIDByTelegramId(ctx.EffectiveUser.Id)
	if userUUID == "" || err != nil {
		ctx.EffectiveMessage.Reply(b, "you must sign in", &gotgbot.SendMessageOpts{})
		return err
	}

	users, err := h.service.User.GetUsersNotMy(cntx,
		userUUID)
	if err != nil {
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		return err
	}
	users_answer := [][]gotgbot.InlineKeyboardButton{{}}
	for _, v := range users {
		users_answer = append(users_answer,
			[]gotgbot.InlineKeyboardButton{
				{Text: v.UserUUID, CallbackData: fmt.Sprintf("pin_user %s", v.UserUUID)},
			})
	}
	_, err = ctx.EffectiveMessage.Reply(b,
		"Click youser to pin in your accaount. After you can add value debtor value",
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

func (h *Handler) pinUserToUser(b *gotgbot.Bot, ctx *ext.Context) error {
	userUUID, err := h.getUserUUIDByTelegramId(ctx.EffectiveUser.Id)
	if userUUID == "" || err != nil {
		ctx.EffectiveMessage.Reply(b, "you must sign in", &gotgbot.SendMessageOpts{})
		return err
	}

	cb := ctx.Update.CallbackQuery
	userSplit := strings.Split(cb.Data, " ")
	if len(userSplit) < 2 {
		err = fmt.Errorf("uncorect data")
		ctx.EffectiveChat.SendMessage(b, err.Error(), nil)
		return err
	}

	userPin := userSplit[1]
	cntx := context.Background()

	err = h.service.User.PinUserToUser(cntx, userUUID, userPin)
	if err != nil {
		ctx.EffectiveMessage.Reply(b, err.Error(), &gotgbot.SendMessageOpts{})
		return err
	}
	_, _, err = cb.Message.EditText(b, fmt.Sprintf("%s successfully pin", userPin), nil)
	return err
}
