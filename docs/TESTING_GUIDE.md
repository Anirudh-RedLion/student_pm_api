# API Testing Guide

## Quick Start

### 1. Start the API Server
```bash
cd student_pm_api
go run main.go
```

The server will start on `http://localhost:8080`

### 2. Test User Registration
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "role": "student"
  }'
```

**Expected Response:**
```json
{"message":"User registered"}
```

### 3. Test User Login
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "password123"
  }'
```

**Expected Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "role": "student"
  }
}
```

### 4. Test Protected Endpoint
```bash
curl -X GET http://localhost:8080/api/users/me \
  -H "Authorization: Bearer YOUR_JWT_TOKEN_HERE"
```

**Expected Response:**
```json
{
  "id": 1,
  "email": "john@example.com",
  "role": "student"
}
```

## Error Testing

### Test Duplicate Registration
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Another User",
    "email": "john@example.com",
    "password": "password123",
    "role": "student"
  }'
```

**Expected Response:**
```json
{"error":"Email already exists"}
```

### Test Wrong Password
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "john@example.com",
    "password": "wrongpassword"
  }'
```

**Expected Response:**
```json
{"error":"Invalid credentials"}
```

### Test Invalid Token
```bash
curl -X GET http://localhost:8080/api/users/me \
  -H "Authorization: Bearer invalid_token"
```

**Expected Response:**
```json
{"error":"Invalid token"}
```

## Using Postman

### 1. User Registration
- **Method**: POST
- **URL**: `http://localhost:8080/api/register`
- **Headers**: `Content-Type: application/json`
- **Body** (raw JSON):
```json
{
  "name": "Test User",
  "email": "test@example.com",
  "password": "password123",
  "role": "student"
}
```

### 2. User Login
- **Method**: POST
- **URL**: `http://localhost:8080/api/login`
- **Headers**: `Content-Type: application/json`
- **Body** (raw JSON):
```json
{
  "email": "test@example.com",
  "password": "password123"
}
```

### 3. Get Current User
- **Method**: GET
- **URL**: `http://localhost:8080/api/users/me`
- **Headers**: 
  - `Content-Type: application/json`
  - `Authorization: Bearer YOUR_JWT_TOKEN`

## Test Script

Create a file called `test_api.sh`:

```bash
#!/bin/bash

echo "Testing Student PM API"
echo "======================"

# Test Registration
echo "1. Testing User Registration..."
REGISTER_RESPONSE=$(curl -s -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "role": "student"
  }')
echo "Registration Response: $REGISTER_RESPONSE"

# Test Login
echo -e "\n2. Testing User Login..."
LOGIN_RESPONSE=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }')
echo "Login Response: $LOGIN_RESPONSE"

# Extract token from login response
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*"' | cut -d'"' -f4)

# Test Protected Endpoint
echo -e "\n3. Testing Protected Endpoint..."
ME_RESPONSE=$(curl -s -X GET http://localhost:8080/api/users/me \
  -H "Authorization: Bearer $TOKEN")
echo "Protected Endpoint Response: $ME_RESPONSE"

echo -e "\nTesting Complete!"
```

Make it executable and run:
```bash
chmod +x test_api.sh
./test_api.sh
```

## Troubleshooting

### Common Issues

1. **Port already in use**
   ```bash
   lsof -ti:8080 | xargs kill -9
   ```

2. **Database locked**
   ```bash
   rm studentpm.db
   go run main.go
   ```

3. **Import errors**
   ```bash
   go mod tidy
   go build ./...
   ```

### Check Server Status
```bash
curl http://localhost:8080/api/register
```
If the server is running, you should get a response (even if it's an error for missing body). 