# Production API

A Go HTTP API built with [Chi](https://github.com/go-chi/chi) for routing and middleware. Clean, minimal, no magic.

## Quick Start

```bash
go run ./cmd
```

Server starts on `:8080`.

## Endpoints

| Method | Path         | Description |
|--------|--------------|-------------|
| GET    | /healthcheck | Returns "working" |

## Project Structure

```
production_api/
├── cmd/
│   ├── main.go    # Entry point, wiring
│   └── api.go     # Router setup, server config
├── internal/
│   └── products/  # Product domain (handlers, service)
├── go.mod
└── go.sum
```

## Stack

- **Go 1.23**
- **Chi v5** — router + middleware (RequestID, RealIP, Logger, Recoverer, 60s timeout)

## Configuration

Currently hardcoded in `main.go`:
- Address: `:8080`
- DB config: empty (placeholders for future use)

## Requirements

- Go 1.23+
