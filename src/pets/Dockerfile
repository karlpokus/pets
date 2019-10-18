FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download # deps will be cached

COPY . .

RUN go build -o pets ./cmd/pets

EXPOSE 9013

CMD ["./pets"]
