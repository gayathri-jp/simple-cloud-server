# Use a Golang base image to build the app
FROM golang:1.18-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o simpletimeservice .

# Start a new stage from a minimal image
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the binary from the previous stage
COPY --from=builder /app/simpletimeservice .

# Expose port 8080
EXPOSE 8080

# Run the executable
CMD ["./simpletimeservice"]
