FROM golang:latest AS builder
WORKDIR /build

ARG VERSION
ARG APP_PATH=./cmd/merchant/main.go
ARG APP_NAME=effectivesolution

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build ${APP_PATH} --race \
    -ldflags="-X 'effectivesolution/common/constants.Version=$VERSION'"  \
    -o /${APP_NAME}

FROM debian:stable-slim

ARG APP_NAME=effectivesolution
ENV APP_NAME=${APP_NAME}

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*
COPY --from=builder /${APP_NAME} /app
RUN chmod +x /app
RUN touch .env
RUN mkdir -p /storage/logs
ENTRYPOINT ["/bin/sh", "-c", "logfile=\"/storage/logs/docker_$(date +%d_%m_%Y__%H_%M).log\" && /app > \"$logfile\" 2>&1"]