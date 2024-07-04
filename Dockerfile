ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

# Set the working directory to the location of main.go
WORKDIR /usr/src/app/cmd/api

# Build the Go application
RUN go build -v -o /run-app .

FROM debian:bookworm

# Install CA certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the built binary from the builder stage
COPY --from=builder /run-app /usr/local/bin/

# Set environment variables for AWS
ENV AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID}
ENV AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY}

# Set the command to run the built binary
CMD ["run-app"]