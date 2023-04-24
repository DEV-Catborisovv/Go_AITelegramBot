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
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
