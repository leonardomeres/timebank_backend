# Use the official Go image with version 1.23.1
FROM golang:1.23.1-alpine

# Set working directory
WORKDIR /app

# Install git for modules (optional, but often required)
RUN apk add --no-cache git

# Copy go module files first (for caching)
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Expose the API port
EXPOSE 8080

# Run the application
CMD ["go", "run", "cmd/main.go"]