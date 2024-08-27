# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Build
# RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
WORKDIR ./cmd

RUN go build -o timeline-subscriber .

# Run
CMD ["./timeline-subscriber"]