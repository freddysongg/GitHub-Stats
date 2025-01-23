# Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app

# Show initial directory structure
RUN echo "Initial directory:" && ls -la

# Copy and install dependencies
COPY go.mod go.sum ./
RUN echo "After copying go.mod/go.sum:" && ls -la
RUN go mod download

# Copy source code
COPY . .
RUN echo "After copying source code:" && ls -la

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o github-dashboard .
RUN echo "After build:" && ls -la

# Runtime stage
FROM alpine:latest
WORKDIR /app

# Copy built binary and verify
COPY --from=builder /app/github-dashboard /app/github-dashboard
COPY .env .

# Verify file exists and set permissions
RUN echo "Runtime stage files:" && ls -la /app && \
    echo "Binary details:" && \
    file /app/github-dashboard && \
    ldd /app/github-dashboard || true

EXPOSE 8080
CMD ["./github-dashboard"]
