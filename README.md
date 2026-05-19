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
git clone https://github.com/LuisCabantac/portfolyo-go-api.git
cd portfolyo-go-api
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables in a `.env.local` file:
```env
PORT=8080
APP_ENV=development
CONVEX_URL=<your_convex_url>
CONVEX_DEPLOY_KEY=<your_convex_deploy_key>
CLERK_SECRET_KEY=<your_clerk_secret_key>
```

### Development

Start the development server:
```bash
go run ./cmd
```

The API will be available at `http://localhost:8080`

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
