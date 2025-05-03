package handlers

import (
	"MyTGbot/internal/storage"

	"gopkg.in/telebot.v3"
)

type StartHandler struct {
	menu *telebot.ReplyMarkup
	db   storage.Database
}

func NewStartHandler(db storage.Database) *StartHandler {
	menu := &telebot.ReplyMarkup{}
	btnAskName := menu.Text("📝 Ввести имя")
	btnCalcSum := menu.Text("🧮 Сумма чисел")
	btnBack := menu.Text("⬅ Назад")

	menu.Reply(
		menu.Row(btnAskName),
		menu.Row(btnCalcSum),
		menu.Row(btnBack),
	)

	return &StartHandler{menu: menu, db: db}
}

func (h *StartHandler) HandleStart(c telebot.Context) error {
	h.db.SetUserState(c.Sender().ID, storage.StateMainMenu)
	return c.Send("Выберите действие:", h.menu)
}
