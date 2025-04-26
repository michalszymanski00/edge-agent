FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -ldflags="-extldflags=-static" -o agent main.go

# Now use small image
FROM alpine:3.20

# Install CA certs in container
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/agent /agent

ENTRYPOINT ["/agent"]
