# Task Management API Documentation

## Overview

The Task Management API allows users to manage tasks, with the ability to create, read, update, and delete tasks. The API includes JWT-based authentication and authorization to ensure that only authenticated users can access the API, and only users with the appropriate roles can perform certain actions.

## Table of Contents

1. [Authentication & Authorization](#authentication--authorization)
   - [Register User](#register-user)
   - [Login](#login)
   - [JWT Authentication](#jwt-authentication)
   - [Role-Based Authorization](#role-based-authorization)
2. [Task Endpoints](#task-endpoints)
   - [Create a Task](#create-a-task)
   - [Get All Tasks](#get-all-tasks)
   - [Get Task by ID](#get-task-by-id)
   - [Update a Task](#update-a-task)
   - [Delete a Task](#delete-a-task)
3. [Promote User](#promote-user)
4. [Error Handling](#error-handling)

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

## Error Handling

All error responses will include a JSON object with an `error` key explaining the error. For example:

**Response:**

```json
{
  "error": "Unauthorized"
}
```

### Common Error Codes

- **400 Bad Request:** Invalid input or request format.
- **401 Unauthorized:** Invalid credentials or missing token.
- **403 Forbidden:** Access denied due to insufficient permissions.
- **404 Not Found:** The requested resource does not exist.
- **500 Internal Server Error:** An unexpected error occurred on the server.

- [API Documentation on Postman](https://documenter.getpostman.com/view/34185326/2sA3s4kVco)
