# Бот для переворачивания сообщений

Этот Telegram бот написан на языке Go с использованием библиотеки [gotgbot](https://github.com/PaulSonOfLars/gotgbot).

## Функциональность

Бот умеет:

- Переворачивать любое входящее сообщение (на "привет" отвечает "тевирп" и т.д.)
- Генерировать короткие истории по команде `/story` с использованием OpenAI API

### Команды

- `/story` - генерирует фантастическую (sci-fi) историю
- `/story [стиль]` - генерирует историю в указанном стиле (например, `/story фэнтези`, `/story детектив`, и т.д.)

## Структура проекта

- `main.go` - основной файл бота
- `utils/stringutils.go` - вспомогательные функции для работы со строками
- `utils/stringutils_test.go` - unit-тесты для функций работы со строками
- `utils/openai_utils.go` - функции для работы с OpenAI API
- `.env` - файл с переменными окружения (не включен в репозиторий, опционально)
- `.env.example` - пример файла с переменными окружения

## Настройка и запуск

### Необходимые переменные окружения

- `BOT_TOKEN` - токен вашего Telegram бота, полученный от @BotFather
- `OPENAI_API_KEY` - ключ API OpenAI для генерации историй (команда `/story`)

### Вариант 1: Использование .env файла

1. Клонировать репозиторий
2. Создать файл `.env` и добавить в него необходимые переменные:
   ```
   BOT_TOKEN=your_bot_token_here
   OPENAI_API_KEY=your_openai_api_key_here
   ```
3. Установить зависимости:
   ```
   go mod tidy
   ```
4. Запустить бота:
   ```
   go run main.go
   ```

### Вариант 2: Использование переменных окружения

1. Клонировать репозиторий
2. Установить переменные окружения:
   ```
   export BOT_TOKEN=your_bot_token_here
   export OPENAI_API_KEY=your_openai_api_key_here
   ```
3. Установить зависимости:
   ```
   go mod tidy
   ```
4. Запустить бота:
   ```
   go run main.go
   ```

## Тестирование

Для запуска unit-тестов выполните:

```
go test -v ./...
```

Или для запуска только тестов в директории utils:

```
cd utils && go test -v
```

## Зависимости

- [gotgbot/v2](https://github.com/PaulSonOfLars/gotgbot) - библиотека для работы с Telegram Bot API
- [godotenv](https://github.com/joho/godotenv) - для загрузки переменных окружения из .env файла
- [openai-go](https://github.com/openai/openai-go) - официальный клиент OpenAI API для Go
.