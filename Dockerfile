FROM golang:1.18-alpine

WORKDIR /app

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build

CMD ./tereus-remixer-c-go
