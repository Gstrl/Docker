# Первый этап: сборка
FROM golang:1.22-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /celen

# необходимые утилиты
RUN apk --no-cache add bash git make gcc gettext musl-dev

# Копируем файлы go.mod и go.sum и устанавливаем зависимости
COPY ["./app/go.mod", "./app/go.sum", "./"]
RUN go mod download

# Копируем исходный код и собираем приложение
COPY app ./
RUN go build -o ./bin/app cmd/main.go

# Второй этап: создание минимального образа
FROM alpine AS runner


# Копируем только собранный бинарный файл из первого этапа
COPY --from=builder /celen/bin/app /

COPY ./app/config/local.yaml ./local.yaml

# Указываем команду для запуска приложения
CMD ["/app"]