# GitHub Dashboard Backend

This is the backend service for the GitHub Stats Dashboard, providing real-time GitHub data through REST APIs and WebSocket connections.

## Features

- GitHub API integration for user and repository statistics
- Real-time updates through Kafka and WebSocket
- Caching for improved performance
- Containerized deployment with Docker

## Prerequisites

- Docker and Docker Compose
- Go 1.21+
- GitHub Personal Access Token

## Setup

1. Copy `.env.example` to `.env` and configure the environment variables:
   ```bash
   cp .env.example .env
   ```

2. Build and start the services:
   ```bash
   docker-compose up --build
   ```

3. The backend will be available at `http://localhost:8080`

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /profile/{username}` - Get GitHub user profile
- `GET /repositories/{username}` - Get user repositories
- `GET /languages/{username}` - Get top programming languages
- `GET /contributions/{username}` - Get contribution statistics

## WebSocket

Connect to `ws://localhost:8080/ws` for real-time updates on:
- New repository stars
- New issues
- New pull requests

## Development

To run locally without Docker:
```bash
go run main.go
```

## Testing

Run tests with:
```bash
go test ./...
