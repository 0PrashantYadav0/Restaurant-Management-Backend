# Use an official Golang image as the base
FROM golang:alpine

# Set the working directory to /app
WORKDIR /app

# Copy the Go module dependencies
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o main main.go

# Expose the port
EXPOSE 8000

# Set the default command to run the application
CMD ["./main"]