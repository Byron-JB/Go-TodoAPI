# Development Guide

This README explains how to run the project with Docker Compose and how to use the .http files in the dev-resources folder to emulate HTTP requests during development.

## Prerequisites
- Docker Desktop (Windows/macOS) or Docker Engine (Linux)
- Docker Compose plugin (bundled with recent Docker Desktop)
- Terminal open at the project root containing docker-compose.yml

Verify:
```shell script
# Bash
docker --version
docker compose version
```


## Typical workflow
```shell script
# Bash
docker compose build        # Build or rebuild images
docker compose up           # Start services (foreground)
docker compose up -d        # Start services (detached/background)
docker compose down         # Stop and remove containers, networks
```


## Commands in detail

### Build images
Builds or rebuilds all service images defined in docker-compose.yml.
```shell script
# Bash
docker compose build
```

Useful flags:
- No cache:
```shell script
# Bash
docker compose build --no-cache
```


### Start services
Start the full stack.
```shell script
# Bash
docker compose up           # Runs in foreground, streams logs
docker compose up -d        # Detached mode
```

Watch logs:
```shell script
# Bash
docker compose logs -f
docker compose logs -f <service-name>
```

### Stop and remove services
```shell script
# Bash
docker compose down
```
### Clean reset
```shell script
# Bash
docker compose down -v --rmi local
```

Optional cleanup:
- Remove volumes (deletes persisted dev data):
```shell script
# Bash
docker compose down -v
```

- Remove images built by compose:
```shell script
# Bash
docker compose down --rmi local
```


### Inspect and troubleshoot
```shell script
# Bash
docker compose ps                   # Status of services
docker compose logs -f api          # Tail logs of a specific service
docker compose exec api sh          # Shell into a service container
```


## Using the dev-resources .http files

The dev-resources folder contains .http files that emulate the requests the system will receive. These files are intended to be run with the built-in HTTP Client in an IDE (such as GoLand), or you can translate them to curl commands if you prefer the terminal.

Typical locations:
- dev-resources/requests/*.http
- dev-resources/*.http

### Running .http files in GoLand
1. Open the .http file in GoLand.
2. Hover over a request; click the green Run icon next to it.
3. View the response right inside the editor.

Tips:
- If requests reference variables, configure an environment file (e.g., http-client.env.json) or inline variables at the top of the .http file.
- Ensure the service is running (docker compose up) before sending requests.

Example .http content:
```
### Get all todos
GET http://localhost:<PORT>/todos
Accept: application/json

### Create multiple todos
POST http://localhost:<PORT>/todos
Content-Type: application/json

[
  {
    "title": "Sample",
    "priority": 1,
    "due_date": "2025-12-31"
  }
]
```

Replace <PORT> with the published port configured in docker-compose.yml.

### Using curl instead (optional)
You can replicate the same requests with curl:
```shell script
# Bash
curl -i http://localhost:<PORT>/todos

curl -i -X POST http://localhost:<PORT>/todos \
  -H "Content-Type: application/json" \
  -d '[{"title":"Sample","priority":1,"due_date":"2025-12-31"}]'
```
