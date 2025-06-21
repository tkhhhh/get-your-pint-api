# Use a minimal base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the pre-built binary into the container
COPY main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]