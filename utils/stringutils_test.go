package utils

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	// Подготавливаем тестовые случаи в виде пар входных и ожидаемых значений
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Пустая строка",
			input:    "",
			expected: "",
		},
		{
			name:     "Односимвольная строка",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Обычная строка",
			input:    "привет",
			expected: "тевирп",
		},
		{
			name:     "Палиндром",
			input:    "шалаш",
			expected: "шалаш",
		},
		{
			name:     "Строка с пробелами",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "Строка с цифрами",
			input:    "123456",
			expected: "654321",
		},
		{
			name:     "Строка со специальными символами",
			input:    "!@#$%^&*()",
			expected: ")(*&^%$#@!",
		},
		{
			name:     "Кириллические символы",
			input:    "абвгдеёжзий",
			expected: "йизжёедгвба",
		},
		{
			name:     "Эмодзи",
			input:    "😀🙂😊",
			expected: "😊🙂😀",
		},
	}

	// Запускаем все тестовые случаи
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ReverseString(tc.input)
			if result != tc.expected {
				t.Errorf("ReverseString(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
} 