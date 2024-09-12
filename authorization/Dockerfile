FROM golang:1.23-alpine AS builder
RUN apk --no-cache add curl
RUN apk --no-cache add bash
WORKDIR /app
COPY . .
# RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o auth_service ./cmd/api/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/auth_service .
EXPOSE 8081
CMD ["./auth_service"]
