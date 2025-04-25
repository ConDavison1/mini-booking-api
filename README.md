#  RISE API (Go + PostgreSQL)

 Built with Go (Fiber web framework), PostgreSQL, Docker, and JWT-based authentication.

 Used Fiber because of its performance (its one of the fastest Go frameworks because of fasthttp) and its great for APIs. It comes with middleware for JWT and CORS. Docker/Cloud Ready also

---

## Setup Instructions (Local)

```bash
git clone https://github.com/ConDavison1/rise-api.git
cd rise-api
docker-compose up --build
```

API will be live at: http://localhost:3000

---

## Dockerized Architecture

| Service | Description        | Port  |
|---------|--------------------|-------|
| api   | Go backend (Fiber) | 3000  |
| db    | PostgreSQL         | 5432  |

Docker handles:
- Go app build
- Postgres container w/ seeded data
- .env-based config

---

## Authentication

Login using:

```
POST /login
{
  "email": "admin@example.com",
  "password": "password123"
}
```

Use the returned JWT for all /bookings routes:

```
Authorization: Bearer <your_token_here>
```

---

## API Endpoints

| Method | Route              | Description                      |
|--------|--------------------|----------------------------------|
| GET    | `/programs`        | List all programs                |
| GET    | `/programs/:id`    | Get program + spots left         |
| POST   | `/login`           | Auth user + return JWT           |
| GET    | `/bookings`        | List bookings (JWT required)     |
| POST   | `/bookings`        | Book a program (JWT required)    |

---



## GCP Deployment

### Deployment via Cloud Run:

The reason for cloud run is because its really good for containerized Go app and has autoscaling support

### What you need to have 

- Google Cloud SDK installed 
- Billing enabled 
- Docker installed

### Steps

```bash
# 1. Submit container image to GCP
gcloud builds submit --tag gcr.io/PROJECT_ID/rise-api
``` 
This uploads the local Dockerfile-based Go app.
```bash

# 2. Deploy the service
gcloud run deploy rise-api \
  --image gcr.io/PROJECT_ID/booking-api \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars DB_HOST=...,DB_USER=...,JWT_SECRET=...
```

### Env Vars and Secrets

Dont hard code secrets always use

- **Secret Manager** for:
  - `DB_PASSWORD`
  - `JWT_SECRET`
- **Cloud SQL IAM connector** or **private IP connection** for:
  - `DB_HOST`

You can inject secrets at runtime using:
```bash
--set-secrets JWT_SECRET=projects/your-project/secrets/jwt-secret:latest
```

## Scaling and Monitoring the app

- Use **Cloud Monitoring** to set alerts (e.g., error spike, latency)
- Enable **Cloud Trace** to analyze request performance
- Cloud Run autoscaling handles demand bursts:
  - Configure min/max instances
  - Concurrency settings (default: 80)
 

