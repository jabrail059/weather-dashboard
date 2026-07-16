# Weather Dashboard

[![Check](https://github.com/jabrail059/weather-dashboard/actions/workflows/check.yml/badge.svg)](https://github.com/jabrail059/weather-dashboard/actions/workflows/check.yml)
[![Build](https://github.com/jabrail059/weather-dashboard/actions/workflows/build.yml/badge.svg)](https://github.com/jabrail059/weather-dashboard/actions/workflows/build.yml)
[![Go](https://img.shields.io/badge/Go-1.25.6-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Docker](https://img.shields.io/badge/Docker-ready-2496ED?logo=docker&logoColor=white)](https://www.docker.com/)

A server-rendered weather dashboard built with Go, Chi, HTMX, SQLite, and Redis.

The application provides current weather conditions, a seven-day forecast, hourly forecast details, favorite cities, and recent search history. Weather data comes from the Open-Meteo API, while Go HTML templates and HTMX provide a responsive interface without a client-side application framework.

## Screenshots

### Main dashboard

![Main dashboard](docs/screenshots/home.png)

### Hourly forecast

![Hourly forecast](docs/screenshots/hourly.png)

## Features

- Search for weather by city
- View current weather conditions
- View a seven-day forecast
- Open hourly forecast details for a selected day
- Add and remove favorite cities
- Review recent search history
- Use an automatically detected city suggestion

## Technical Highlights

- Server-side rendering with Go HTML templates
- HTTP routing with Chi
- Partial page updates with HTMX
- Weather data retrieval from Open-Meteo
- Redis caching for external API responses
- SQLite persistence for favorites and search history
- Cookie-based user sessions
- Graceful shutdown of the HTTP server and storage connections
- Docker Compose environment for the application and Redis
- Automated linting, tests, and Docker image builds with GitHub Actions

## Tech Stack

| Area | Technology |
|---|---|
| Language | Go 1.25.6 |
| HTTP router | Chi |
| Rendering | Go HTML templates, HTMX |
| Frontend | HTML, CSS |
| Database | SQLite |
| Cache | Redis |
| Weather API | Open-Meteo |
| Testing | Go testing package, Testify |
| Infrastructure | Docker, Docker Compose |
| CI | GitHub Actions, golangci-lint |

## Architecture

The project follows a layered structure that keeps HTTP handling, application logic, persistence, external API access, and rendering separate.

```text
Browser
   │
   ▼
Chi router
   │
   ▼
HTTP handlers
   │
   ▼
Service layer
   ├── Open-Meteo API
   ├── Redis cache
   └── SQLite storage
   │
   ▼
Go HTML templates and HTMX responses
```

Main package responsibilities:

- `internal/handlers` — HTTP request handling and response preparation
- `internal/service` — application logic and weather data processing
- `internal/storage/sqlite` — favorites and search history persistence
- `internal/storage/redis` — weather response caching
- `internal/session` — cookie-based session handling
- `internal/view` — template rendering
- `internal/server` — route and static file configuration
- `internal/app` — application dependency initialization

## Project Structure

```text
.
├── cmd/app/                 # Application entry point
├── docs/screenshots/        # README screenshots
├── internal/
│   ├── app/                 # Dependency initialization
│   ├── config/              # Environment configuration
│   ├── handlers/            # HTTP handlers
│   ├── models/              # Application models
│   ├── server/              # Router configuration
│   ├── service/             # Application and weather logic
│   ├── session/             # User session handling
│   ├── storage/
│   │   ├── redis/           # Redis cache
│   │   └── sqlite/          # SQLite storage
│   ├── view/                # Template rendering
│   └── weather/             # Weather code descriptions
├── migrations/              # SQLite migrations
├── static/                  # CSS and static assets
├── templates/               # Go HTML templates
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

## Quick Start

### Requirements

- Docker
- Docker Compose
- GNU Make for the recommended commands

### 1. Clone the repository

```bash
git clone https://github.com/jabrail059/weather-dashboard.git
cd weather-dashboard
```

### 2. Create the environment file

```bash
cp .env.example .env
```

The provided example is configured for Docker Compose:

```env
APP_PORT=8081
SQLITE_PATH=/data/weather.db
REDIS_ADDR=redis:6379
REDIS_PASSWORD=
REDIS_DB=0
```

### 3. Build and start the application

The Makefile is the recommended command interface for this project:

```bash
make build
make up
```

Open the application at:

```text
http://localhost:8081
```

SQLite data is persisted in the local `data/` directory.

View application logs:

```bash
make logs
```

Stop the application:

```bash
make down
```

### Without Make

Use the equivalent Docker Compose command:

```bash
docker compose up --build -d
```

To view logs and stop the services:

```bash
docker compose logs -f app
docker compose down
```

## Local Development

### Requirements

- Go 1.25.6 or newer
- Redis

For local development outside Docker, update `.env` to use local paths and the local Redis address:

```env
APP_PORT=8081
SQLITE_PATH=./data/weather.db
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

Create the data directory and download dependencies:

```bash
mkdir -p data
go mod download
```

Run the application through the Makefile:

```bash
make run
```

Equivalent direct command:

```bash
go run ./cmd/app/
```

The application will be available at `http://localhost:8081`.

## Makefile Commands

The Makefile provides a small, consistent interface for common development and Docker tasks.

| Command | Purpose | Equivalent command |
|---|---|---|
| `make build` | Build the application image | `docker compose build app` |
| `make up` | Start the services in the background | `docker compose up -d` |
| `make down` | Stop and remove the services | `docker compose down` |
| `make logs` | Follow application logs | `docker compose logs -f app` |
| `make restart` | Restart the application container | `docker compose restart app` |
| `make run` | Run the application locally | `go run ./cmd/app/` |
| `make fmt` | Format Go source files | `go fmt ./...` |
| `make vet` | Run static checks with `go vet` | `go vet ./...` |
| `make test` | Run the complete test suite | `go test ./...` |

Using Make is recommended because it keeps project commands short and consistent. The equivalent commands are documented so the project remains easy to use on systems where `make` is unavailable.

## Configuration

| Variable | Description | Docker Compose value |
|---|---|---|
| `APP_PORT` | HTTP port used by the application. Keep it at `8081` with the current Docker Compose configuration | `8081` |
| `SQLITE_PATH` | Path to the SQLite database file | `/data/weather.db` |
| `REDIS_ADDR` | Redis server address | `redis:6379` |
| `REDIS_PASSWORD` | Redis password; leave empty when authentication is disabled | empty |
| `REDIS_DB` | Redis database index | `0` |

## HTTP Routes

| Method | Route | Description |
|---|---|---|
| `GET` | `/` | Main dashboard |
| `GET` | `/weather` | Weather forecast for a selected city |
| `GET` | `/weather/hourly` | Hourly forecast for a selected day |
| `GET` | `/favorites` | Favorite cities |
| `POST` | `/favorites/add` | Add a city to favorites |
| `DELETE` | `/favorites/delete` | Remove a city from favorites |
| `GET` | `/searchhistory` | Recent search history |

## Testing and Code Quality

Run the complete test suite:

```bash
make test
```

Run formatting and static checks:

```bash
make fmt
make vet
```

The current test suite includes:

- Unit tests for service-layer data processing and validation
- Integration tests for SQLite storage

Direct command equivalents:

```bash
go test ./...
go fmt ./...
go vet ./...
```

## Continuous Integration

GitHub Actions runs two workflows:

- **Check** — runs `golangci-lint`, then executes the test suite with the configured Go version matrix
- **Build** — verifies that the Docker image builds successfully on pushes to `main`

## Request Flow

1. The user searches for a city.
2. The application resolves the city's coordinates.
3. The service checks Redis for cached weather data.
4. If the cache does not contain the response, the service requests it from Open-Meteo.
5. The service processes the response into application models.
6. Go templates render the page or HTMX fragment.
7. Favorites and search history are stored in SQLite and associated with the user's session.
