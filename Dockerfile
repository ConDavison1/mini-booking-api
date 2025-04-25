# syntax=docker/dockerfile:1

# 1. Start from official Go image
FROM golang:1.22

# 2. Set working directory inside the container
WORKDIR /app

# 3. Copy Go files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# 4. Build the Go app
RUN go build -o main ./cmd

# 5. Expose the port and run the app
EXPOSE 3000
CMD ["/app/main"]
