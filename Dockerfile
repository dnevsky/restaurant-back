# Используем базовый образ для Go
FROM golang:1.24.0-alpine

# Создадим директорию
RUN mkdir /app

# Скопируем всё в директорию
ADD . /app/

# Установим рабочей папкой директорию
WORKDIR /app

# Сертификаты и таймзоны
RUN apk update && apk add ca-certificates && apk add tzdata

# Соберём приложение
RUN go build -buildvcs=false -o restaurant ./cmd/app

FROM alpine:3.17.2

WORKDIR /app/

RUN apk update && apk add ca-certificates && apk add tzdata

COPY --from=0  /app/restaurant /app/restaurant

EXPOSE 8000

ENTRYPOINT /app/restaurant
