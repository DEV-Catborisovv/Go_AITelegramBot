// Start Command Bot Handler
// Created by Catborisovv (c) 2020-2024

package startcommand

import (
	"fmt"
	"log"
	"telegramBot/configs"
	"telegramBot/internal/app/database"
	"telegramBot/internal/app/filereader"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Call-back from Go
func StartCommandHandler(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	go func() {
		ErrorRead, StringContent := filereader.ReadStartFile()
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
				err = database.RunRequest(fmt.Sprintf("INSERT INTO `users` (`id`, `chat_id`, `username`, `requests`, `admin`) VALUES (NULL, '%d', '%s', '100', '0');", update.Message.Chat.ID, update.Message.From.UserName))
				if err != nil {
					log.Fatalf("Возникла ошибка при обраблотке значения: %s\n", err)
				}
			} else {
				fmt.Printf("[LOG] Пользователь %s (%d) использовал команду /start, но он зарегистрирован\n", update.Message.From.UserName, update.Message.Chat.ID)
			}
		}
	}()
}
