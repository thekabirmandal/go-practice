# Go Language Fundamentals & Notes

---

## 1. Toolchain & Essential Commands

The Go toolchain provides built-in tools for building, formatting, testing, and managing dependencies without needing third-party task runners.

| Command | Purpose | Example |
| :--- | :--- | :--- |
| `go run` | Compiles and executes code in memory (useful during development) | `go run main.go` |
| `go build` | Compiles packages and dependencies into a standalone binary | `go build -o app.exe main.go` |
| `go fmt` | Automatically formats Go source files according to Go style guides | `go fmt ./...` |
| `go vet` | Examines Go source code and reports suspicious constructs / bugs | `go vet ./...` |
| `go test` | Runs unit and benchmark tests (files ending in `_test.go`) | `go test ./...` |
| `go mod tidy` | Adds missing module requirements and removes unused ones | `go mod tidy` |
| `go get` | Downloads and adds/updates dependencies in `go.mod` | `go get github.com/gin-gonic/gin` |

### Compilation Examples
```bash
# Build the root project binary
go build -o bin/server.exe .

# Build a specific package/command path
go build -o bin/api.exe ./cmd/api/v1
```

---

## 2. Go Modules & Dependency Management

Go uses **Go Modules** as its official dependency management system.

* **`go.mod`**: Describes the module's path, Go language version, and direct/indirect dependencies.
* **`go.sum`**: Contains expected cryptographic checksums of the content of specific module versions to ensure reproducibility and security.

### Common Workflow
```bash
# 1. Initialize a new Go module
go mod init github.com/username/project-name

# 2. Install an external package
go get github.com/google/uuid

# 3. Clean up and synchronize dependencies
go mod tidy

# 4. Download dependencies to local cache
go mod download
```

---

## 3. Basic Program Structure

Every runnable Go program starts with `package main` and contains an entrypoint function `func main()`.

### Example: `main.go`

```go
// 1. Package Declaration
// Every Go file must belong to a package. 'package main' produces an executable binary.
package main

// 2. Import Declarations
// Imports standard library packages or third-party dependencies.
import (
	"fmt"
	"net/http"
)

// 3. Package-level Variables / Constants
const appPort = ":8080"

// 4. Main Entrypoint Function
// Execution starts and ends here. Takes no arguments and returns no value.
func main() {
	fmt.Printf("Server starting on port %s...\n", appPort)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "pong"}`))
	})

	if err := http.ListenAndServe(appPort, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
```

### Key Conventions
* **Package Name**: Usually matches the directory name (except `package main` for binaries).
* **Exported vs. Unexported Identifiers**:
  * Capital first letter = **Exported** (public, accessible outside the package, e.g., `fmt.Println`, `http.StatusOK`).
  * Lowercase first letter = **Unexported** (package-private, e.g., `appPort`, `calculateTotal()`).

---

## 4. Core Language Syntax Cheatsheet

### Variable Declarations
```go
// Short variable declaration (only inside functions)
name := "Go Developer"
age := 25

// Explicit declaration with type
var count int = 10
var isReady bool // Default zero value: false

// Grouped declaration
var (
	host = "localhost"
	port = 5432
)
```

### Error Handling Pattern
Go functions explicitly return errors as their last return value instead of throwing exceptions:
```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Idiomatic error checking
result, err := divide(10, 2)
if err != nil {
	log.Fatalf("Operation failed: %v", err)
}
fmt.Println("Result:", result)
```