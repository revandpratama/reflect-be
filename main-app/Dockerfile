# Use the official Golang image to build the Go app
FROM golang:1.24.2-alpine AS builder

WORKDIR /main-app

COPY . .

# Install dependencies  
RUN go mod download
RUN go mod tidy
RUN go build -o main .

# Use a minimal base image for running the Go app
FROM alpine:latest

WORKDIR /root

RUN apk add --no-cache libc6-compat

# Copy the compiled Go binary from the builder image
COPY --from=builder /main-app/main .

COPY --from=builder /main-app/.env /root/.env

# Expose the port the app will run on
EXPOSE 8080

# Run the Go binary
CMD ["./main"]