package main

import (
	"log"
	"os"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"

	"github.com/joho/godotenv"
	
	"telegrambot/utils"
)

func main() {
	// Пробуем загрузить переменные окружения из .env файла
	// Игнорируем ошибку, если файл не найден - можно использовать export
	_ = godotenv.Load()

	// Получаем токен бота из переменных окружения
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN не найден в переменных окружения. Установите через export BOT_TOKEN=your_token или создайте .env файл")
	}

	// Проверяем наличие API ключа OpenAI
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Println("OPENAI_API_KEY не найден в переменных окружения. Команда /story не будет работать.")
	}

	// Создаем нового бота
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatal("Ошибка при создании бота: ", err)
	}

	// Создаем диспетчер для обработки обновлений
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("Ошибка при обработке обновления:", err.Error())
			return ext.DispatcherActionNoop
		},
	})

	// Добавляем обработчик для команды /story
	dispatcher.AddHandler(handlers.NewCommand("story", handleStoryCommand))

	// Добавляем обработчик для всех остальных сообщений
	dispatcher.AddHandler(handlers.NewMessage(message.All, handleMessage))

	// Запускаем получение обновлений
	updater := ext.NewUpdater(dispatcher, nil)
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
	})
	if err != nil {
		log.Fatal("Ошибка при запуске получения обновлений: ", err)
	}

	log.Printf("Бот @%s успешно запущен", bot.User.Username)

	// Ожидаем завершения работы бота
	updater.Idle()
}

// Обработчик команды /story
func handleStoryCommand(b *gotgbot.Bot, ctx *ext.Context) error {
	// Проверяем наличие API ключа OpenAI
	if os.Getenv("OPENAI_API_KEY") == "" {
		_, err := ctx.EffectiveMessage.Reply(b, "Для генерации историй необходимо настроить OPENAI_API_KEY в переменных окружения", nil)
		return err
	}
	
	// Получаем стиль из аргументов команды, по умолчанию "sci-fi"
	style := "sci-fi"
	if len(ctx.Args()) > 0 {
		style = strings.Join(ctx.Args(), " ")
	}
	
	// Отправляем сообщение, что бот генерирует историю
	_, err := ctx.EffectiveMessage.Reply(b, "Генерирую историю в стиле "+style+"...", nil)
	if err != nil {
		return err
	}
	
	// Генерируем историю
	story, err := utils.GenerateStory(style, 400)
	if err != nil {
		_, err = ctx.EffectiveMessage.Reply(b, "Ошибка при генерации истории: "+err.Error(), nil)
		return err
	}
	
	// Отправляем готовую историю пользователю
	_, err = ctx.EffectiveMessage.Reply(b, story, nil)
	return err
}

// Обработчик всех входящих сообщений
func handleMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	// Пропускаем команды
	if ctx.EffectiveMessage.Text != "" && ctx.EffectiveMessage.Text[0] == '/' {
		return nil
	}
	
	// Получаем текст сообщения
	messageText := ctx.EffectiveMessage.Text
	
	// Переворачиваем текст с помощью функции из utils
	reversedText := utils.ReverseString(messageText)
	
	// Отправляем перевернутый текст обратно пользователю
	_, err := ctx.EffectiveMessage.Reply(b, reversedText, nil)
	if err != nil {
		return err
	}
	
	return nil
} 