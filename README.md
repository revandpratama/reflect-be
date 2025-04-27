# Reflect API

Reflect API is a microservices-based backend architecture designed for scalable applications requiring robust authentication, efficient logging, and flexible REST APIs. Built with a focus on separation of concerns, it features dedicated services for authentication, business logic, and logging—all communicating via Kafka messaging. With MinIO integration for object storage and PASETO tokens for secure authentication, Reflect API provides a modern, maintainable foundation for building resilient API.

## 📋 Table of Contents

- [Architecture Overview](#architecture-overview)
- [Services](#services)
  - [Main App](#main-app)
  - [Auth Service](#auth-service)
  - [Logging Service](#logging-service)
- [Technologies](#technologies)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Development](#development)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## 🏗️ Architecture Overview

Reflect API is built on a microservices architecture with three main components:

```
/project-root
  /auth-service     # Authentication and authorization
  /main-app         # User and application logic
  /logging-service  # Kafka log consumer
  /shared/proto     # Proto files for gRPC (if used)
  /deployments      # Docker, Kubernetes, or Compose files
```

The services communicate via Kafka messaging, with the logging service consuming messages from both the auth service and main app.

![Architecture Diagram](https://i.imgur.com/VVHg8Mv.png)

## 🔌 Services

### Main App

The main application serves as the public-facing REST API that exposes various services to users.

**Responsibilities:**
- User-related functionality (CRUD operations, profiles)
- Application-specific logic
- REST API for Posts, Comments, and Users
- Communication with auth-service for authentication via gRPC protocol
- Sending logs to the logging service via Kafka
- Integration with MinIO for object storage
- Utilizing Redis for data caching

### Auth Service

A dedicated service that handles authentication and authorization in a secure, isolated environment.

**Responsibilities:**
- User authentication (login)
- User registration with validation
- Password hashing and security
- PASETO token generation for authentication

### Logging Service

A Kafka-based logging system that captures logs from all other services.

**Responsibilities:**
- Collects logs from different microservices
- Stores logs in Kafka topics for later analysis
- Can be extended to monitoring tools (Prometheus, ELK Stack, etc.)

## 🛠️ Technologies

- **Backend**: Go, Fiber HTTP library, grpc
- **Authentication**: PASETO tokens
- **Message Broker**: Apache Kafka (segment.io/kafka-go)
- **Object Storage**: MinIO
- **Caching**: Redis
- **Database**: Postgresql
- **API Documentation**: Swagger
- **Containerization**: Docker
- **Orchestration**: docker-compose

## 🚀 Getting Started

### Prerequisites

- Docker and Docker Compose
- Kafka
- MinIO
- Redis

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/revandpratama/reflect-be.git
   cd reflect-be
   ```

2. Set up environment variables:
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```


## 📚 API Documentation

### Main App API Endpoints

#### User Endpoints
- `POST /api/users` - Create a new user
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user
- `DELETE /api/users/:id` - Delete user

#### Post Endpoints
- `POST /api/posts` - Create a new post
- `GET /api/posts` - Get all posts
- `GET /api/posts/:id` - Get post by ID
- `PUT /api/posts/:id` - Update post
- `DELETE /api/posts/:id` - Delete post

#### Comment Endpoints
- `POST /api/posts/:postId/comments` - Create a new comment
- `GET /api/posts/:postId/comments` - Get all comments for a post
- `PUT /api/comments/:id` - Update comment
- `DELETE /api/comments/:id` - Delete comment

### Auth Service API Endpoints

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - Authenticate user and get token

## 💻 Development

### Setting Up a Development Environment

1. [Development setup instructions]

### Running Tests

```bash
# Implemented Soon :)
```

## 🌐 Deployment

### Docker Deployment

```bash

# Run necessary service
cd /deployments/service-deploy
docker-compose up --build -d

# Run go app
cd /deplyoments
docker-compose up --build -d

```

## 🤝 Contributing

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT LICENSE - checkout the LICENSE file for details.
