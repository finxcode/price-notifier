# syntax=docker/dockerfile:1

FROM golang:1.17.6-alpine AS builder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE=on

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o divergence ./cmd

FROM alpine

WORKDIR /app

COPY --from=builder /build /app

CMD ["./divergence"]