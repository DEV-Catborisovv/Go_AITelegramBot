// Обработчик команды /start
// Created by Catborisovv (c) 2020-2024

package settings

import (
	"fmt"
	"log"
	"telegramBot/internal/app/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Call-back settings command from Go
func SettingsCommandHandler(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	go func() {
		err, u := database.SelectData(fmt.Sprintf("SELECT * FROM `users` WHERE `chat_id` = '%d';", update.Message.Chat.ID))
		if err != nil {
			log.Fatalf("Возникла ошибка при обработке запроса к базе данных:\n%v", err)
		}
		UserDataF := fmt.Sprintf("⚙️ %s, Вы перешли в настройки.\n\n- У Вас осталось %d запросов.\n\n😺 Ваш Telegram ID: %d (#%d)", update.Message.From.FirstName, u.Requests, update.Message.From.ID, u.ID)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, UserDataF)
		bot.Send(msg)
	}()
}
