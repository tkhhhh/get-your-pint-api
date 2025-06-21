# Go API Server

This project is a simple API server built with Go. It includes a basic health check endpoint to verify that the server is running.

## Project Structure

```
go-api-server
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── handlers
│   │   └── health.go    # Health check handler
│   └── routes
│       └── router.go     # Route setup
├── go.mod               # Module definition and dependencies
└── go.sum               # Dependency checksums
```

## Getting Started

To run the server, navigate to the project directory and use the following command:

```
go run cmd/main.go
```

## Endpoints

- `GET /health`: Returns a simple message indicating that the server is running.

## Build

- GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go
- docker run -p 8080:8080 go-api-server