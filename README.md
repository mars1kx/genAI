![image](https://github.com/user-attachments/assets/f9710604-af34-412d-a501-416659f0b1da)
# Telegram Бот на Go

Простой Telegram бот, который отвечает "hello" на любое сообщение.

## Установка

1. Убедитесь, что у вас установлен Go (версия 1.16 или выше)
2. Клонируйте репозиторий
3. Установите зависимости:
   ```
   go mod tidy
   ```

## Настройка

1. Создайте бота в Telegram через [@BotFather](https://t.me/BotFather) и получите токен
2. Установите токен как переменную окружения:

   - Windows:
   ```
   set TELEGRAM_BOT_TOKEN=ваш_токен
   ```
   
   - Linux/Mac:
   ```
   export TELEGRAM_BOT_TOKEN=ваш_токен
   ```

## Запуск

```
go run main.go
```

После запуска бот будет отвечать "hello" на все ваши сообщения. 
