# Start from a Golang base image
FROM golang:1.17 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o url-shortener .

# Start a new stage from scratch
FROM debian:buster-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/url-shortener .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./url-shortener"]

