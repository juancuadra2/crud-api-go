# CRUD API Go - Hexagonal Architecture

A RESTful API built with Go using Hexagonal Architecture (Ports and Adapters pattern) for creating, reading, updating, and deleting resources.

## 📋 Table of Contents

- [About the Project](#about-the-project)
- [Architecture](#architecture)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Development](#development)
- [Deployment](#deployment)
- [Contributing](#contributing)
- [License](#license)

## 🎯 About the Project

This project demonstrates a clean implementation of a CRUD REST API in Go following Hexagonal Architecture principles. The architecture promotes separation of concerns, testability, and maintainability by isolating business logic from external dependencies.

### Key Features

- ✅ RESTful API design
- ✅ Hexagonal Architecture (Ports and Adapters)
- ✅ Clean separation of concerns
- ✅ Dependency injection
- ✅ Easy to test and mock
- ✅ Database agnostic design
- ✅ Environment-based configuration
- ✅ Structured logging
- ✅ Error handling middleware
- ✅ API versioning

### Built With

- [Go](https://golang.org/) - Programming language
- [Gin](https://github.com/gin-gonic/gin) or [Echo](https://echo.labstack.com/) - HTTP web framework
- [GORM](https://gorm.io/) or [sqlx](https://github.com/jmoiron/sqlx) - Database ORM/toolkit
- [PostgreSQL](https://www.postgresql.org/) - Database (can be swapped)
- [Docker](https://www.docker.com/) - Containerization

## 🏗️ Architecture

This project follows the **Hexagonal Architecture** (also known as Ports and Adapters) pattern, which organizes code into three main layers:

### Domain Layer (Core/Business Logic)
- **Entities**: Core business objects
- **Value Objects**: Immutable objects that represent concepts
- **Use Cases/Services**: Business logic and application rules
- **Repository Interfaces (Ports)**: Abstract definitions for data access

### Application Layer (Ports)
- **Input Ports**: Interfaces for use cases (e.g., service interfaces)
- **Output Ports**: Interfaces for external dependencies (e.g., repository interfaces)

### Infrastructure Layer (Adapters)
- **Primary/Driving Adapters**: HTTP handlers, CLI, gRPC servers
- **Secondary/Driven Adapters**: Database repositories, external APIs, message queues

### Benefits

- **Independence**: Business logic is independent of frameworks, UI, and databases
- **Testability**: Easy to test core logic without external dependencies
- **Flexibility**: Easy to swap implementations (e.g., change database)
- **Maintainability**: Clear separation of concerns

```
┌─────────────────────────────────────────────────────┐
│              Primary Adapters (HTTP/CLI)            │
│                    (Driving Side)                   │
└──────────────────┬──────────────────────────────────┘
                   │
        ┌──────────▼──────────┐
        │   Application Port  │
        │    (Use Cases)      │
        └──────────┬──────────┘
                   │
        ┌──────────▼──────────┐
        │   Domain/Core       │
        │  (Business Logic)   │
        │    (Entities)       │
        └──────────┬──────────┘
                   │
        ┌──────────▼──────────┐
        │  Repository Port    │
        │   (Interfaces)      │
        └──────────┬──────────┘
                   │
┌──────────────────▼──────────────────────────────────┐
│         Secondary Adapters (DB/External APIs)       │
│                    (Driven Side)                    │
└─────────────────────────────────────────────────────┘
```

## 📁 Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── domain/                     # Domain layer (entities, value objects)
│   │   ├── entity/
│   │   │   └── user.go            # Domain entities
│   │   └── repository/
│   │       └── user_repository.go # Repository interfaces (ports)
│   ├── application/                # Application layer (use cases)
│   │   └── service/
│   │       └── user_service.go    # Business logic/use cases
│   └── infrastructure/             # Infrastructure layer (adapters)
│       ├── adapter/
│       │   ├── http/              # HTTP handlers (primary adapter)
│       │   │   ├── handler/
│       │   │   │   └── user_handler.go
│       │   │   ├── middleware/
│       │   │   │   └── error_handler.go
│       │   │   └── router/
│       │   │       └── router.go
│       │   └── repository/        # Database implementation (secondary adapter)
│       │       └── postgres/
│       │           └── user_repository.go
│       ├── config/                 # Configuration
│       │   └── config.go
│       └── database/               # Database connection
│           └── connection.go
├── pkg/                            # Public reusable packages
│   ├── logger/
│   └── validator/
├── test/                           # Integration and e2e tests
├── docker/
│   └── Dockerfile
├── .env.example                    # Environment variables template
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile                        # Build and development commands
└── README.md
```

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.21 or higher - [Download](https://golang.org/dl/)
- **Docker** and **Docker Compose** - [Download](https://www.docker.com/get-started)
- **PostgreSQL** 15+ (if running locally without Docker)
- **Make** (optional, for using Makefile commands)

## 🚀 Installation

### 1. Clone the repository

```bash
git clone https://github.com/juancuadra2/crud-api-go.git
cd crud-api-go
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Copy environment configuration

```bash
cp .env.example .env
```

Edit `.env` file with your configuration values.

## ⚙️ Configuration

The application uses environment variables for configuration. Create a `.env` file based on `.env.example`:

```env
# Application
APP_NAME=crud-api-go
APP_ENV=development
APP_PORT=8080

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=crud_db
DB_SSLMODE=disable

# JWT (if using authentication)
JWT_SECRET=your-secret-key
JWT_EXPIRATION=24h

# Logging
LOG_LEVEL=debug
```

## 🏃 Running the Application

### Using Docker Compose (Recommended)

```bash
# Start all services (API + Database)
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down
```

The API will be available at `http://localhost:8080`

### Using Make Commands

```bash
# Run the application
make run

# Run with hot reload (using air)
make dev

# Build the application
make build

# Run tests
make test

# Run linter
make lint
```

### Manual Execution

```bash
# Start PostgreSQL (if not using Docker)
# Make sure PostgreSQL is running on port 5432

# Run migrations (if applicable)
make migrate-up

# Run the application
go run cmd/api/main.go
```

## 📚 API Documentation

### Base URL

```
http://localhost:8080/api/v1
```

### Endpoints

#### Users

| Method | Endpoint | Description | Auth |
|--------|----------|-------------|------|
| GET    | `/users` | Get all users | No |
| GET    | `/users/:id` | Get user by ID | No |
| POST   | `/users` | Create new user | No |
| PUT    | `/users/:id` | Update user | No |
| DELETE | `/users/:id` | Delete user | No |

### Request/Response Examples

#### Create User

**Request:**
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 30
  }'
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 30,
  "created_at": "2025-11-15T10:00:00Z",
  "updated_at": "2025-11-15T10:00:00Z"
}
```

#### Get All Users

**Request:**
```bash
curl http://localhost:8080/api/v1/users
```

**Response:**
```json
{
  "data": [
    {
      "id": "123e4567-e89b-12d3-a456-426614174000",
      "name": "John Doe",
      "email": "john.doe@example.com",
      "age": 30,
      "created_at": "2025-11-15T10:00:00Z",
      "updated_at": "2025-11-15T10:00:00Z"
    }
  ],
  "total": 1
}
```

#### Error Response

```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Invalid input data",
    "details": [
      "email is required",
      "age must be greater than 0"
    ]
  }
}
```

### Status Codes

- `200 OK` - Request successful
- `201 Created` - Resource created successfully
- `400 Bad Request` - Invalid request data
- `404 Not Found` - Resource not found
- `500 Internal Server Error` - Server error

## 🧪 Testing

### Unit Tests

Run unit tests for individual components:

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run tests for specific package
go test ./internal/application/service/...
```

### Integration Tests

```bash
# Run integration tests
make test-integration

# Run with Docker
docker-compose -f docker-compose.test.yml up --abort-on-container-exit
```

### Test Structure

```
internal/
├── domain/
│   └── entity/
│       └── user_test.go
├── application/
│   └── service/
│       └── user_service_test.go
└── infrastructure/
    └── adapter/
        ├── http/
        │   └── handler/
        │       └── user_handler_test.go
        └── repository/
            └── postgres/
                └── user_repository_test.go
```

## 💻 Development

### Code Style

This project follows Go standard conventions:

- Use `gofmt` for formatting
- Follow [Effective Go](https://golang.org/doc/effective_go.html) guidelines
- Use meaningful variable and function names
- Write comments for exported functions

### Linting

```bash
# Run linter
make lint

# Auto-fix linting issues
make lint-fix
```

### Adding New Features

When adding new features following hexagonal architecture:

1. **Define the entity** in `internal/domain/entity/`
2. **Create repository interface** in `internal/domain/repository/`
3. **Implement use case** in `internal/application/service/`
4. **Create HTTP handler** in `internal/infrastructure/adapter/http/handler/`
5. **Implement repository** in `internal/infrastructure/adapter/repository/`
6. **Register routes** in router
7. **Write tests** for each layer

### Database Migrations

```bash
# Create new migration
make migrate-create name=create_users_table

# Run migrations
make migrate-up

# Rollback migrations
make migrate-down

# Check migration status
make migrate-status
```

## 🚢 Deployment

### Docker

Build and push Docker image:

```bash
# Build image
docker build -t crud-api-go:latest -f docker/Dockerfile .

# Run container
docker run -p 8080:8080 --env-file .env crud-api-go:latest
```

### Using Docker Compose in Production

```bash
# Build and start services
docker-compose -f docker-compose.prod.yml up -d

# Scale API instances
docker-compose -f docker-compose.prod.yml up -d --scale api=3
```

### Environment Variables for Production

Ensure the following environment variables are properly set:

- Use strong `JWT_SECRET`
- Set `APP_ENV=production`
- Configure proper `DB_SSLMODE`
- Use secure database credentials
- Set appropriate `LOG_LEVEL`

### Cloud Deployment

The application can be deployed to:

- **AWS**: ECS, EKS, or Elastic Beanstalk
- **Google Cloud**: Cloud Run, GKE
- **Azure**: Container Instances, AKS
- **Heroku**: Using container deployment
- **DigitalOcean**: App Platform or Kubernetes

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

### Coding Guidelines

- Write tests for new features
- Follow Go conventions and best practices
- Maintain hexagonal architecture principles
- Update documentation as needed
- Ensure all tests pass before submitting PR

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

Juan Cuadrado - [@juancuadra2](https://github.com/juancuadra2)

Project Link: [https://github.com/juancuadra2/crud-api-go](https://github.com/juancuadra2/crud-api-go)

## 🙏 Acknowledgments

- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)
- [Go Best Practices](https://github.com/golang-standards/project-layout)
- [Clean Architecture by Robert C. Martin](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Domain-Driven Design](https://martinfowler.com/tags/domain%20driven%20design.html)