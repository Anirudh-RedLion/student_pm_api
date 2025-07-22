# Student PM API Documentation

## Overview
The Student PM API is a RESTful backend service built with Go, Gin, GORM, and SQLite. It provides authentication and user management functionality for the Student Profile Management application.

## Base URL
```
http://localhost:8080
```

## Authentication
The API uses JWT (JSON Web Tokens) for authentication. Protected endpoints require a valid Bearer token in the Authorization header.

## Endpoints

### 1. User Registration
**POST** `/api/register`

Register a new user account.

#### Request Body
```json
{
  "name": "string",
  "email": "string",
  "password": "string",
  "role": "string"
}
```

#### Request Example
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "role": "student"
  }'
```

#### Success Response (201)
```json
{
  "message": "User registered"
}
```

#### Error Response (400)
```json
{
  "error": "Email already exists"
}
```

#### Error Response (400)
```json
{
  "error": "Invalid request"
}
```

#### Error Response (500)
```json
{
  "error": "Failed to hash password"
}
```

---

### 2. User Login
**POST** `/api/login`

Authenticate a user and receive a JWT token.

#### Request Body
```json
{
  "email": "string",
  "password": "string"
}
```

#### Request Example
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

#### Success Response (200)
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "Test User",
    "email": "test@example.com",
    "role": "student"
  }
}
```

#### Error Response (400)
```json
{
  "error": "Invalid request"
}
```

#### Error Response (401)
```json
{
  "error": "Invalid credentials"
}
```

#### Error Response (500)
```json
{
  "error": "Failed to generate token"
}
```

---

### 3. Get Current User
**GET** `/api/users/me`

Get the current authenticated user's information.

#### Headers
```
Authorization: Bearer <jwt_token>
```

#### Request Example
```bash
curl -X GET http://localhost:8080/api/users/me \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

#### Success Response (200)
```json
{
  "id": 1,
  "email": "test@example.com",
  "role": "student"
}
```

#### Error Response (401)
```json
{
  "error": "Missing or invalid token"
}
```

#### Error Response (401)
```json
{
  "error": "Invalid token"
}
```

#### Error Response (401)
```json
{
  "error": "Invalid token claims"
}
```

---

## Testing Results

### ✅ Successful Tests

1. **User Registration**
   - ✅ Successfully registered user with valid data
   - ✅ Proper response format and status code (201)

2. **User Login**
   - ✅ Successfully authenticated with valid credentials
   - ✅ JWT token generated and returned
   - ✅ User information included in response

3. **Protected Endpoint Access**
   - ✅ Successfully accessed `/api/users/me` with valid JWT
   - ✅ User information returned correctly

### ✅ Error Handling Tests

1. **Duplicate Email Registration**
   - ✅ Properly rejected duplicate email registration
   - ✅ Correct error message: "Email already exists"

2. **Invalid Login Credentials**
   - ✅ Properly rejected login with wrong password
   - ✅ Correct error message: "Invalid credentials"

3. **Invalid JWT Token**
   - ✅ Properly rejected request with invalid token
   - ✅ Correct error message: "Invalid token"

---

## Data Models

### User Model
```go
type User struct {
    ID           uint   `gorm:"primaryKey" json:"id"`
    Name         string `json:"name"`
    Email        string `gorm:"unique" json:"email"`
    PasswordHash string `json:"-"`
    Role         string `json:"role"`
}
```

---

## Security Features

1. **Password Hashing**: Passwords are hashed using bcrypt with cost factor 14
2. **JWT Authentication**: Secure token-based authentication with 72-hour expiration
3. **Input Validation**: Request validation and sanitization
4. **SQL Injection Protection**: Using GORM ORM for safe database operations

---

## Database

- **Type**: SQLite
- **File**: `studentpm.db`
- **Auto-migration**: User table is automatically created on startup

---

## Error Codes

| Status Code | Description |
|-------------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 500 | Internal Server Error |

---

## Development

### Running the API
```bash
cd student_pm_api
go run main.go
```

### Building the API
```bash
go build -o student_pm_api main.go
```

### Dependencies
- Gin (Web Framework)
- GORM (ORM)
- SQLite Driver
- JWT (Authentication)
- bcrypt (Password Hashing)

---

## Notes

- The API runs on port 8080 by default
- JWT tokens expire after 72 hours
- Email addresses are automatically converted to lowercase
- Passwords are never returned in responses
- All timestamps are in Unix format 