# Используем актуальный базовый образ с Go
FROM golang:1.22.8 AS builder

WORKDIR /app

# Копируем файлы go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./

# Загружаем модули
RUN go mod download

# Копируем весь исходный код
COPY . .

# Сборка Go-приложения
RUN go build -o cfupb ./cmd/main.go

# Создаем финальный легкий контейнер
FROM debian:bookworm-slim

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированное приложение
COPY --from=builder /app/cfupb ./

# Копируем конфигурационный файл в контейнер
COPY schema/* ./schema/
COPY .env ./

# Указываем команду по умолчанию для запуска приложения
ENTRYPOINT ["./cfupb"]
