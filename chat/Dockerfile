FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o websocket-chat-server ./cmd/api/main.go

FROM scratch
COPY --from=builder /app/websocket-chat-server /websocket-chat-server
COPY --from=builder /app/static /static
EXPOSE 8080
CMD ["/websocket-chat-server"]
