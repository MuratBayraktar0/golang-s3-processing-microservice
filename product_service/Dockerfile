# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Murat Bayraktar"

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install the Go dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application without setting environment variable
RUN go build -o app .

# Set the entry point for the container
ENTRYPOINT ["./app"]