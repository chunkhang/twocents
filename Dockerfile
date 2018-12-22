# BUILD
FROM golang:1.11.2-alpine AS builder

# Install git for fetching dependencies
RUN apk update && apk add --no-cache git

# Copy source files
COPY . /app
WORKDIR /app

# Fetch dependencies
RUN go get -d -v

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /app/main

# EXECUTABLE
FROM scratch

# Copy binary
COPY --from=builder /app/main /app/main

# Run binary
ENTRYPOINT ["/app/main"]
