package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// GenerateStory создает короткую историю в указанном стиле с ограничением по количеству символов
func GenerateStory(style string, maxChars int) (string, error) {
	// Получаем ключ API из переменных окружения
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY не найден в переменных окружения")
	}

	// Создаем клиент OpenAI с ключом API
	client := openai.NewClient(option.WithAPIKey(apiKey))
	ctx := context.Background()

	// Формируем запрос на генерацию истории
	prompt := fmt.Sprintf("Напиши короткую %s историю на максимум %d символов. История должна быть интересной, с неожиданным поворотом.", style, maxChars)

	// Настраиваем параметры запроса
	params := openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: openai.ChatModelGPT3_5Turbo,
	}

	// Выполняем запрос к API
	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("ошибка при обращении к API OpenAI: %w", err)
	}

	// Возвращаем сгенерированную историю
	story := completion.Choices[0].Message.Content
	return story, nil
} 