# Project boilerplate: Go + Gin + Docker + Air

A containerised [Gin](https://github.com/gin-gonic/gin) app boilerplate, with a containerised development environment using Air and Docker Compose.

| Component        | Choice                                  |
| ---------------- | --------------------------------------- |
| Language         | [Go](https://go.dev/)                   |
| Framework        | [Gin](https://github.com/gin-gonic/gin) |
| Hot Reloading    | [Air](https://github.com/cosmtrek/air)  |
| Containerisation | [Docker](https://www.docker.com/)       |

## Development

Install the dependancies:

> This project uses [Go mod](https://blog.golang.org/using-go-modules), the official module manager, to handle Go modules in a portable way without having to worry about GOPATH.

```bash
go mod download
go mod verify
```

Define environment variables for your development environment:

> These are passed to the Docker container via `docker-compose.yaml` in development. When running in production, the environment variables must be passed to the container when it is run.

```bash
cp .env.example .env
vim .env
```

Run locally:

> This builds the Docker image and runs it automatically with the config defined in `docker-compose.yaml`. This saves you having to build the docker image and then run a manual `docker run` command with all the flags (for environment variables, ports, etc).

```bash
docker compose up
```

## Production

> Note: Environment variables are never baked into the image, or they wouldn't be _environment_ variables. The production environment will start a Docker container based on this image, but it will have to pass the environment variables to the container when it runs it.

```bash
# Build Docker image for production:
docker build -t gin-air-docker-boilerplate -f Dockerfile.production .

# Example manually running a container with environment variables and ports defined:
docker run -p 8080:8080 -e SUPER_SECRET_KEY=abc123 gin-air-docker-boilerplate
```

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=tobyscott25/gin-postgres-air-docker-starter&type=Date)](https://star-history.com/#tobyscott25/gin-postgres-air-docker-starter&Date)
