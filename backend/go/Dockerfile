FROM golang:latest

WORKDIR /app

COPY . .

RUN go build ./cmd/root/main.go

EXPOSE 8080

CMD ["./main"]
