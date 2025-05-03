package app

import (
	"MyTGbot/internal/config"
	"MyTGbot/internal/handlers"
	"MyTGbot/internal/storage"
	"log"

	"gopkg.in/telebot.v3"
)

func Run(cfg *config.Config) {
	// Инициализация хранилища
	store := storage.NewMemoryStorage()

	// Создание бота
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  cfg.Token,
		Poller: &telebot.LongPoller{Timeout: 10},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Регистрация обработчиков
	handlers.Register(bot, store)

	log.Println("Бот запущен...")
	bot.Start()
}
