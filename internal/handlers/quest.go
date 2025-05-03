package handlers

import (
	"MyTGbot/internal/storage"

	"gopkg.in/telebot.v3"
)

type QuestHandler struct {
	//menu *telebot.ReplyMarkup
	db storage.Database
}

func (h *QuestHandler) HandleAskName(c telebot.Context) error {
	h.db.SetUserState(c.Sender().ID, storage.StateWaitingForName)
	return c.Send("Введите ваше имя:")
}

func (h *QuestHandler) HandleCalcSum(c telebot.Context) error {
	h.db.SetUserState(c.Sender().ID, storage.StateWaitingForNumbers)
	return c.Send("Введите два числа через пробел (например: 5 10):")
}
