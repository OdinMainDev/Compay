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

ARG APP_NAME=compay
ARG VERSION
ENV APP_NAME=${APP_NAME}
ENV VERSION=${VERSION}

LABEL version="${VERSION}"

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*
COPY --from=builder /${APP_NAME} /app
RUN chmod +x /app
RUN touch .env
RUN mkdir -p /resources/logs
ENTRYPOINT ["/app"]