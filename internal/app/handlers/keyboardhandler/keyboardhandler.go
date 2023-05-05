package keyboardhandler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ShopKeyBoardHandler(update tgbotapi.Update, bot tgbotapi.BotAPI) {
	go func() {
		callbackData := update.CallbackQuery.Data
		if callbackData == "100ReqCall" || callbackData == "200ReqCall" || callbackData == "500ReqCall" || callbackData == "1000ReqCall" {
			msg := tgbotapi.NewMessage(update.CallbackQuery.From.ID, "😅 В данный момент система оплаты ещё не готова.\n\nВы можете оплатить услуги за использование у администратора: @catborisovv")
			bot.Send(msg)
		}
	}()
}
