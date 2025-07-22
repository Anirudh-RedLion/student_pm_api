# Deployment Guide: Railway & Render

This guide explains how to deploy your Student PM API to two popular cloud platforms: **Railway** and **Render**. Both support Docker and environment variables, making them ideal for portable, multi-environment deployments.

---

## 1. Deploying to Railway

### A. Prerequisites
- GitHub account
- Railway account ([Sign up free](https://railway.app/))

### B. Steps
1. **Push your code to GitHub** (if not already).
2. **Go to [Railway Dashboard](https://railway.app/dashboard)** and click **"New Project"**.
3. **Select "Deploy from GitHub repo"** and connect your repository.
4. **Railway will auto-detect your Dockerfile.**
5. **Set environment variables:**
   - Go to the "Variables" tab.
   - Add `PORT`, `DB_PATH`, and `JWT_SECRET` (see [docs/ENVIRONMENT.md](ENVIRONMENT.md)).
6. **Deploy!**
   - Railway will build and run your app.
   - You’ll get a public URL (e.g., `https://student-pm-api.up.railway.app`).
7. **Update:**
   - Push changes to GitHub; Railway auto-redeploys.

### C. Tips
- Use Railway’s persistent volumes if you want your SQLite DB to persist across deploys.
- Free tier may sleep after inactivity.
- [Railway Docs](https://docs.railway.app/)

---

## 2. Deploying to Render

### A. Prerequisites
- GitHub account
- Render account ([Sign up free](https://render.com/))

### B. Steps
1. **Push your code to GitHub** (if not already).
2. **Go to [Render Dashboard](https://dashboard.render.com/)** and click **"New Web Service"**.
3. **Connect your GitHub repo.**
4. **Select "Docker" as the environment.**
5. **Set environment variables:**
   - In the "Environment" tab, add `PORT`, `DB_PATH`, and `JWT_SECRET`.
6. **Deploy!**
   - Render will build and run your app.
   - You’ll get a public URL (e.g., `https://student-pm-api.onrender.com`).
7. **Update:**
   - Push changes to GitHub; Render auto-redeploys.

### C. Tips
- Use Render’s "Persistent Disk" if you want your SQLite DB to persist across deploys.
- Free tier may sleep after inactivity.
- [Render Docs](https://render.com/docs)

---

## 3. General Cloud Deployment Notes
- Always set a strong, unique `JWT_SECRET` in production.
- Use different DB files or databases for each environment.
- Monitor your app’s health and logs via the platform dashboard.
- Both platforms support custom domains and HTTPS.

---

## 4. Troubleshooting
- **App won’t start?** Check logs for missing env vars or port conflicts.
- **DB not persisting?** Use persistent storage/volumes.
- **Quota exhausted?** Deploy to a second vendor and update your DNS or API gateway.

---

## 5. References
- [docs/ENVIRONMENT.md](ENVIRONMENT.md)
- [docs/MULTI_ENVIRONMENT.md](MULTI_ENVIRONMENT.md)
- [Railway Docs](https://docs.railway.app/)
- [Render Docs](https://render.com/docs) 