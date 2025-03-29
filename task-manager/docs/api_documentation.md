# Task Manager API Documentation

## Overview
The Task Manager API is a backend service built using **Gin (Go) and MongoDB**. It provides user authentication with **JWT tokens** and allows users to create, update, delete, and retrieve tasks.

## Features
- **User Authentication:** JWT-based authentication.
- **Task Management:** CRUD operations (Create, Read, Update, Delete).
- **User-Specific Tasks:** Users can only access their own tasks.
- **Role-Based Access Control (RBAC):** Admins have additional privileges.

## Technologies Used
- **Go (Gin Framework)** – For handling HTTP requests.
- **MongoDB** – For storing tasks and users.
- **JWT (JSON Web Tokens)** – For authentication.
- **bcrypt** – For password hashing.

---

## API Endpoints

### **Authentication**
#### **Register User**
```http
POST /auth/register
```
**Body:**
```json
{
  "username": "johndoe",
  "password": "password123",
  "role": "user"
}
```
**Response:**
```json
{
  "message": "User registered successfully"
}
```

#### **Login User**
```http
POST /auth/login
```
**Body:**
```json
{
  "username": "johndoe",
  "password": "password123"
}
```
**Response:**
```json
{
  "token": "jwt_token_here"
}
```

---

### **Task Management**
#### **Create a Task**
```http
POST /tasks
Authorization: Bearer <JWT_TOKEN>
```
**Body:**
```json
{
  "title": "Complete assignment",
  "status": "pending"
}
```
**Response:**
```json
{
  "id": "task_id_here",
  "user_id": "user_id_here",
  "title": "Complete assignment",
  "status": "pending"
}
```

#### **Get User's Tasks**
```http
GET /tasks
Authorization: Bearer <JWT_TOKEN>
```
**Response:**
```json
[
  {
    "id": "task_id_here",
    "title": "Complete assignment",
    "status": "pending"
  }
]
```

#### **Update a Task**
```http
PUT /tasks/:id
Authorization: Bearer <JWT_TOKEN>
```
**Body:**
```json
{
  "status": "completed"
}
```

#### **Delete a Task**
```http
DELETE /tasks/:id
Authorization: Bearer <JWT_TOKEN>
```

---

## Middleware
### **Authentication Middleware (`AuthMiddleware`)**
- Extracts JWT token from the `Authorization` header.
- Validates the token and extracts the `userID`.
- Stores `userID` in the request context.

---

## Running the API
To start the server:
```sh
go run main.go
```

The API will run on **`http://localhost:8080`**.

---




api documentation : follow the link below:

-> https://documenter.getpostman.com/view/37500379/2sAYdeNC2N
