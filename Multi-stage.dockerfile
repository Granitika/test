# Используем официальный образ Golang для сборки
FROM golang:1.21 AS builder

# Устанавливаем рабочую директорию внутри образа
WORKDIR /app

# Копируем исходный код внутрь образа
COPY *.go ./
RUN go mod init MSSQL-test && go mod tidy

# Компилируем исходный код в исполняемый файл
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Создаем отдельный образ для запуска приложения
FROM alpine:latest

# Устанавливаем рабочую директорию внутри образа
WORKDIR /app

RUN echo "CTF{asdljka}" > /secret.txt


# Копируем скомпилированный исполняемый файл из предыдущего образа
COPY --from=builder /app/main.go .
COPY --from=builder /app/main .

RUN chmod 555 /app/main /secret.txt

USER nobody 

# Запускаем скомпилированное приложение
CMD ["./main"]
