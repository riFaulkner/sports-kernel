# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY sk-serve/go.mod ./
COPY sk-serve/go.sum ./

RUN go mod download

COPY sk-serve/ ./

RUN go build -o main .

EXPOSE 8080

CMD [ "/app/main"]