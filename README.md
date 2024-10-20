# Project Documentation

## Overview

This project is a modern web application that leverages multiple technologies to provide a seamless user experience. It integrates **MongoDB** for document storage, **Socket.IO** for real-time communication with an Express backend, and **gRPC** for efficient inter-service communication with a Gin-based backend in Go. The application is designed to manage user authentication, provide real-time updates, and facilitate smooth interactions across different components.

## Technologies Used

- **TypeScript**: A superset of JavaScript that adds static types, enhancing code quality and maintainability for the backend.
- **MongoDB**: A NoSQL database for storing and retrieving document-based content.
- **PostgreSQL**: A relational database for handling structured data and complex queries.
- **Redis**: An in-memory data structure store is used for session management and caching.
- **Go (Gin)**: The backend framework for user authentication and management.
- **Node.js (Express)**: The framework for real-time backend communication.
- **Socket.IO**: Enables real-time, bidirectional communication between clients and servers.
- **gRPC**: A high-performance RPC framework is used for efficient communication between services.
- **Docker**: This is for containerizing the application and managing dependencies.
- **Nginx**: Acts as a reverse proxy to handle incoming requests and route them to the appropriate service.

## Getting Started

These instructions will help you set up a copy of the project on your local machine for development and testing purposes. 

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
- Nginx (proxy)


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
