# Start from a newer golang base image
FROM golang:1.21-alpine AS builder

# Set the Current Working Directory as app
WORKDIR /app

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Generate Swagger documentation
RUN swag init -g cmd/server/main.go --parseDependency --parseInternal

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./cmd/server

# Optimizing for a lean docker image which has built binary and minimal packages
FROM alpine:latest 

# Set the Current Working Directory as root
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
