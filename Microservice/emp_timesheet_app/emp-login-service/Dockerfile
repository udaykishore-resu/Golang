# Use the official Golang image as a parent image
FROM golang:1.21.5

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire service directory
COPY . .

# Build the application
# Adjust the path to where your main.go is located
RUN go build -o /app/main ./cmd/main.go

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["/app/main"]
