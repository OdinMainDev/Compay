FROM golang:latest AS builder
WORKDIR /build
ARG VERSION

COPY go.mod go.sum ./
RUN go mod download
#RUN go mod vendor

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build ./cmd/merchant/main.go --race \
    -ldflags="-X 'effectivesolution/common/constants.Version=$VERSION'"  \
    -o /effectivesolution

FROM debian:stable-slim
RUN apt-get update && apt-get install -y ca-certificates && apt-get clean && rm -rf /var/lib/apt/lists/*
COPY --from=builder /effectivesolution /effectivesolution
COPY files files/
RUN chmod +x /effectivesolution
RUN touch .env
RUN mkdir -p /storage/logs
ENTRYPOINT ["/bin/sh", "-c", "logfile=\"/storage/logs/docker_$(date +%d_%m_%Y__%H_%M).log\" && /effectivesolution > \"$logfile\" 2>&1"]