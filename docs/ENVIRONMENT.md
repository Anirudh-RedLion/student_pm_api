# Environment Variables for Student PM API

This backend service is configured entirely via environment variables. This makes it easy to deploy to any cloud or local environment.

## Required Variables

| Variable     | Description                                      | Example Value           | Default         |
|--------------|--------------------------------------------------|------------------------|-----------------|
| PORT         | Port for the HTTP server                         | 8080                   | 8080            |
| DB_PATH      | Path to the SQLite database file                  | studentpm.db           | studentpm.db    |
| JWT_SECRET   | Secret key for signing JWT tokens (required!)     | my_super_secret_key    | (no default)    |

## Using a `.env` File (Local Development)

1. Copy `.env.example` to `.env` in the `student_pm_api` directory:
   ```bash
   cp .env.example .env
   ```
2. Edit `.env` and set your values.
3. The app will automatically load `.env` when running locally.

## Example `.env` File
```env
PORT=8080
DB_PATH=studentpm.db
JWT_SECRET=your_super_secret_key_here
```

## Notes
- **Never commit your real `.env` file with secrets to version control.**
- In production, set these variables in your cloud provider’s dashboard or deployment pipeline.
- If `JWT_SECRET` is not set, the app will not start (for security). 