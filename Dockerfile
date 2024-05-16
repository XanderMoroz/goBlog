# Собираем первый образ на основе Linux alpine + golang
FROM golang:alpine as builder  

# ENV GO111MODULE=on

# Устанавливаем git
# Чтоб корректно установить зависимости
RUN apk update && apk add --no-cache git

# Устанавливаем рабочую директорию и переходим в нее 
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Устанавливаем зависимости из go.mod
RUN go mod download 

# Копируем все содержимое текущей директории в рабочую (т.е. в /app)
COPY . .

# Компилируем приложение в один бинарный файл
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Собираем второй образ на основе Linux alpine (from scratch)
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию и переходим в нее 
WORKDIR /root/

# Копируем бинарный файл и переменные окружения из первого образа
COPY --from=builder /app/main .
COPY --from=builder /app/.env .       

# Открываем порт 3000 
EXPOSE 3000

# Запускаем бинарник приложения
CMD ["./main"]