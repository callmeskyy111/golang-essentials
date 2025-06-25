### **What is GoLang?** 🔵

Go (or GoLang) is an open-source, statically typed, compiled programming language developed by Google in 2007 and released in 2009. It was designed to improve productivity in software development, combining the performance of a compiled language with the simplicity and ease of use of dynamically typed languages.

Go is particularly known for its efficiency in handling concurrent programming, making it ideal for building scalable and high-performance applications.

---

### **Key Features of GoLang:**

1. **Simple and Readable Syntax**

   * Go has a clean and minimalistic syntax, making it easy to learn and read.
   * No complex features like classes, inheritance, or explicit memory management.

2. **Compiled and High Performance**

   * Unlike interpreted languages (e.g., Python, JavaScript), Go is compiled, resulting in faster execution.
   * Produces standalone binaries with minimal dependencies.

3. **Garbage Collection (GC)**

   * Go has an efficient garbage collector that manages memory automatically.

4. **Concurrency Support (Goroutines)**

   * Go has built-in concurrency using lightweight threads called **goroutines**.
   * It uses **channels** for safe communication between goroutines, making it more efficient than traditional threading.

5. **Static Typing with Type Inference**

   * Statically typed, reducing runtime errors.
   * Also supports type inference, making code more concise.

6. **Cross-Platform**

   * Go supports cross-platform development.
   * A single Go program can be compiled to run on Windows, Linux, and macOS without changes.

7. **Standard Library**

   * Comes with a powerful standard library, eliminating the need for third-party packages for common tasks like HTTP handling, JSON parsing, and cryptography.

8. **Built-in Testing**

   * Provides an in-built testing framework (`testing` package) for unit testing and benchmarking.

9. **Efficient Networking & Web Development**

   * Optimized for building high-performance web services and APIs.
   * Used in popular frameworks like Gin and Fiber.

10. **Security & Simplicity**

* Avoids issues like dangling pointers and memory leaks.
* Enforces good coding practices and prevents unsafe operations.

---

### **Why Should We Learn GoLang?**

1. **High Demand for Go Developers**

   * Companies like Google, Uber, Dropbox, Netflix, and PayPal use Go.
   * Growing job opportunities in backend development, cloud computing, and DevOps.

2. **Ideal for Backend Development**

   * Perfect for building RESTful APIs, microservices, and backend systems.
   * Frameworks like Gin, Fiber, and Echo make web development easy.

3. **Great for Cloud Computing & DevOps**

   * Used in cloud-native technologies like Docker, Kubernetes, and Terraform.
   * Many DevOps tools are built using Go.

4. **Concurrency & Scalability**

   * Handles thousands of concurrent tasks efficiently.
   * Suitable for applications requiring real-time data processing (e.g., chat apps, streaming services).

5. **Ease of Learning**

   * Simple syntax with fewer concepts to master compared to Java or C++.
   * Good starting point for learning system programming.

6. **Performance & Efficiency**

   * Almost as fast as C/C++, with better memory management.
   * Lower CPU and RAM usage, making it ideal for large-scale applications.

7. **Community & Ecosystem**

   * Active open-source community.
   * Strong ecosystem with numerous libraries and frameworks.

---

### **Should We Learn GoLang?**

Yes, especially if we're interested in:
✅ **Backend development** (APIs, microservices)
✅ **Cloud computing & DevOps**
✅ **Building high-performance, scalable applications**
✅ **Concurrency & distributed systems**
✅ **Learning a simple yet powerful language with fast execution**

Here’s a **detailed explanation** of all the Go commands:

---

### **1. `go bug`**

* Opens a web page to report a bug in the Go project.
* It automatically collects system information to help with debugging.

📝 **Usage:**

```sh
go bug
```

---

### **2. `go build`**

* Compiles the Go source code and generates an executable binary.
* It compiles all dependencies but does not install them.

📝 **Usage:**

```sh
go build [package]
```

Example:

```sh
go build main.go
```

* This compiles `main.go` into an executable file (`main.exe` on Windows, `main` on Linux/macOS).

---

### **3. `go clean`**

* Removes compiled object files and cached data from the Go build system.
* Useful for cleaning up unnecessary files.

📝 **Usage:**

```sh
go clean
```

* Removes all build artifacts from the current module.

```sh
go clean -cache -modcache -testcache
```

* Cleans the cache, module cache, and test cache.

---

### **4. `go doc`**

* Displays documentation for a Go package or function.

📝 **Usage:**

```sh
go doc fmt.Println
```

* Shows documentation for `fmt.Println`.

```sh
go doc -all fmt
```

* Shows full documentation for the `fmt` package.

---

### **5. `go env`**

* Prints information about the Go environment variables (like `GOROOT`, `GOPATH`).

📝 **Usage:**

```sh
go env
```

* Shows all environment variables.

```sh
go env GOPATH
```

* Prints only the `GOPATH` value.

---

### **6. `go fix`**

* Automatically updates Go code to use the latest API changes.

📝 **Usage:**

```sh
go fix ./...
```

* Fixes all Go files in the current directory and subdirectories.

---

### **7. `go fmt`**

* Formats Go source code to follow standard Go coding style.

📝 **Usage:**

```sh
go fmt main.go
```

* Formats `main.go`.

```sh
go fmt ./...
```

* Formats all `.go` files in the current directory and subdirectories.

---

### **8. `go generate`**

* Runs code generation tools specified in Go files using `//go:generate` comments.
* Often used for generating boilerplate code, mocks, or database models.

📝 **Usage:**

```sh
go generate
```

* Executes commands specified in `//go:generate` comments.

Example in a Go file:

```go
//go:generate go run mygenerator.go
```

---

### **9. `go get`**

* Fetches and installs dependencies for a Go module.

📝 **Usage:**

```sh
go get github.com/gorilla/mux
```

* Installs the `gorilla/mux` package.

```sh
go get -u
```

* Updates all dependencies in the project.

---

### **10. `go install`**

* Compiles and installs a package or program into `$GOBIN`.
* Useful for installing CLI tools.

📝 **Usage:**

```sh
go install github.com/user/tool@latest
```

* Installs the latest version of `tool`.

```sh
go install ./...
```

* Installs all executables from the current module.

---

### **11. `go list`**

* Lists installed Go packages or modules.

📝 **Usage:**

```sh
go list all
```

* Lists all available packages.

```sh
go list -m
```

* Shows details about the current module.

---

### **12. `go mod`**

* Manages Go modules (Go’s dependency management system).

📝 **Common Subcommands:**

```sh
go mod init myproject
```

* Initializes a new module.

```sh
go mod tidy
```

* Removes unused dependencies from `go.mod`.

```sh
go mod verify
```

* Checks the integrity of dependencies.

---

### **13. `go work`**

* Manages workspaces containing multiple Go modules.

📝 **Usage:**

```sh
go work init
```

* Initializes a workspace file (`go.work`).

```sh
go work use ./module1 ./module2
```

* Adds multiple modules to the workspace.

---

### **14. `go run`**

* Compiles and runs a Go program **without creating an executable file**.

📝 **Usage:**

```sh
go run main.go
```

* Runs `main.go` without generating an executable.

```sh
go run .
```

* Runs all Go files in the current directory.

---

### **15. `go telemetry`**

* Manages telemetry data sent to Go developers.

📝 **Usage:**

```sh
go telemetry on
```

* Enables telemetry.

```sh
go telemetry off
```

* Disables telemetry.

---

### **16. `go test`**

* Runs unit tests in the current package.

📝 **Usage:**

```sh
go test
```

* Runs all tests in the current directory.

```sh
go test -v
```

* Runs tests in verbose mode.

```sh
go test -cover
```

* Shows code coverage details.

---

### **17. `go tool`**

* Runs various Go tools like `pprof` (profiling), `vet` (error checking), etc.

📝 **Usage:**

```sh
go tool vet main.go
```

* Checks for possible mistakes in `main.go`.

```sh
go tool pprof profile.out
```

* Analyzes a performance profile.

---

### **18. `go version`**

* Prints the installed Go version.

📝 **Usage:**

```sh
go version
```

* Example output:

  ```
  go version go1.21.0 windows/amd64
  ```

---

### **19. `go vet`**

* Detects potential bugs and incorrect code patterns.

📝 **Usage:**

```sh
go vet main.go
```

* Checks `main.go` for issues.

```sh
go vet ./...
```

* Checks all files in the current project.

---

### **Summary Table:**

| Command        | Description                          |
| -------------- | ------------------------------------ |
| `go bug`       | Reports a Go bug.                    |
| `go build`     | Compiles Go code into an executable. |
| `go clean`     | Removes compiled files and cache.    |
| `go doc`       | Shows package documentation.         |
| `go env`       | Displays Go environment settings.    |
| `go fix`       | Updates Go code to new API versions. |
| `go fmt`       | Formats Go source files.             |
| `go generate`  | Runs code generation tools.          |
| `go get`       | Fetches and installs dependencies.   |
| `go install`   | Installs executables in `$GOBIN`.    |
| `go list`      | Lists installed packages.            |
| `go mod`       | Manages Go modules.                  |
| `go work`      | Manages workspaces.                  |
| `go run`       | Compiles and runs a Go program.      |
| `go telemetry` | Manages telemetry settings.          |
| `go test`      | Runs unit tests.                     |
| `go tool`      | Runs additional Go tools.            |
| `go version`   | Prints the installed Go version.     |
| `go vet`       | Detects code mistakes.               |

---

This covers all the **Go commands** in detail.

### 1. Hello World 👋🏻
- go build 1_hello_world/main.go
- ./main (.exe file)
- 💡 Soln. - go run 1_hello_world/main.go

### 2. All the basic data types in Go (Golang) 💻

## ✅ 1. **Boolean Type (`bool`)**

* Represents a truth value: `true` or `false`
* Default value: `false`

```go
var isActive bool = true
var isLoggedIn = false // inferred type
```

> Use in logical expressions, conditions, or control flow (e.g., `if`, `for`).

---

## ✅ 2. **Numeric Types**

### A. **Integer Types**

Go supports both **signed** and **unsigned** integers.

#### 🔹 Signed Integers:

| Type    | Size                                      | Range                           |
| ------- | ----------------------------------------- | ------------------------------- |
| `int8`  | 8-bit                                     | -128 to 127                     |
| `int16` | 16-bit                                    | -32,768 to 32,767               |
| `int32` | 32-bit                                    | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64-bit                                    | ±9.22×10^18                     |
| `int`   | 32 or 64-bit (depends on OS/architecture) |                                 |

#### 🔹 Unsigned Integers:

| Type     | Size                      | Range                     |
| -------- | ------------------------- | ------------------------- |
| `uint8`  | 8-bit                     | 0 to 255                  |
| `uint16` | 16-bit                    | 0 to 65,535               |
| `uint32` | 32-bit                    | 0 to 4,294,967,295        |
| `uint64` | 64-bit                    | 0 to 18,446,744,073×10^18 |
| `uint`   | 32 or 64-bit (like `int`) |                           |

#### Example:

```go
var a int8 = -100
var b uint16 = 65000
```

> ⚠️ Always choose the smallest type that fits your range to optimize memory.

---

### B. **Floating Point Types**

Used for **decimal or real numbers**.

| Type      | Size   | Precision           |
| --------- | ------ | ------------------- |
| `float32` | 32-bit | \~6 decimal places  |
| `float64` | 64-bit | \~15 decimal places |

#### Example:

```go
var pi float64 = 3.14159265358979
var price float32 = 19.99
```

> Use `float64` by default for better precision.

---

### C. **Complex Numbers**

Go supports complex numbers with real and imaginary parts.

| Type         | Description                      |
| ------------ | -------------------------------- |
| `complex64`  | float32 real and imaginary parts |
| `complex128` | float64 real and imaginary parts |

#### Example:

```go
var c1 complex64 = 1 + 2i
var c2 = complex(5.5, 3.3) // real: 5.5, imag: 3.3
```

> Rarely used, but handy in scientific computing.

---

## ✅ 3. **String Type**

A **string** is a sequence of bytes (UTF-8 encoded characters).

* Immutable
* Can be indexed and sliced

```go
var name string = "Skyy"
var greeting = "Hello, " + name
```

#### String Operations:

```go
fmt.Println(len(name))      // Length
fmt.Println(name[0])        // First byte
fmt.Println(name[0:2])      // Slice "Sk"
```

> Strings are byte-based; to work with Unicode characters (like emojis), use `rune`.

---

## ✅ 4. **Rune Type (`rune`)**

* Alias for `int32`
* Represents a **Unicode code point**
* Useful for **multibyte characters**

```go
var ch rune = 'あ' // Japanese character
fmt.Println(ch)     // Unicode point
```

> Runes allow working with characters, not just bytes.

---

## ✅ 5. **Byte Type (`byte`)**

* Alias for `uint8`
* Used for raw binary data or ASCII characters

```go
var b byte = 'A'
```

> Often used in file I/O, buffers, and byte slices (`[]byte`).

---

## 🧠 Type Inference

Go uses type inference with `:=`:

```go
age := 30           // inferred as int
pi := 3.14          // inferred as float64
flag := true        // inferred as bool
```

---

## 🗂 Default Zero Values

| Type      | Default Value |
| --------- | ------------- |
| `bool`    | `false`       |
| `int`     | `0`           |
| `float64` | `0.0`         |
| `string`  | `""` (empty)  |
| `pointer` | `nil`         |

---

## ⚙️ Best Practices

* Use `int` unless a specific size is required (e.g., for network data).
* Use `float64` for decimals unless performance/memory optimization is critical.
* Use `string` for text, `[]byte` for binary data.
* Use type aliases like `byte` and `rune` when appropriate for clarity.
