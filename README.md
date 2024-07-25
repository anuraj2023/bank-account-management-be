# Bank Account Management Service

This project is a simple bank account management service built with Go and the Echo framework. It uses immudb for storing account information securely.

## Features

- Store and retrieve account information
- RESTful API for account management
- Docker support for easy deployment

## Prerequisites

- Go 1.16+
- Docker and Docker Compose
- immudb

## Getting Started

There are two ways to run this project:

### 1. Running Locally

1. Clone the repository:
   ```
   git clone https://github.com/anuraj2023/bank-account-management-be.git
   cd bank-account-management-be
   ```

2. Copy the `.env.example` file to `.env` and fill in all the values:
   ```
   cp .env.example .env
   ```

3. Start immudb in a Docker container:
   ```
   docker run -it -d -p 3322:3322 -p 9497:9497 --name immudb codenotary/immudb:latest
   ```

4. Run the main.go file to start the server locally:
   ```
   go run cmd/server/main.go
   ```

### 2. Running with Docker Compose

1. Clone the repository:
   ```
   git clone https://github.com/anuraj2023/bank-account-management-be.git
   cd bank-account-management-be
   ```

2. Build and run the service using Docker Compose, passing all required environment variables:
   ```
   SERVER_ADDRESS=":8080" \
   IMMUDB_ADDRESS="immudb" \
   IMMUDB_USERNAME="immudb" \
   IMMUDB_PASSWORD="immudb" \
   docker-compose up
   ```
   Make sure to replace the values with your actual configuration.

The service will be available at `http://localhost:8080`

## Project Structure

- `cmd/server`: Contains the main application entry point
- `internal/api`: Handles HTTP routing and request handling
- `internal/config`: Manages application configuration
- `internal/models`: Defines data models
- `internal/repository`: Implements data storage and retrieval
- `pkg/immudb`: Provides a client for interacting with immudb

## To-Do

1. Adding Swagger API documentation
2. Implementing unit tests
3. Adding pagination to GET APIs

## License

This project is licensed under the MIT License.