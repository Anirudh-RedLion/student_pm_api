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

## Environment Variables

This project is configured via environment variables for maximum portability and security. See [docs/ENVIRONMENT.md](docs/ENVIRONMENT.md) for a full list and usage instructions.

- Copy `.env.example` to `.env` and set your values for local development.
- In production, set these variables in your deployment environment. 

## Docker Usage

This project includes a Dockerfile for easy deployment anywhere.

### Build the Docker image
```bash
docker build -t student-pm-api .
```

### Run the Docker container
```bash
docker run --env-file .env -p 8080:8080 student-pm-api
```
- The `--env-file .env` flag loads environment variables from your `.env` file.
- The `-p 8080:8080` flag maps the container port to your local machine.

### Overriding Environment Variables
You can override any variable at runtime:
```bash
docker run -e PORT=9090 -e JWT_SECRET=anothersecret -p 9090:9090 student-pm-api
```

See [docs/ENVIRONMENT.md](docs/ENVIRONMENT.md) for all available variables. 