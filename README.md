# Project docs

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application


```bash
make build
```

- [Docker](https://www.docker.com/) and Docker-compose (for containerized setup)
- [Node.js](https://nodejs.org/en/) (for frontend development and real-time backend)
- [Go](https://golang.org/) (for authentication and user management backend)
- [PostgreSQL](https://www.postgresql.org/) (for the database)
- [Redis](https://redis.io/) (for session management)
- Nginx 


Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
