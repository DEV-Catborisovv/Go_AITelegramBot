// KeyBoards-Getter, :D
// Created by Catborisovv (c) 2020-2024

package keyboards

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// функция получения клавиатуры для оплаты бота
func GetShopKeyboard() tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("100 запросов", "100ReqCall"),
			tgbotapi.NewInlineKeyboardButtonData("200 запросов", "200ReqCall"),
			tgbotapi.NewInlineKeyboardButtonData("500 запросов", "500ReqCall"),
			tgbotapi.NewInlineKeyboardButtonData("1000 запросов", "1000ReqCall"),
		),
	)
	return keyboard
}
