package chatgpt

import (
	"fmt"
	"log"
	"telegramBot/configs"
	"telegramBot/internal/app/database"
	"telegramBot/internal/app/keyboards"
	"telegramBot/internal/pkg/openai"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Call-back from Go
func ChatGPTHandler(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	// При использовании функции запускаем ее в отдельном потоке, чтобы не перегружать бота
	go func() {
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

			// Отнятие одного запроса у пользователя
			err = database.RunRequest(fmt.Sprintf("UPDATE `users` SET `requests` = `requests` - '1' WHERE `chat_id` = '%d';", update.Message.Chat.ID))
			if err != nil {
				log.Printf("Возникла ошибка обработки запроса:\n\n%v\n", err)
			}
		}
	}()
}
