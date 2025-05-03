package handlers

import (
	"MyTGbot/internal/storage"

	"gopkg.in/telebot.v3"
)

func Register(bot *telebot.Bot, db storage.Database) {
	startHandler := NewStartHandler(db)
	questHandler := &QuestHandler{db: db}
	textHandler := &TextHandler{db: db, menu: startHandler.menu}

	bot.Handle("/start", startHandler.HandleStart)
	bot.Handle(&startHandler.menu.ReplyKeyboard[0][0], questHandler.HandleAskName) // Кнопка "Ввести имя"
	bot.Handle(&startHandler.menu.ReplyKeyboard[1][0], questHandler.HandleCalcSum) // Кнопка "Сумма чисел"
	bot.Handle(telebot.OnText, textHandler.HandleText)
}
