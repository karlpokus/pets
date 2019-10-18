FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download # deps will be cached

COPY . .

ARG VERSION
RUN export VERSION=`cat version`
RUN go build -ldflags "-X pets.Version=$VERSION" -o pets ./cmd/pets

EXPOSE 9013

CMD ["./pets"]
