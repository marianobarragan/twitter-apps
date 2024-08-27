# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .
WORKDIR ./cmd
# Build the Go app
RUN go build -o users-service .

EXPOSE 8083
# Command to run the executable
CMD ["./users-service"]