# Project boilerplate: Go + Gin + Docker + Air
A containerised [Gin](https://github.com/gin-gonic/gin) app boilerplate, with a containerised development environment using Air and Docker Compose.

## Development

Install the dependancies:

> This project uses [Go mod](https://blog.golang.org/using-go-modules), the official module manager, to handle Go modules in a portable way without having to worry about GOPATH.

```bash
go mod download
go mod vendor
go mod verify
```

Run locally:

> This builds the Docker image and runs it automatically with the config defined in `docker-compose.yaml`. This saves you having to build the docker image and then run a manual `docker run` command with all the flags (for environment variables, ports, etc).

```bash
docker compose up
```

## Production

```bash
# Build Docker image for production:
docker build -t gin-air-docker-boilerplate -f Dockerfile.production .
```