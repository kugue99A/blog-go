FROM golang:1.22.7

WORKDIR /api
COPY . /api

RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go mod tidy

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
