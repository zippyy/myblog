FROM golang:1.25-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/techrelay-terminal ./cmd/techrelay-terminal

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S techrelay \
    && adduser -S -G techrelay -h /app techrelay \
    && mkdir -p /data \
    && chown techrelay:techrelay /data

COPY --from=build /out/techrelay-terminal /usr/local/bin/techrelay-terminal

USER techrelay
WORKDIR /app

ENV TECHRELAY_FEED_URL=https://techrelay.xyz/index.json \
    TECHRELAY_SITE_URL=https://techrelay.xyz \
    TECHRELAY_LISTEN_ADDR=0.0.0.0:23234 \
    TECHRELAY_HEALTH_ADDR=0.0.0.0:8080 \
    TECHRELAY_HOST_KEY=/data/ssh_host_ed25519 \
    TECHRELAY_REFRESH_INTERVAL=15m

EXPOSE 23234 8080
VOLUME ["/data"]

ENTRYPOINT ["/usr/local/bin/techrelay-terminal"]
