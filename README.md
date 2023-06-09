## Go Telegram Bot
👋 Привет! Это мой Pet-Проект. Настоящее название этого бота - CatAI-Bot.
Он умеет обращаться к API ChatGPT-3 (Модель давиничи-3) и имеет БД в виде MYSQL.

Бот написан на языке программирования Go с использованием таких библиотек (фреймворков) как:
- github.com/go-sql-driver/mysql (Драйвер базы данных MYSQL)
- github.com/go-telegram-bot-api/telegram-bot-api/v5 (API-Клиент для Telegram-ботов)
- github.com/sashabaranov/go-openai (API-Клиент Open-AI)

*Предупреждение: Это мой первый проект на языке Go, он не закончен и требует прохождения тестирований и оптимизации.*
# Установка
Перед началом работы бота Вам нужно установить язык Go стандарта 1.20 или выше.
После клонирования и настройки проекта Вам потребуется установить нужные модули. Вы можете это сделать при помощи пакетного менеджера:

*Установка драйвера базы данных:*

    go get -u github.com/go-sql-driver/mysql
*Установка API клиента для работы с Telegram:*

    github.com/go-telegram-bot-api/telegram-bot-api

*Установка API клиента для работы с Open-AI:*

    github.com/sashabaranov/go-openai
# Настройка

После установки проекта Вам потребуется настроить конфиг. 
*p.s. Я не стал переводить config на toml, поскольку он еще не закончен.*

Конфиг располагается по пути:

    configs/config.go
В конфиге Вам потребуется установить Ваш токен от бота, open-ai и данные БД. Все эти данные нужно внести в соответствующие константы:

        package  configs
        
        // Telegram-Token
        const TELEGRAM_API_TOKEN =  "TELEGRAM_BOT_TOKEN"
    		
    	// OPENAI-API-Token
    	const OPEN_AI_TOKEN =  "OpenAI-API-Token"
    	
	    // DEBUG MODE
	    const DEBUG_MODE =  false
	    
	    // TIMEOUT
	    const TIMEOUT =  60
	    
	    // DataBase connection config
	    const PG_USER =  "DB_USER"
	    const PG_PASS =  "DB_PASS"
	    const PG_DB =  "DB_DB"
	    const PG_HOST =  "DB_HOST"
# Основной функционал

- Написание текста-запроса к боту. Бот обработает входные данные через ChatGPT.
- Команда /settings позволяет узнать различные значения пользователя
- Команда /start запускает бота и регистрирует пользователя при надобности

# Авторство
🐈 Created By Catborisovv (c) 2020-2024


**Enjoy!**
