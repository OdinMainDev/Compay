FROM golang:latest AS builder
WORKDIR /build

ARG VERSION
ARG APP_PATH=./cmd/merchant/main.go
ARG APP_NAME=compay

COPY go.mod go.sum ./
COPY efw ./efw
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -race \
    -ldflags="-X 'efw/config.AppVersion=$VERSION'"  \
    -o /${APP_NAME} ${APP_PATH}

FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*
COPY --from=builder /compay /compay
COPY files files/
RUN chmod +x /compay
RUN touch .env
RUN mkdir -p /storage/logs
ENTRYPOINT ["/bin/sh", "-c", "logfile=\"/storage/logs/docker_$(date +%d_%m_%Y__%H_%M).log\" && /compay > \"$logfile\" 2>&1"]