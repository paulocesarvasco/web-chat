FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o auth_service ./cmd/api/main.go

FROM scratch
COPY --from=builder /app/auth_service .
COPY --from=builder /app/.env .
EXPOSE 8081
CMD ["./auth_service"]
