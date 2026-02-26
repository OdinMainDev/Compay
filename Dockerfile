FROM golang:latest AS builder
WORKDIR /build

ARG VERSION
ARG APP_PATH=./cmd/merchant/main.go

COPY go.mod go.sum ./
COPY efw ./efw
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -race \
    -ldflags="-X 'efw/config.AppVersion=$VERSION'"  \
    -o /compay ${APP_PATH}

FROM debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates curl && apt-get clean && rm -rf /var/lib/apt/lists/*
COPY --from=builder /compay /compay
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /compay /entrypoint.sh
RUN touch .env
RUN mkdir -p /storage/logs
RUN mkdir -p /storage/uploads
ENTRYPOINT ["/entrypoint.sh"]
