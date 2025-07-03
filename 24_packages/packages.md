**Packages in Go**, are a core part of how Go structures code, promotes reusability, and keeps things organized.

---

## 📦 What Is a Package in Go?

In Go, a **package** is a collection of Go source files in the same directory that are compiled together. Every Go file **must belong to a package**.

A Go program is made up of many packages:

* **Standard library packages** (like `fmt`, `os`, `net/http`)
* **Your own custom packages**
* **Third-party packages** (via `go get` or `go mod`)

---

## 🗂️ Package Anatomy

```go
// greetings.go
package greetings

import "fmt"

func Hello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
```

* `package greetings`: defines the package name
* This file is part of the `greetings` package
* We can now use `greetings.Hello("Skyy")` in another package

---

## 🏁 `main` Package

```go
package main

import (
    "fmt"
    "yourapp/greetings"
)

func main() {
    fmt.Println(greetings.Hello("Skyy"))
}
```

* `main` is **the entry point** of every Go program.
* It must have a `main()` function.
* Without the `main` package, a Go program **cannot run**, only compile into a library.

---

## ✅ Naming Rules

* All `.go` files in the same directory must declare the **same package name**
* A package name should be **short and descriptive**
* We typically name the package **after the folder** it’s in

---

## 🔐 Exported vs Unexported

Go uses **capitalization** to determine visibility.

| Identifier | Visibility           |
| ---------- | -------------------- |
| `Hello`    | Exported (public)    |
| `hello`    | Unexported (private) |

This applies to:

* Functions
* Variables
* Types
* Constants

So:

```go
func Hello() {}  // can be accessed from other packages
func hello() {}  // internal only
```

---

## 📁 Example: Creating and Using a Custom Package

### Folder Structure

```
goapp/
├── main.go
└── utils/
    └── mathutil.go
```

---

### utils/mathutil.go

```go
package utils

func Add(a, b int) int {
	return a + b
}

func subtract(a, b int) int { // not exported
	return a - b
}
```

---

### main.go

```go
package main

import (
	"fmt"
	"goapp/utils"
)

func main() {
	sum := utils.Add(5, 7)
	fmt.Println("Sum:", sum)

	// utils.subtract() → ❌ can't access unexported function
}
```

---

## 📦 Import Path

Import paths depend on:

* **Your Go module name** (declared in `go.mod`)
* **File structure**

```sh
go mod init goapp
```

Now `utils` is accessible as `goapp/utils`.

---

## 🧱 Go Modules and Packages

Go modules (`go.mod`) tell Go where our project starts and handles dependencies.

**Example `go.mod`:**

```
module goapp

go 1.22
```

This makes our imports structured like:

```go
import "goapp/utils"
```

---

## 🔍 How Go Resolves Packages

Go looks for imports in this order:

1. Standard library (`fmt`, `os`, etc.)
2. Our module-defined packages (based on `go.mod`)
3. Third-party packages in the Go module cache (downloaded via `go get`)

---

## 📚 Standard Library Packages (a few examples)

| Package         | Description                      |
| --------------- | -------------------------------- |
| `fmt`           | Formatting and printing          |
| `os`            | OS-level file & process handling |
| `io`            | Basic I/O interfaces             |
| `bufio`         | Buffered I/O                     |
| `net/http`      | Web servers and HTTP clients     |
| `encoding/json` | JSON encoding/decoding           |
| `time`          | Time and date utilities          |
| `math`          | Mathematical operations          |

We can browse them at: [https://pkg.go.dev/std](https://pkg.go.dev/std)

---

## ⚙️ Internal Packages (Restricted Access)

Go allows you to create **internal** packages that can only be used within the same module.

**Folder Structure:**

```
goapp/
├── internal/
│   └── db/
│       └── connection.go
├── main.go
```

Anything inside `/internal/...` can **only** be imported by packages under the same module root.

✅ Used to **hide implementation details** (clean architecture)

---

## 🧪 Testing Packages

Go supports testing with a built-in tool.

```go
// utils/mathutil_test.go
package utils

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}
```

Then run:

```sh
go test ./...
```

---

## 🧠 Best Practices

| Practice                    | Why It Matters                            |
| --------------------------- | ----------------------------------------- |
| Group related functions     | Packages should be cohesive               |
| Keep small & focused        | Don't overload a single package           |
| Use `internal/` for privacy | Prevent misuse by other modules           |
| Name packages simply        | Avoid `util`, prefer specific names       |
| Avoid cyclic dependencies   | Go will refuse to compile if cycles exist |

---

## ✅ Summary

| Concept                | Details                                               |
| ---------------------- | ----------------------------------------------------- |
| `package`              | Declares the namespace of the file                    |
| `main` package         | Required to run an app                                |
| Exported vs unexported | Uppercase = exported, lowercase = internal            |
| Imports                | Based on path, module, or standard library            |
| Go modules             | Helps manage versions, dependencies, and import paths |
| Internal package       | Used to restrict usage to same module                 |

---
