# Student PM API (Go Backend)

This repository contains the Go backend for the Student Profile Management platform. It exposes RESTful APIs for authentication, student/course/job/query management, analytics, and more.

## Project Structure

```
student_pm_api/
├── main.go           # Entry point
├── go.mod            # Go module definition
├── handlers/         # HTTP handlers for each resource (students, courses, jobs, etc.)
├── models/           # Data models (structs for DB and JSON)
├── db/               # Database connection, queries, migrations
├── middleware/       # Auth, logging, CORS, etc.
├── routes/           # Route registration and grouping
├── utils/            # Helpers (JWT, hashing, validation, etc.)
```

## Key Features
- JWT authentication
- RESTful endpoints for all resources
- Role-based access control
- Clean, scalable architecture
- Ready for deployment (Docker, cloud, etc.)

## Next Steps
- Scaffold handlers, models, and routes for core resources
- Implement authentication and sample endpoints
- Connect to a database (Postgres/MySQL)
- Add API documentation 