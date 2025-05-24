# Go-URL-Shortner

A simple URL shortener service built with Go and Gin.

## Features

- Shorten long URLs to short, shareable links
- Redirect short URLs to the original URLs
- Rate limiting to prevent abuse
- Request tracing with trace IDs in logs and responses
- URL reachability validation before shortening

## Project Structure

```
cmd/                # Application entry point
internal/
  handler/          # HTTP handlers
  middleware/       # Gin middleware (logging, tracing)
  models/           # Request/response models
  routes/           # Route definitions
  service/          # Business logic for URL shortening
  utils/            # Utility functions (URL validation)
pkg/
  common/           # Shared context keys, logger helpers
  config/           # Environment variable config
```

## Getting Started

### Prerequisites

- Go 1.24+
- Git

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/Go-URL-Shortner.git
   cd Go-URL-Shortner
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

### Configuration

Set the following environment variable (optional):

- `APP_DOMAIN`: The domain for generated short URLs (default: `localhost:8080`)

You can create a `.env` file or set the variable in your shell.

### Running the Application

```sh
go run cmd/main.go
```

The server will start on `localhost:8080`.

## API Endpoints

### Shorten a URL

- **POST** `/short`
- **Headers:** `traceId: <your-trace-id>`
- **Body:**
  ```json
  {
    "url": "https://example.com"
  }
  ```
- **Response:**
  ```json
  {
    "short_url": "localhost:8080/abc123"
  }
  ```

### Redirect to Original URL

- **GET** `/:shortKey`
- **Headers:** `traceId: <any-value>`
- **Response:** Redirects to the original URL.

## Rate Limiting

- Maximum 5 requests per 10 seconds per IP.
- Exceeding the limit returns HTTP 429.

## Logging & Tracing

- Each request is logged with a trace ID.
- The trace ID is required in the `traceId` header (except for redirects, where it uses the short key).
