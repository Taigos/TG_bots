package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
	//myToken := "7008104386:AAG4_P77JyB2GwZB5sOPfRqOtCJKEPNTl0Q"

	"MyTGbot/internal/app"
	"MyTGbot/internal/config"
)

// Состояния бота
const (
	StateMainMenu = iota
	StateWaitingForName
	StateWaitingForNumbers
)

// Глобальные переменные
var (
	userStates = make(map[int64]int) // Храним состояния пользователей
)

func oldFunc() {
	myToken := ""
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  myToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Главное меню
	mainMenu := &telebot.ReplyMarkup{}
	btnAskName := mainMenu.Text("📝 Ввести имя")
	btnCalcSum := mainMenu.Text("🧮 Сумма чисел")
	btnBack := mainMenu.Text("⬅ Назад")

	mainMenu.Reply(
		mainMenu.Row(btnAskName),
		mainMenu.Row(btnCalcSum),
		mainMenu.Row(btnBack),
	)

	// Обработчики
	bot.Handle(&btnBack, func(c telebot.Context) error {
		userStates[c.Sender().ID] = StateMainMenu
		return c.Send("Главное меню:", mainMenu)
	})
	bot.Handle("/start", func(c telebot.Context) error {
		userStates[c.Sender().ID] = StateMainMenu
		return c.Send("Выберите действие:", mainMenu)
	})

	bot.Handle(&btnAskName, func(c telebot.Context) error {
		userStates[c.Sender().ID] = StateWaitingForName
		return c.Send("Введите ваше имя:")
	})

	bot.Handle(&btnCalcSum, func(c telebot.Context) error {
		userStates[c.Sender().ID] = StateWaitingForNumbers
		return c.Send("Введите два числа через пробел (например: 5 10):")
	})

	// Обработка текстовых сообщений
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		userID := c.Sender().ID
		text := c.Text()

		switch userStates[userID] {
		case StateWaitingForName:
			userStates[userID] = StateMainMenu
			return c.Send("Привет, "+text+"! 😊", mainMenu)

		case StateWaitingForNumbers:
			nums := strings.Fields(text)
			if len(nums) != 2 {
				return c.Send("Нужно 2 числа через пробел!")
			}

			a, err1 := strconv.Atoi(nums[0])
			b, err2 := strconv.Atoi(nums[1])
			if err1 != nil || err2 != nil {
				return c.Send("Это не числа!")
			}

			userStates[userID] = StateMainMenu
			return c.Send(fmt.Sprintf("Сумма: %d", a+b), mainMenu)

		default:
			return c.Send("Выберите действие:", mainMenu)
		}
	})

	bot.Start()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Паника: %v", r)
			time.Sleep(10 * time.Second) // Чтобы увидеть ошибку
		}
	}()
}
func main() {
	cfg := config.Load()
	app.Run(cfg)

	oldFunc()
}
