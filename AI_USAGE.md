# AI Usage

## Tools Used
- ChatGPT (GPT-4) via ChatGPT Plus
- Github Copilot

---

## Prompts Used

I used AI to assist with:

- "How to retry Postgres connection using pgx in Go"

- "Dockerize a Go Fiber app and connect it with Postgres using docker-compose"
- "Add pagination to a GET route in Fiber using query parameters"
- "Fix Go scan error: cannot scan date into string"

---

## How Output Was Used 


- **DB Connection (pgx)**: AI helped generate a retry loop that waits for Postgres to start. I adapted the logic to work with Docker Compose by handling connect: connection refused errors.

- **Docker Fixes**: When Docker builds failed (e.g., Go 1.24 issue), I asked for compatibility fixes and integrated the solution by adjusting go.mod and Docker base image.

- **Pagination**: I asked how to implement pagination using LIMIT and OFFSET in Fiber. I tweaked it to include metadata in the JSON response (e.g., current page, results count).

- **Error Debugging**: AI helped explain scan errors (e.g., cannot scan date into string). I applied the fix by switching to time.Time and restructured the model fields accordingly.

---

## Github Copilot

- **Writing error logs**
  - Writing repetitive log.Println statements for error tracing
  - Auto-completing if err != nil blocks
  - Suggesting struct field names when creating JSON models
  - Quickly generating consistent handler function signatures

I made sure to review autofilled code after to ensure clarity and match the over structure of the task