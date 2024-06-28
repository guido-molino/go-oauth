FROM golang:1.22.1

ARG GIN_MODE=release
ARG PORT=8080

ENV GIN_MODE=$GIN_MODE
ENV PORT=$PORT

WORKDIR /app

COPY go.mod go.sum ./

RUN go install github.com/mitranim/gow@latest
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./...

EXPOSE $PORT