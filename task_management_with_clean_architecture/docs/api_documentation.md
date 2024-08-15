# Task Management API - Clean Architecture Refactor

## Overview

This project is a refactor of the existing Task Management API using Clean Architecture principles. The refactor aims to improve maintainability, testability, and scalability by organizing the code into distinct layers with clear separation of concerns. The architecture is structured to ensure that core business logic remains decoupled from external frameworks and services, allowing for easier updates, testing, and future enhancements.

## Project Structure

The project is organized into the following layers:

```
task-manager/
├── Delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── Domain/
│   └── domain.go
├── Infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── Repositories/
│   ├── task_repository.go
│   └── user_repository.go
└── Usecases/
    ├── task_usecases.go
    └── user_usecases.go
```

### Delivery/

- **main.go:** Initializes the HTTP server, sets up dependencies, and defines routing.
- **controllers/controller.go:** Handles incoming HTTP requests and invokes appropriate use case methods.
- **routers/router.go:** Configures routes and initializes the Gin router.

### Domain/

- **domain.go:** Contains core business entities such as `Task` and `User` structs.

### Infrastructure/

- **auth_middleWare.go:** Middleware for handling authentication and authorization via JWT tokens.
- **jwt_service.go:** Functions for generating and validating JWT tokens.
- **password_service.go:** Functions for hashing and comparing passwords for secure credential storage.

### Repositories/

- **task_repository.go:** Interface and implementation for task-related data access.
- **user_repository.go:** Interface and implementation for user-related data access.

### Usecases/

- **task_usecases.go:** Implements use cases for tasks, including creating, updating, retrieving, and deleting tasks.
- **user_usecases.go:** Implements use cases for user-related operations like registration, login, and promotion.

## Key Principles

This refactor adheres to Clean Architecture principles, including:

- **Separation of Concerns:** Clear boundaries between different layers of the application.
- **Dependency Inversion:** Higher-level layers depend on abstractions provided by lower-level layers.
- **Encapsulation of Business Logic:** Business rules are centralized within use cases, ensuring reusability and maintainability.

## Getting Started

### Prerequisites

- Go 1.XX
- MongoDB
- Gin Web Framework

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Ephrem-shimels21/A2SV_Go_Backend.git
   cd task_management_with_clean_architecture
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up the environment variables for MongoDB connection:

   ```bash
   export MONGODB_URI=mongodb://localhost:27017
   ```

4. Run the application:
   ```bash
   go run Delivery/main.go
   ```

## Usage

- **Register User:** `POST /register`
- **Login User:** `POST /login`
- **Promote User:** `PUT /promote`
- **Create Task:** `POST /tasks`
- **Get Tasks by User ID:** `GET /tasks/user/:id`

## Testing

- Unit tests are provided for core components. Run the tests using:
  ```bash
  go test ./...
  ```

## Documentation

- The API documentation is available [here](#).
- Design decisions and guidelines for future development are documented in the [docs](./docs) directory.

## Postman API documentation

- [API Documentation on Postman](https://documenter.getpostman.com/view/34185326/2sA3s4kVco)
