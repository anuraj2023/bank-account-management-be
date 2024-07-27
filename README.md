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

## Deployment

Code is already deployed on render platform and it is available here:  
[https://bank-account-management-be.onrender.com/swagger/index.html](https://bank-account-management-be.onrender.com/swagger/index.html)

## Getting Started

There are two ways to run this project:

### 1. Running Locally

1. Clone the repository:
   ```
   git clone https://github.com/anuraj2023/bank-account-management-be.git
   cd bank-account-management-be
   ```

2. Download all dependencies
   ```
   go mod download
   ```

3. Copy the `.env.example` file to `.env` and fill in all the values:
   ```
   cp .env.example .env
   ```

4. Run the main.go file to start the server locally:
   ```
   cd cmd/server
   go run main.go
   ```

### 2. Running with Docker Compose

1. Clone the repository:
   ```
   git clone https://github.com/anuraj2023/bank-account-management-be.git
   cd bank-account-management-be
   ```

2. Build and run the service using Docker Compose, passing all required environment variables:
   ```
   SERVER_PORT=:8080\
   IMMUDB_URL="https://vault.immudb.io/ics/api/v1/ledger"\
   IMMUDB_API_KEY=""\
   docker-compose up
   ```
   Make sure to replace the IMMUDB_API_KEY with your API Key value.

The service will be available at `http://localhost:8080` <br/>
For API docs in local, you can visit this URL: `http://localhost:8080/swagger/index.html`

## Project Structure

- `cmd/server`: Contains the main application entry point
- `internal/api`: Handles HTTP routing and request handling
- `internal/config`: Manages application configuration
- `internal/models`: Defines data models
- `internal/repository`: Implements data storage and retrieval
- `pkg/immudb`: Provides a client for interacting with immudb vault

## To-Do

1. Implementing unit tests

## Note:

- In case you need to generate swagger docs after making some change in API code, run the below command:
   ```
   swag init -g cmd/server/main.go --parseDependency --parseInternal
   ```

## License

This project is licensed under the MIT License.