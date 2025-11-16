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

- [Go](https://golang.org/) 1.25.3 - Programming language
- [Gin](https://github.com/gin-gonic/gin) v1.11.0 - HTTP web framework
- [GORM](https://gorm.io/) v1.31.1 - Database ORM
- [MySQL](https://www.mysql.com/) - Database
- [UUID](https://github.com/google/uuid) - ID generation

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
│       └── main.go                  # Application entry point
├── config/
│   └── config.go                    # Configuration management
├── internal/
│   ├── core/                        # Core domain layer
│   │   ├── domain/
│   │   │   └── models/
│   │   │       └── user.go          # Domain entities
│   │   ├── ports/                   # Interfaces (ports)
│   │   │   ├── user_repository.go   # Repository port
│   │   │   └── user_service.go      # Service port
│   │   └── services/
│   │       └── user_use_case.go     # Business logic/use cases
│   └── adapters/                    # Adapters layer
│       ├── primary/                 # Primary/Driving adapters
│       │   └── rest/
│       │       ├── dtos/            # Data Transfer Objects
│       │       │   ├── create_user_dto.go
│       │       │   ├── update_user_dto.go
│       │       │   ├── user_dto.go
│       │       │   └── error_response_dto.go
│       │       ├── handlers/
│       │       │   └── user_handler.go  # HTTP handlers
│       │       └── router/
│       │           └── router.go        # Route configuration
│       └── secondary/               # Secondary/Driven adapters
│           └── mysql_adapter/
│               ├── user_adapter.go  # MySQL implementation
│               └── entity/
│                   └── user_entity.go
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.25.3 or higher - [Download](https://golang.org/dl/)
- **MySQL** 8.0+ (running locally or accessible remotely)
- **Git** - For cloning the repository

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

### 3. Configure environment variables

You can configure the application using environment variables or use the default values defined in `config/config.go`.

Available environment variables:
- `DB_HOST` - Database host (default: "localhost")
- `DB_PORT` - Database port (default: "3306")
- `DB_USER` - Database user (default: "jcuadrado")
- `DB_PASS` - Database password (default: "jcuadrado")
- `DB_NAME` - Database name (default: "go_lab")
- `SERVER_PORT` - Server port (default: "8080")

## ⚙️ Configuration

The application uses environment variables for configuration. You can set these variables or use the default values:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=jcuadrado
DB_PASS=jcuadrado
DB_NAME=go_lab

# Server Configuration
SERVER_PORT=8080
```

The configuration is managed through `config/config.go`, which loads environment variables with fallback to default values.

## 🏃 Running the Application

### Prerequisites
1. Ensure MySQL is running and accessible
2. Create the database: `CREATE DATABASE go_lab;`
3. Configure environment variables (optional, default values will be used)

### Run the application

```bash
# Run directly with Go
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

### Build and run

```bash
# Build the application
go build -o bin/api cmd/api/main.go

# Run the built binary
./bin/api
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
    "password": "secret123"
  }'
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

#### Get All Users

**Request:**
```bash
curl http://localhost:8080/api/v1/users
```

**Response:**
```json
[
  {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "John Doe",
    "email": "john.doe@example.com"
  },
  {
    "id": "456e7890-e89b-12d3-a456-426614174001",
    "name": "Jane Smith",
    "email": "jane.smith@example.com"
  }
]
```

#### Get User by ID

**Request:**
```bash
curl http://localhost:8080/api/v1/users/123e4567-e89b-12d3-a456-426614174000
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

#### Update User

**Request:**
```bash
curl -X PUT http://localhost:8080/api/v1/users/123e4567-e89b-12d3-a456-426614174000 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com"
  }'
```

**Response:**
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "name": "John Updated",
  "email": "john.updated@example.com"
}
```

#### Delete User

**Request:**
```bash
curl -X DELETE http://localhost:8080/api/v1/users/123e4567-e89b-12d3-a456-426614174000
```

**Response:**
```json
{
  "message": "User deleted successfully"
}
```

#### Error Response

```json
{
  "message": "name is required"
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
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests for specific package
go test ./internal/core/services/...
```

### Test Structure

```
internal/
├── core/
│   ├── domain/
│   │   └── models/
│   │       └── user_test.go
│   └── services/
│       └── user_use_case_test.go
└── adapters/
    ├── primary/
    │   └── rest/
    │       └── handlers/
    │           └── user_handler_test.go
    └── secondary/
        └── mysql_adapter/
            └── user_adapter_test.go
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
golangci-lint run

# Auto-fix linting issues
golangci-lint run --fix
```

### Adding New Features

When adding new features following hexagonal architecture:

1. **Define the entity** in `internal/core/domain/models/`
2. **Create repository interface** in `internal/core/ports/`
3. **Implement use case** in `internal/core/services/`
4. **Create HTTP handler** in `internal/adapters/primary/rest/handlers/`
5. **Implement repository** in `internal/adapters/secondary/mysql_adapter/`
6. **Register routes** in `internal/adapters/primary/rest/router/`
7. **Write tests** for each layer

### Database Migrations

The application uses GORM's AutoMigrate feature. Add your entity to the migration in `cmd/api/main.go`:

```go
// Automigrate database schema
if err := db.AutoMigrate(&entity.UserEntity{}); err != nil {
    log.Fatal("Failed to migrate database schema:", err)
}
```

## 🚢 Deployment

### Build for Production

```bash
# Build the application
go build -o bin/api cmd/api/main.go

# Run the built binary
./bin/api
```

### Environment Variables for Production

Ensure the following environment variables are properly set:

- Use secure database credentials
- Configure proper `DB_HOST` and `DB_PORT`
- Set appropriate `SERVER_PORT`

### Cloud Deployment

The application can be deployed to:

- **AWS**: EC2, ECS, or Elastic Beanstalk
- **Google Cloud**: Compute Engine, Cloud Run
- **Azure**: Virtual Machines, Container Instances
- **Heroku**: Using Go buildpack
- **DigitalOcean**: Droplets or App Platform

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