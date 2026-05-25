# Portfolyo API

Backend service for the [Portfolyo](https://portfolyo.luiscabantac.com) mobile app - handles portfolio screenshot generation and management.


![image](https://portfolyo.luiscabantac.com/og.jpg)


## Features

- Portfolio screenshot generation with theme support
- Clerk authentication integration
- Convex database integration

## Tech Stack

- **Language**: Go 1.26.3+
- **HTTP Framework**: Chi
- **Database**: Convex
- **Authentication**: Clerk SDK
- **Screenshot Generation**: Rod

## Getting Started

### Prerequisites

- Go 1.26.3 or higher

### Installation

1. Clone the repository:
```bash
git clone https://github.com/LuisCabantac/go-portfolyo-api.git
cd go-portfolyo-api
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables in a `.env` file:
```env
PORT=8080
APP_ENV=development

# Clerk Authentication
CLERK_SECRET_KEY=<your_clerk_secret_key>
CLERK_PUBLISHABLE_KEY=<your_clerk_publishable_key>

# Convex Cloud Configuration
CONVEX_DEPLOYMENT=<your_convex_deployment_id>
CONVEX_URL=https://<your_project>.convex.cloud
CONVEX_SITE_URL=https://<your_project>.convex.site
CONVEX_DEPLOY_KEY=<your_convex_deploy_key>
```

### Development

Start the development server:
```bash
go run ./cmd
```

The API will be available at `http://localhost:8080`

### Docker

The project includes Docker configuration for containerized deployment.

#### Build and Run with Docker

1. Build the Docker image:
```bash
docker build -t portfolyo-api .
```

2. Run the container:
```bash
docker run -p 8080:8080 --env-file .env portfolyo-api
```

#### Docker Compose

For easier setup with volume management, use Docker Compose:

```bash
docker compose up
```

This will:
- Build the Docker image from the Dockerfile
- Start the service on port 8080
- Load environment variables from `.env`
- Mount a volume for Rod browser cache to persist across container restarts

#### Docker Configuration Details

- **Base Image**: `debian:bookworm-slim` (optimized for production)
- **Build Stage**: Uses `golang:1.26.3-bookworm` for compilation
- **Dependencies**: Includes browser dependencies required by Rod for screenshot generation
- **Volume**: `rod-browser-cache` persists browser cache between runs, improving performance

### Building

Build the binary:
```bash
go build -o portfolyo-api ./cmd/main.go
```

## Project Structure

```
cmd/
├── main.go          # Application entry point
└── api.go           # Router and server configuration

internal/
├── apperrors/       # Error types
├── health/          # Health check service
├── json/            # JSON utilities
├── portfolios/      # Portfolio management
└── screenshot/      # Screenshot utilities
```

## License

MIT License - see the LICENSE file for details.
