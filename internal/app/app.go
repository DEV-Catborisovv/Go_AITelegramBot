// Created by Catborisvv (c) 2020-2024
// Telegram-Bot, APP-Source

package app

import (
	"log"
	"telegramBot/configs"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot() {
	// Запуск бота, создается объект бота (API-Токен взят из конфига)
	bot, err := tgbotapi.NewBotAPI(configs.TELEGRAM_API_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	// DEBUG MODE?
	bot.Debug = configs.DEBUG_MODE
	log.Printf("Авторизация бота: %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = configs.TIMEOUT

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			if update.Message.Text == "/start" {
				ErrorRead, StringContent := ReadStartFile()
				var msg tgbotapi.MessageConfig

				if ErrorRead != nil {
					log.Printf("Возникла ошибка при прочтении файла:\n%v\n", ErrorRead)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, configs.StartMessageError)
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, StringContent)
				}
				msg.ReplyToMessageID = update.Message.MessageID
				log.Println(update.Message.From.ID)

				bot.Send(msg)
			}
		}
	}
}
