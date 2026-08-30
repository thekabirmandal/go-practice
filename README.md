# Go Practice & Learning Roadmap

Welcome to your Go practice workspace! This project is initialized with Go modules, **Gin** web framework, and **Google UUID**.

---

## 🚀 Quick Start Commands

```powershell
# Run the current development server
go run .

# Build executable binary (outputs to .gitignore ignored binary)
go build -o app.exe .

# Tidy dependencies
go mod tidy
```

### Current Endpoints
- `GET http://localhost:8080/ping` &rarr; `{"message": "pong"}`
- `GET http://localhost:8080/uuid` &rarr; `{"uuid": "<generated-uuid>"}`

---

## 📋 Next Practice Sessions

Here are high-value practical topics ready for your next session:

### 1. In-Memory CRUD API (Recommended First Step)
- **Goal:** Build a complete RESTful API for managing resources (e.g. `Todo`, `User`, `Product`).
- **Key Concepts:**
  - Struct definition with JSON tags (`json:"title" binding:"required"`).
  - Pointers and slice manipulation (`append`, delete by index).
  - HTTP verbs: `GET`, `POST`, `PUT`, `DELETE`.
  - Route parameters (`c.Param("id")`) and Query parameters (`c.Query("search")`).

---

### 2. Custom Gin Middleware
- **Goal:** Intercept and process requests before reaching route handlers.
- **Key Concepts:**
  - **Request Logger:** Measuring execution latency with `time.Since()`.
  - **Auth Guard:** Verifying `Authorization: Bearer <token>` header.
  - **Error Recovery:** Handling panics gracefully with `c.AbortWithStatusJSON()`.

---

### 3. Production Project Structure
- **Goal:** Modularize code across packages instead of a single `main.go`.
- **Target Architecture:**
  ```text
  go-practice/
  ├── cmd/
  │   └── api/
  │       └── main.go       # App entrypoint
  ├── internal/
  │   ├── handlers/         # Gin HTTP handlers
  │   ├── models/           # Data structs
  │   └── services/         # Business logic
  ├── .gitignore
  ├── go.mod
  └── go.sum
  ```

---

### 4. Database Persistence (SQLite + GORM)
- **Goal:** Persist data to a local file database using an ORM.
- **Packages:** `gorm.io/gorm` and `gorm.io/driver/sqlite`.
- **Key Concepts:** `AutoMigrate()`, `Create()`, `Find()`, `First()`, `Updates()`, `Delete()`.

---

### 5. Go Concurrency & Goroutines
- **Goal:** Master concurrent execution and channel communication.
- **Key Concepts:**
  - Spawning goroutines with `go func()`.
  - Channels (`chan string`, `chan int`) for message passing.
  - `sync.WaitGroup` to wait for multiple workers to complete.
  - `select` statements and worker pool patterns.

---

### 6. Automated Unit & HTTP Testing
- **Goal:** Test API endpoints without running a live server.
- **Key Concepts:**
  - Standard Go testing (`testing.T`).
  - `net/http/httptest` with `httptest.NewRecorder()`.
  - Running tests: `go test -v ./...`.

---

## 🛠 Useful Tools to Add
- **Task Runner:** `Taskfile.yml` or `Makefile` for custom scripts.
- **Live Reload:** `Air` (`go install github.com/air-verse/air@latest`) for auto-reloading on file save.
