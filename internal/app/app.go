// Created by Catborisvv (c) 2020-2024
// Telegram-Bot, APP-Source

package app

import (
	"fmt"
	"log"
	"telegramBot/configs"

	"telegramBot/internal/app/database"
	"telegramBot/internal/app/keyboards"

	openai "telegramBot/internal/pkg/openai"

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
		if update.Message == nil {
			continue
		}

		// ChatGPT Handler
		if !update.Message.IsCommand() {
			err, ureq := database.GetUserRequest(update.Message.Chat.ID)

			if err != nil {
				log.Printf("[LOG]: Возникла ошибка при обработке значений пользователя %d\n%v", update.Message.Chat.ID, err)
			}

			// Проверка, не закончились ли запросы у пользователя
			if ureq <= 0 {
				ErrorMsg := tgbotapi.NewMessage(update.Message.Chat.ID, configs.RequestsLimitError)
				ErrorMsg.ReplyMarkup = keyboards.GetShopKeyboard()
				bot.Send(ErrorMsg)
			} else {
				WaitMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "⌛ Пожалуйста, подождите... Ваш запрос обрабатывается")
				bot.Send(WaitMsg)

				err, resp := openai.GenerateResponse(update.Message.Text)
				if err != nil {
					resp = configs.ChatGPTGenerateError
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp)
				bot.Send(msg)
			}
			// Отнятие одного запроса у пользователя
			err = database.RunRequest(fmt.Sprintf("UPDATE `users` SET `requests` = `requests` - '1' WHERE `chat_id` = '%d';", update.Message.Chat.ID))
			if err != nil {
				log.Printf("Возникла ошибка обработки запроса:\n\n%v\n", err)
			}
		}

		// Start Command handler
		switch update.Message.Command() {
		case "start":
			{
				ErrorRead, StringContent := ReadStartFile()
				var msg tgbotapi.MessageConfig

				if ErrorRead != nil {
					log.Printf("Возникла ошибка при прочтении файла:\n%v\n", ErrorRead)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, configs.StartMessageError)
				} else {
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, StringContent)
				}
				bot.Send(msg)

				err, u := database.SelectData(fmt.Sprintf("SELECT * FROM `users` WHERE `chat_id` = '%d';", update.Message.Chat.ID))
				if err != nil {
					log.Printf("Возникла ошибка при получении данных пользователя")
				} else {
					if u.Chat_id == 0 {
						database.RunRequest(fmt.Sprintf("INSERT INTO `users` (`id`, `chat_id`, `username`, `requests`, `admin`) VALUES (NULL, '%d', '%s', '100', '0');", update.Message.Chat.ID, update.Message.From.UserName))
					} else {
						fmt.Printf("[LOG] Пользователь %s (%d) использовал команду /start, но он зарегистрирован\n", update.Message.From.UserName, update.Message.Chat.ID)
					}
				}
			}
		}
	}
}
