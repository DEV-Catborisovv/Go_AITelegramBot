// TelegramBot-AI. Main-Source file
// Created by Catborisovv (c) 2020-2024

package main

import (
	"telegramBot/internal/app"
	"telegramBot/internal/app/database"
)

func main() {
	database.ConnectToDataBase()
	app.StartBot()
}
