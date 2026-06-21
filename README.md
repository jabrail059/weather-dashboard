# Weather Dashboard

A server-rendered Go web application for checking weather forecasts by city.

The project combines Go HTML templates with HTMX for partial page updates, stores user-related data in SQLite, and uses Redis to cache external weather API responses.

## Screenshots

### Main dashboard

![Main dashboard](docs/screenshots/home.png)

### Hourly forecast

![Hourly forecast](docs/screenshots/hourly.png)

## Features

- Search weather by city
- View current weather conditions
- View a 7-day forecast
- Open hourly forecast details for a selected day
- Save and remove favorite cities
- View recent search history
- Use an auto-detected city suggestion
- Update page sections dynamically without full page reloads

## Tech Stack

- Go
- Chi
- HTMX
- HTML templates
- CSS
- SQLite
- Redis
- Open-Meteo API

## Technical Overview

The application is built as a server-rendered Go web app.

- HTTP routing is handled with Chi.
- Pages are rendered on the server using Go HTML templates.
- HTMX is used for partial page updates.
- Weather data is requested from the Open-Meteo API.
- Redis is used to cache external API responses.
- SQLite stores favorite cities and search history.
- Configuration is provided through environment variables.
- The HTTP server supports graceful shutdown.

## Configuration

Create a local `.env` file from the example file:

```bash
cp .env.example .env
```

Example configuration:

```env
APP_ADDR=:8081
SQLITE_PATH=./data/weather.db
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

| Variable | Description |
| --- | --- |
| `APP_ADDR` | HTTP server address |
| `SQLITE_PATH` | Path to the SQLite database file |
| `REDIS_ADDR` | Redis server address |
| `REDIS_PASSWORD` | Redis password, if required |
| `REDIS_DB` | Redis database index |

## Running Locally

### 1. Clone the repository

```bash
git clone https://github.com/jabrail059/weather-dashboard.git
cd weather-dashboard
```

### 2. Start Redis

Redis is required for caching weather API responses.

```bash
redis-server
```

### 3. Create the environment file

```bash
cp .env.example .env
```

### 4. Run the application

```bash
go run ./cmd/app
```

The application will be available at:

```text
http://localhost:8081
```

## Main Routes

| Route | Description |
| --- | --- |
| `/` | Main dashboard |
| `/weather` | Weather forecast for a selected city |
| `/weather/hourly` | Hourly forecast for a selected day |
| `/favorites` | Favorite cities |
| `/searchhistory` | Recent search history |
