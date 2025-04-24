# Добавляем Go в PATH
$env:Path += ";C:\Program Files\Go\bin"

# Установка токена
$env:TELEGRAM_BOT_TOKEN = "7655801304:AAEwZ9e-ytBPecjCvuv4c5P9cw_aPJFRD6s"

# Вывод информации
Write-Host "Токен установлен: $env:TELEGRAM_BOT_TOKEN"
Write-Host "Версия Go: $(go version)"

# Убедимся, что зависимости установлены
Write-Host "Устанавливаем зависимости..."
go mod tidy

# Запускаем бота
Write-Host "Запускаем бота..."
go run main.go 