# Task Management API Documentation

## Overview

The Task Management API allows users to manage tasks with the ability to create, read, update, and delete tasks. The API includes JWT-based authentication and authorization to ensure that only authenticated users can access the API, and only users with the appropriate roles can perform certain actions.

## Table of Contents

1. [Folder Structure](#folder-structure)
2. [Getting Started](#getting-started)
   - [Prerequisites](#prerequisites)
   - [Installation](#installation)
   - [Running the Application](#running-the-application)
3. [Authentication & Authorization](#authentication--authorization)
   - [Register User](#register-user)
   - [Login](#login)
   - [JWT Authentication](#jwt-authentication)
   - [Role-Based Authorization](#role-based-authorization)
4. [Task Endpoints](#task-endpoints)
   - [Create a Task](#create-a-task)
   - [Get All Tasks](#get-all-tasks)
   - [Get Task by ID](#get-task-by-id)
   - [Update a Task](#update-a-task)
   - [Delete a Task](#delete-a-task)
5. [Promote User](#promote-user)
6. [Error Handling](#error-handling)
7. [Notes](#notes)

## Folder Structure

```
task_manager/
├── main.go
├── controllers/
│ └── controller.go
├── models/
│ ├── task.go
│ └── user.go
├── data/
│ ├── task_service.go
│ └── user_service.go
├── middleware/
│ └── auth_middleware.go
├── router/
│ └── router.go
├── docs/
│ └── api_documentation.md
├── config.json
└── go.mod
```

- **main.go**: Entry point of the application.
- **controllers/controller.go**: Handles incoming HTTP requests and invokes the appropriate service methods for both tasks and user authentication.
- **models/task.go**: Defines the Task struct.
- **models/user.go**: Defines the User struct.
- **data/task_service.go**: Contains business logic and data manipulation functions for tasks.
- **data/user_service.go**: Contains business logic and data manipulation functions for users, including password hashing and JWT generation.
- **middleware/auth_middleware.go**: Implements middleware to validate JWT tokens for authentication and authorization.
- **router/router.go**: Sets up the routes and initializes the Gin router.
- **docs/api_documentation.md**: Contains this API documentation and other related documentation.
- **config.json**: Stores configuration settings, such as the JWT secret key.
- **go.mod**: Defines the module and its dependencies.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher
- [MongoDB](https://www.mongodb.com/try/download/community) installed and running

### Installation

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yourusername/task_manager.git
   cd task_manager_authentication_authorization
   ```

2. **Install Dependencies:**
   Ensure you have Go installed, then run:

   ```bash
   go mod tidy
   ```

3. **Configure the Application:**
   Create a `config.json` file in the root directory with the following content:
   ```json
   {
     "jwt_secret_key": "your-secret-key"
   }
   ```
   Replace `"your-secret-key"` with a secure secret key for JWT token generation.

### Running the Application

1. **Start MongoDB:**
   Make sure your MongoDB instance is running.

2. **Run the Application:**
   ```bash
   go run main.go
   ```
   The application will start on `http://localhost:8080`.

## Authentication & Authorization

### Register User

**Endpoint:** `POST /register`

**Description:** Registers a new user with a unique username and password.

**Request:**

```json
{
  "username": "uniqueUsername",
  "password": "yourPassword"
}
```

**Response:**

- **201 Created:** User successfully registered.
- **400 Bad Request:** Username already exists.

**Example:**

```bash
curl -X POST http://localhost:8080/register -d '{"username":"uniqueUsername", "password":"yourPassword"}' -H "Content-Type: application/json"
```

### Login

**Endpoint:** `POST /login`

**Description:** Authenticates a user and generates a JWT token for subsequent requests.

**Request:**

```json
{
  "username": "yourUsername",
  "password": "yourPassword"
}
```

**Response:**

- **200 OK:** Returns a JWT token.
- **401 Unauthorized:** Invalid username or password.

**Example:**

```bash
curl -X POST http://localhost:8080/login -d '{"username":"yourUsername", "password":"yourPassword"}' -H "Content-Type: application/json"
```

### JWT Authentication

JWT tokens are used to authenticate requests. Include the token in the `Authorization` header for all protected endpoints.

**Header:**

```
Authorization: Bearer <your-jwt-token>
```

### Role-Based Authorization

- **Admin**: Can create, update, delete tasks, and promote users.
- **User**: Can only view tasks.

## Task Endpoints

### Create a Task

**Endpoint:** `POST /tasks`

**Description:** Creates a new task. Only accessible by admins.

**Request:**

```json
{
  "title": "Task Title",
  "description": "Task Description",
  "dueDate": "2024-08-20T12:34:56Z"
}
```

**Response:**

- **201 Created:** Task successfully created.
- **403 Forbidden:** Access denied if the user is not an admin.

**Example:**

```bash
curl -X POST http://localhost:8080/tasks -d '{"title":"Task Title", "description":"Task Description", "dueDate":"2024-08-20T12:34:56Z"}' -H "Authorization: Bearer <your-jwt-token>" -H "Content-Type: application/json"
```

### Get All Tasks

**Endpoint:** `GET /tasks`

**Description:** Retrieves a list of all tasks. Accessible by both users and admins.

**Response:**

- **200 OK:** Returns a list of tasks.

**Example:**

```bash
curl -X GET http://localhost:8080/tasks -H "Authorization: Bearer <your-jwt-token>"
```

### Get Task by ID

**Endpoint:** `GET /tasks/:id`

**Description:** Retrieves a task by its ID. Accessible by both users and admins.

**Response:**

- **200 OK:** Returns the task details.
- **404 Not Found:** Task not found.

**Example:**

```bash
curl -X GET http://localhost:8080/tasks/60c72b2f5b3c3c7a7a9c93a5 -H "Authorization: Bearer <your-jwt-token>"
```

### Update a Task

**Endpoint:** `PUT /tasks/:id`

**Description:** Updates an existing task. Only accessible by admins.

**Request:**

```json
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "dueDate": "2024-08-21T12:34:56Z"
}
```

**Response:**

- **200 OK:** Task successfully updated.
- **403 Forbidden:** Access denied if the user is not an admin.
- **404 Not Found:** Task not found.

**Example:**

```bash
curl -X PUT http://localhost:8080/tasks/60c72b2f5b3c3c7a7a9c93a5 -d '{"title":"Updated Task Title", "description":"Updated Task Description", "dueDate":"2024-08-21T12:34:56Z"}' -H "Authorization: Bearer <your-jwt-token>" -H "Content-Type: application/json"
```

### Delete a Task

**Endpoint:** `DELETE /tasks/:id`

**Description:** Deletes a task by its ID. Only accessible by admins.

**Response:**

- **200 OK:** Task successfully deleted.
- **403 Forbidden:** Access denied if the user is not an admin.
- **404 Not Found:** Task not found.

**Example:**

```bash
curl -X DELETE http://localhost:8080/tasks/60c72b2f5b3c3c7a7a9c93a5 -H "Authorization: Bearer <your-jwt-token>"
```

## Promote User

**Endpoint:** `POST /promote`

**Description:** Promotes a user to admin. Only accessible by admins.

**Request:**

```json
{
  "username": "usernameToPromote"
}
```

**Response:**

- **200 OK:** User successfully promoted.
- **403 Forbidden:** Access denied if the user is not an admin.

**Example:**

```bash
curl -X POST http://localhost:8080/promote -d '{"username":"usernameToPromote"}' -H "Authorization: Bearer <your-jwt-token>" -H "Content-Type: application/json"
```

**Response:**

- [API Documentation on Postman](https://documenter.getpostman.com/view/34185326/2sA3s4kVco)
