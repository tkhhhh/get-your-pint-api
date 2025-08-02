# Use a minimal base image
FROM golang:1.24

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the Go application
RUN go build -o main ./cmd/main.go 

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
