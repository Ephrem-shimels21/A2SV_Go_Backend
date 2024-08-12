# Task Management API - README

---

## Overview

This Task Management API is designed to manage tasks effectively with CRUD operations (Create, Read, Update, Delete).

### Key Features

- **CRUD Operations:** Supports creating, reading, updating, and deleting tasks.
- **Error Handling:** Proper error handling is implemented for all MongoDB operations.
- **Validation:** Input validation ensures data integrity.
- **Backward Compatibility:** The API remains backward compatible with the previous in-memory version.

---

## Folder Structure

```plaintext
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

- **main.go:** Entry point of the application.
- **controllers/task_controller.go:** Handles incoming HTTP requests and invokes appropriate service methods.
- **models/task.go:** Defines the data structure for tasks.
- **data/task_service.go:** Contains business logic and data manipulation functions.
- **router/router.go:** Sets up routes and initializes the Gin router.
- **docs/api_documentation.md:** Contains API documentation and related documentation.
- **go.mod:** Defines the module and its dependencies.

## API Endpoints

- **GET /tasks:** Retrieves a list of all tasks.
- **GET /tasks/:id:** Retrieves a specific task by its ID.
- **POST /tasks:** Creates a new task.
- **PUT /tasks/:id:** Updates an existing task.
- **DELETE /tasks/:id:** Deletes a specific task.

---

## Data Model

### Task

```go
type Task struct {
    ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    DueDate     string             `json:"due_date" bson:"due_date"` // Expected format: YYYY-MM-DD
    Status      string             `json:"status" bson:"status"` // Allowed values: "Done", "Need Help", "In Progress"
}
```

---

## Setup Instructions

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Ephrem-shimels21/A2SV_Go_Backend.git
   cd  task_management
   ```

2. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the Application:**

   ```bash
   go run main.go
   ```

4. **Access the API:**
   - The API will be running at `http://localhost:8080`.
   - Use Postman or any other tool to interact with the API.

---

## Validation Rules

- **Due Date:** Must be a valid date string in the format `YYYY-MM-DD`.
- **Status:** Must be one of `"Done"`, `"Need Help"`, or `"In Progress"`.

## Error Handling

- All MongoDB-related errors, such as connection issues, validation errors, or operation failures, are handled gracefully and returned to the client with appropriate status codes.

## Testing

- Use Postman or any other API testing tool to test the endpoints.
- Ensure that tasks are correctly created, updated, retrieved, and deleted.
- Verify data persistence by restarting the application and checking that the data remains intact in MongoDB.

## Documentation

- The full API documentation, including examples of request and response formats, can be found in `docs/api_documentation.md`.

### Postman Documentation

- [API Documentation on Postman](https://documenter.getpostman.com/view/34185326/2sA3s3Js5z)
