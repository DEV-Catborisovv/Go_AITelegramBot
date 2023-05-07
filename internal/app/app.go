// Created by Catborisvv (c) 2020-2024
// Telegram-Bot, APP-Source

package app

import (
	"log"
	"telegramBot/configs"

	"telegramBot/internal/app/handlers/chatgpt"
	"telegramBot/internal/app/handlers/keyboardhandler"
	"telegramBot/internal/app/handlers/settings"
	"telegramBot/internal/app/handlers/startcommand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Функция пулинга бота (запуск обработки сообщений)
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

		if update.CallbackQuery != nil {
			keyboardhandler.ShopKeyBoardHandler(update, *bot)
		}

		if update.Message == nil {
			continue
		}

		// ChatGPT Handler
		if !update.Message.IsCommand() {
			chatgpt.ChatGPTHandler(update, *bot)
		}

		// Start Command handler
		switch update.Message.Command() {
		case "start":
			{
				startcommand.StartCommandHandler(update, *bot)
			}
		case "settings":
			{
				settings.SettingsCommandHandler(update, *bot)
			}
		}
	}
}
