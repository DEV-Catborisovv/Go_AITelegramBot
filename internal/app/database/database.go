// DB-Client
// Created by Catborisovv (c) 2020-2024

// Пакет базы данных
package database

import (
	"database/sql"
	"fmt"
	"telegramBot/configs"

	_ "github.com/go-sql-driver/mysql"
)

type SData struct {
	ID       uint
	Chat_id  int64
	Username string
	Requests uint
	Admin    uint
}

// Константа в которой содержатся значения из нашего конфига
const connectData = configs.PG_USER + ":" + configs.PG_PASS + "@tcp(" + configs.PG_HOST + ")" + "/" + configs.PG_DB

// Функция подключения к базе данных
func ConnectToDataBase() *sql.DB {
	db, err := sql.Open("mysql", connectData)
	if err != nil {
		panic(err)
	}
	return db
}

// Функция вставки новых данных в таблицу бд
func InsertData(Data string) error {
	db := ConnectToDataBase()
	insert, err := db.Query(Data)
	if err != nil {
		return err
	}

	// Закрываем сессию подключения к бд
	defer db.Close()
	defer insert.Close()

	return nil
}

// Метод получения запросов пользователя
func GetUserRequest(UserChat int64) (error, int) {
	// Новая сессия БД, получение данных из запроса
	db := ConnectToDataBase()

	// Запрос получения запросов пользователя из таблицы
	reqSelect := fmt.Sprintf("SELECT `requests` FROM `users` WHERE `chat_id` = '%v';", UserChat)
	res, err := db.Query(reqSelect)
	// Обработка ошибки при получении запроса
	if err != nil {
		return err, 0
	}

	var result int

	for res.Next() {
		err = res.Scan(&result)
		if err != nil {
			return err, 0
		}
	}

	// Завершение сессии базы данных
	defer db.Close()
	defer res.Close()
	return nil, result
}

// Функция получения данных из таблицы
// Принимает на вход SQL-Запрос, записывает данные в структуру, после чего возвращает указатель на структуру
func SelectData(Data string) (error, *SData) {
	// Создание новой сессии подключения к базе данных и выполнение запроса из аргумента функции
	db := ConnectToDataBase()
	res, err := db.Query(Data)

	// Проверка успешности выполнения запроса
	if err != nil {
		return err, nil
	}

	var user SData
	// Получение результата из таблицы в БД
	for res.Next() {
		err = res.Scan(&user.ID, &user.Chat_id, &user.Username, &user.Requests, &user.Admin)
		if err != nil {
			return err, nil
		}
	}

	// Завершение сессии БД
	defer db.Close()
	defer res.Close()
	return nil, &user
}
