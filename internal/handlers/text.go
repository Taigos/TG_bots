package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"MyTGbot/internal/storage"

	"gopkg.in/telebot.v3"
)

type TextHandler struct {
	menu *telebot.ReplyMarkup
	db   storage.Database
}

func (h *TextHandler) HandleText(c telebot.Context) error {
	userID := c.Sender().ID
	text := c.Text()

	switch h.db.GetUserState(userID) {
	case storage.StateWaitingForName:
		h.db.SetUserState(userID, storage.StateMainMenu)
		return c.Send("Привет, "+text+"! 😊", h.menu)

	case storage.StateWaitingForNumbers:
		nums := strings.Fields(text)
		if len(nums) != 2 {
			return c.Send("Нужно 2 числа через пробел!")
		}

		a, err1 := strconv.Atoi(nums[0])
		b, err2 := strconv.Atoi(nums[1])
		if err1 != nil || err2 != nil {
			return c.Send("Это не числа!")
		}

		h.db.SetUserState(userID, storage.StateMainMenu)
		return c.Send(fmt.Sprintf("Сумма: %d", a+b), h.menu)

	default:
		return c.Send("Выберите действие:", h.menu)
	}
}
