# Multi-Environment Deployment Guide

This guide explains how to set up and manage multiple environments (development, staging, production) for your Student PM API using environment variables and Docker.

---

## 1. Why Multi-Environment?
- **Development**: Local testing, debugging, rapid iteration.
- **Staging**: Pre-production, QA, integration testing.
- **Production**: Live, real users, highest security.

Each environment may use different settings (database, secrets, ports, etc.).

---

## 2. Environment Variable Files

- `.env.development` — for local dev
- `.env.staging` — for staging server
- `.env.production` — for production

**Example:**
```env
# .env.development
PORT=8080
DB_PATH=studentpm-dev.db
JWT_SECRET=dev_secret

# .env.staging
PORT=8080
DB_PATH=studentpm-staging.db
JWT_SECRET=staging_secret

# .env.production
PORT=80
DB_PATH=studentpm.db
JWT_SECRET=super_secret_prod_key
```

---

## 3. Local Development

- Copy `.env.development` to `.env`:
  ```bash
  cp .env.development .env
  ```
- Run the app (Go or Docker):
  ```bash
  go run main.go
  # or
  docker run --env-file .env -p 8080:8080 student-pm-api
  ```

---

## 4. Staging & Production

- On your staging/production server or cloud platform, set the appropriate environment variables.
- For Docker, upload the correct `.env.staging` or `.env.production` and run:
  ```bash
  docker run --env-file .env.staging -p 8080:8080 student-pm-api
  # or
  docker run --env-file .env.production -p 80:80 student-pm-api
  ```
- **Cloud platforms** (Railway, Render, Fly.io, etc.) let you set env vars in their dashboard or CI/CD pipeline.

---

## 5. Best Practices

- **Never commit real secrets to git.** Only commit `.env.example` and sample files.
- **Use strong, unique JWT secrets in production.**
- **Use different DB files or databases for each environment.**
- **Automate environment selection in your CI/CD pipeline.**
- **Document which environment is running where.**

---

## 6. Example Workflow

1. **Local Dev:**
   - Use `.env.development` locally.
2. **Staging:**
   - Deploy to a staging server with `.env.staging`.
3. **Production:**
   - Deploy to production with `.env.production`.
4. **Switching Environments:**
   - Change the `--env-file` flag or set env vars in your cloud dashboard.

---

## 7. Advanced: Multi-Cloud & Failover

- Deploy the same Docker image to multiple vendors (Railway, Render, Fly.io, etc.) with different env files.
- Use DNS or a load balancer to route traffic as needed.
- Monitor health and switch traffic if one environment is down or quota is exhausted.

---

## 8. References
- [docs/ENVIRONMENT.md](ENVIRONMENT.md)
- [README.md](../README.md) 