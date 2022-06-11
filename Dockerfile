FROM golang:1.18-alpine

WORKDIR /app

RUN apk add git build-base

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags musl

CMD ./tereus-transpiler-c-go
