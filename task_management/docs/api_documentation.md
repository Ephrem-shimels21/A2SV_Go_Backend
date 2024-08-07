# Task Management API Documentation

## Endpoints

### GET /tasks

**Description:** Get a list of all tasks.

**Response:**

- Status: 200 OK
- Body: Array of tasks

### GET /tasks/:id

**Description:** Get the details of a specific task.

**Response:**

- Status: 200 OK
- Body: Task object

### POST /tasks

**Description:** Create a new task.

**Request:**

- Body: JSON object containing task's title, description, due date, and status

**Response:**

- Status: 201 Created
- Body: Created task object

### PUT /tasks/:id

**Description:** Update a specific task.

**Request:**

- Body: JSON object with the new details of the task

**Response:**

- Status: 200 OK
- Body: Updated task object

### DELETE /tasks/:id

**Description:** Delete a specific task.

**Response:**

- Status: 204 No Content
