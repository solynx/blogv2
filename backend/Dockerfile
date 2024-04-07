FROM golang:1.21-alpine AS builder

LABEL alidoan="Van Doan Dev"

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download


RUN apk add --update --no-cache \
		ca-certificates \
		openssl

RUN openssl genpkey -algorithm RSA -out private_key.pem -aes256 -pass pass:sammy_welcome


RUN openssl rsa -in private_key.pem -outform PEM -out private-key.pem -passin pass:sammy_welcome


COPY . .

RUN rm -rf .env

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -ldflags="-s -w" -o server .

FROM alpine:3.15 AS bin

WORKDIR /app

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder /build/server server
COPY --from=builder /build/private-key.pem private-key.pem
CMD ["./server"]