# Step 1: Use the official Golang image to build the application
FROM golang:1.18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application for the current platform
RUN go build -o websocket-chat-server ./cmd/api/main.go

# Step 2: Use a smaller base image to run the application
FROM alpine:latest

# Set the working directory inside the new image
WORKDIR /app

# Copy the built application from the previous stage
COPY --from=builder /app/websocket-chat-server .

# Copy the static files (e.g., HTML, CSS, JS) to the container
COPY --from=builder /app/clients ./clients

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./websocket-chat-server"]
