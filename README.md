# Docs-Backend

This project is a real-time collaborative document editing platform, inspired by Google Docs. Users can create, edit, and share documents with real-time updates and secure authentication.

## Features

- Real-time document editing using Socket.io
- Collaborative editing (multiple users can edit the same document simultaneously)
- OAuth-based user authentication
- Secure storage of user data and documents
- Backend built with Express and TypeScript for real-time collaboration
- User data and authentication handled by Go (Fiber) with PostgreSQL and Redis for session management
- Containerized services using Docker for seamless deployment
- Ai suggestions

## Tech Stack

- **Backend**: 
  - **Real-time collaboration**: [Express](https://expressjs.com/) + [Socket.io](https://socket.io/) with TypeScript
  - **User authentication and data management**: [Go (Fiber)](https://gofiber.io/) with [PostgreSQL](https://www.postgresql.org/) and [Redis](https://redis.io/) for caching and session management
- **Authentication**: OAuth 2.0 for secure user login
- **WebSocket Communication**: Socket.io for real-time collaboration

## Prerequisites

Make sure you have the following installed:

- [Docker](https://www.docker.com/) (for containerized setup)
- [Node.js](https://nodejs.org/en/) (for frontend development and real-time backend)
- [Go](https://golang.org/) (for authentication and user management backend)
- [PostgreSQL](https://www.postgresql.org/) (for the database)
- [Redis](https://redis.io/) (for session management)

