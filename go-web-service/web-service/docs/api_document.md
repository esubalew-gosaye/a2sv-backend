## 📘 Task API Documentation

Base URL: `http://localhost:8080`

---

### [GET] Get All Tasks  
Retrieves a list of all tasks.

**URL:**  
`GET http://localhost:8080/tasks`

---

### [GET] Get Task by ID  
Fetch a single task using its ID.

**URL:**  
`GET http://localhost:8080/tasks/4`

---

### [POST] Create New Task  
Creates a new task with the following fields:

**URL:**  
`POST http://localhost:8080/tasks`

**Request Body:**
```json
{
  "id": "3",
  "title": "Task 3",
  "description": "Third task",
  "due_date": "2025-04-22T15:04:05Z",
  "status": "Pending"
}
```

---

### [PUT] Update Existing Task  
Replaces all fields of a task. Make sure to send all required fields.

**URL:**  
`PUT http://localhost:8080/tasks/{id}`

**Request Body:**
```json
{
  "id": "4",
  "title": "Updated Task",
  "description": "Completely updated task",
  "due_date": "2025-04-25T10:00:00Z",
  "status": "In Progress"
}
```

---

### [PATCH] Partially Update Task  
Updates only specified fields (e.g., just the status).

**URL:**  
`PATCH http://localhost:8080/tasks/{id}`

**Request Body:**
```json
{
  "status": "Completed"
}
```

---

### [DELETE] Delete Task  
Deletes the task with the specified ID.

**URL:**  
`DELETE http://localhost:8080/tasks/{id}`