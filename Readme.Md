## Additional, curated Reading-Material 📖👨🏻‍💻

## CODE-EXPLANATION 💻
Let's break down our Go code in `bank.go` in **clear sections**, so we understand **what each part does**, why it matters, and how Go concepts like **file handling**, **error handling**, **loops**, and **control structures** work together in this basic banking CLI app.

---

## ✅ OVERVIEW

This is a **command-line banking app** called **GoBank**, built using Go. It supports:

* Reading and writing the balance from a file (`balance.txt`)
* Checking balance
* Depositing money
* Withdrawing money
* Graceful error handling (e.g., missing file or bad data)

---

## 🔹 1. IMPORTS

```go
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
```

We import:

* `fmt`: For input/output formatting
* `os`: For file handling
* `strconv`: For converting strings to float
* `errors`: For custom error messages

---

## 🔹 2. GLOBAL CONSTANT

```go
const accountBalanceFile = "balance.txt"
```

Declares a **constant** file name to avoid hardcoding `"balance.txt"` everywhere. Centralized and clean.

---

## 🔹 3. Function: `readBalanceFromFile()`

### Purpose:

To **read and return the balance** stored in `balance.txt`. If it doesn’t exist or has bad data, return `1000.00` as default.

### Code:

```go
func readBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000.00, errors.New("file doesn't exist ⚠️")
	}
	balanceStr := string(data)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		return 1000.00, errors.New("failed to parse stored balance ⚠️")
	}
	return balance, nil
}
```

### Breakdown:

* Reads file content as a `[]byte`
* Converts to string
* Parses string to `float64`
* Handles:

  * Missing file (`os.ReadFile` fails)
  * Corrupted data (`strconv.ParseFloat` fails)
* Returns a **default balance (1000.00)** on error

---

## 🔹 4. Function: `main()`

### Part A: Error Handling on Startup

```go
var accBalance, err = readBalanceFromFile()
if err != nil {
	fmt.Println("ERROR:", err)
	fmt.Println("----------------------")
	panic("Exiting the process.. 🔴")
}
```

* Reads initial balance from file
* If error: logs and exits using `panic`

---

### Part B: Infinite Loop (Banking Menu)

```go
for {
    ...
}
```

The loop lets us:

* Repeatedly ask user for banking actions
* Exit when the user gives an invalid option (anything other than 1, 2, or 3)

---

## 🔹 5. Menu System

```go
fmt.Println("\nYour amount is: $", accBalance)
fmt.Println("What do you want to do?")
fmt.Println("1️⃣. Check balance")
fmt.Println("2️⃣. Deposit")
fmt.Println("3️⃣. Withdraw")
fmt.Println("🔘.OTHER - Exit")
```

Prompts the user with 3 options (switch-case handles input next).

---

## 🔹 6. `switch choice` Block

### Case 1: Check Balance

```go
case 1:
	fmt.Println("Your balance is: $", accBalance)
```

Just prints current balance.

---

### Case 2: Deposit Money

```go
case 2:
	fmt.Print("💰 How much do you wanna deposit?: +$")
	var depositAmt float64
	fmt.Scan(&depositAmt)

	if depositAmt <= 0 {
		fmt.Println("INVALID AMOUNT!.. AMOUNT must not be -ve")
		continue
	}
	accBalance += depositAmt
	formattedBalance := fmt.Sprintf("%.2f", accBalance)
	fmt.Println("Deposited ✅.. Your updated account-balance: $", formattedBalance)

	os.WriteFile("balance.txt", []byte(formattedBalance), 0644)
```

* Reads deposit amount
* Rejects 0 or negative input
* Adds to balance
* Updates `balance.txt` using `os.WriteFile` (overwrites with new value)
* Uses `0644` (read/write owner, read others) for file permissions

---

### Case 3: Withdraw Money

```go
case 3:
	fmt.Print("💰 How much do you wanna withdraw?: -$")
	var withdrawAmt float64
	fmt.Scan(&withdrawAmt)

	if withdrawAmt <= 0 {
		fmt.Println("INVALID AMOUNT!.. AMOUNT must not be -ve")
		continue
	}
	if withdrawAmt > accBalance {
		fmt.Println("Insufficient Balance :(")
		continue
	}
	accBalance -= withdrawAmt
	formattedBalance := fmt.Sprintf("%.2f", accBalance)
	fmt.Println("Amount withdrawn ✅.. Your updated account-balance: $", formattedBalance)
```

* Reads withdrawal amount
* Rejects 0 or negative input
* Checks for sufficient balance
* Updates balance

⚠️ **Missing file update!** → You forgot to save new balance to `balance.txt` here. You should add:

```go
os.WriteFile("balance.txt", []byte(formattedBalance), 0644)
```

---

### Default Case: Exit

```go
default:
	fmt.Println("Exiting.. Thanks for choosing GoBank")
	return
```

Any other input → exits app cleanly using `return`.

---

## 🔚 Summary of Key Concepts Used

| Concept                        | Used Where                                          |
| ------------------------------ | --------------------------------------------------- |
| **File I/O**                   | `os.ReadFile`, `os.WriteFile`                       |
| **Error handling**             | `if err != nil`, custom `errors.New()`              |
| **Control structure**          | `switch`, `for`, `continue`, `return`               |
| **CLI input**                  | `fmt.Scan()`                                        |
| **String to float conversion** | `strconv.ParseFloat()`                              |
| **Formatting floats**          | `fmt.Sprintf("%.2f", value)`                        |
| **Infinite loop**              | `for { ... }`                                       |
| **Panic**                      | Exits app with a red flag when something goes wrong |

---

## ✅ Improvements You Can Make

* ✅ Save balance after withdrawal (missing now)
* ✅ Validate string input (e.g., avoid crashing if user types "abc" instead of a number)
* ✅ Use functions to modularize logic: `deposit()`, `withdraw()`, `writeToFile()`
* ✅ Graceful `os.Exit(1)` instead of `panic`

---

## Loops in GoLang ➰
Let’s break down **loops in Go** in a clear and beginner-friendly way — including **syntax**, **types**, **examples**, and **key differences** from other languages like JavaScript or C.

---

## 🔁 1. What Are Loops?

In Go, **loops allow us to repeat a block of code** multiple times. Go has only **one loop construct**: `for`.

There is **no `while` or `do-while`** like in other languages — everything is done using `for`.

---

## 🟩 2. Basic `for` Loop

### Syntax:

```go
for initializer; condition; post {
    // loop body
}
```

### Example:

```go
for i := 0; i < 5; i++ {
    fmt.Println("i is", i)
}
```

* `i := 0` → Initialization
* `i < 5` → Condition
* `i++` → Post-expression (increment)

This prints:

```
i is 0
i is 1
i is 2
i is 3
i is 4
```

---

## 🔄 3. `for` as a While Loop

Go has **no `while` keyword**, but we can mimic it like this:

```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

This behaves exactly like a while loop in other languages.

---

## ♾️ 4. Infinite Loop

We can create an infinite loop using just:

```go
for {
    fmt.Println("Looping forever...")
}
```

You’ll usually combine this with `break` or `return` inside the loop to exit at some point.

---

## 🛑 5. `break` and `continue`

* `break` → Immediately exits the loop.
* `continue` → Skips current iteration and moves to the next.

### Example:

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // skip 3
    }
    if i == 5 {
        break // stop loop when i == 5
    }
    fmt.Println(i)
}
```

Output:

```
1
2
4
```

---

## 🔂 6. Looping Over Collections (Array/Slice/Map)

### Over a slice:

```go
nums := []int{10, 20, 30}
for index, value := range nums {
    fmt.Println("Index:", index, "Value:", value)
}
```

### Over a map:

```go
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Println("Key:", key, "Value:", value)
}
```

You can also ignore unused variables with `_`:

```go
for _, v := range nums {
    fmt.Println(v)
}
```

---

## 📌 Recap Table

| Loop Type         | Syntax Example                        |
| ----------------- | ------------------------------------- |
| Traditional `for` | `for i := 0; i < 10; i++ { ... }`     |
| `while` style     | `for condition { ... }`               |
| Infinite loop     | `for { ... }`                         |
| With `range`      | `for i, v := range slice/map { ... }` |

---

## 🚀 Common Use Cases

* Reading user input until they quit → `for { ... }` with `break`
* Iterating over list of items → `for _, item := range ...`
* Implementing retry logic → `for i := 0; i < 3; i++ {}`

---

## Error-Handling in GoLang ⚠️
Let’s dive deep into **error-handling in GoLang** — a very important and idiomatic part of Go programming.

---

## 🧱 **1. Go’s Philosophy on Errors**

Go does **not** use `try/catch/finally` like JavaScript, Python, or Java.

Instead, **errors are values**, returned as the **last return value** from functions, allowing **explicit and predictable error handling**.

> ✅ Idiomatic Go: *"Don't panic. Handle errors explicitly."*

---

## 🧾 **2. The `error` Type**

The `error` type is a built-in **interface** in Go:

```go
type error interface {
    Error() string
}
```

Any type that implements the `Error()` method (returning a string) satisfies the `error` interface.

---

## ⚙️ **3. Returning and Handling Errors**

### A. Returning an error from a function:

```go
import "errors"

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

### B. Handling the error:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

> ✅ Go programmers check `if err != nil` **all the time**.

---

## 🔧 **4. Creating Custom Errors**

### A. With `errors.New`:

```go
import "errors"
err := errors.New("something went wrong")
```

### B. With `fmt.Errorf` (for dynamic messages):

```go
name := "Skyy"
err := fmt.Errorf("user %s not found", name)
```

---

## 🧰 **5. Custom Error Types (Advanced)**

You can define a struct that implements the `Error()` method.

```go
type DivideError struct {
    Numerator   float64
    Denominator float64
}

func (e *DivideError) Error() string {
    return fmt.Sprintf("cannot divide %.2f by %.2f", e.Numerator, e.Denominator)
}

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, &DivideError{a, b}
    }
    return a / b, nil
}
```

Now you can even **type-assert** or inspect error details!

---

## 🧪 **6. Sentinel Errors (Comparing Errors)**

You might define reusable global error values:

```go
var ErrNotFound = errors.New("item not found")

func findItem(id int) error {
    return ErrNotFound
}

err := findItem(10)
if errors.Is(err, ErrNotFound) {
    fmt.Println("Handle item-not-found specifically")
}
```

Since Go 1.13+, use `errors.Is` and `errors.As`.

---

## 🪜 **7. `errors.Is` and `errors.As`**

### A. `errors.Is` → Check if an error **matches** a known error

```go
if errors.Is(err, ErrNotFound) { ... }
```

### B. `errors.As` → Check if error is of a specific **type**

```go
var myErr *DivideError
if errors.As(err, &myErr) {
    fmt.Println("DivideError with values:", myErr.Numerator, myErr.Denominator)
}
```

---

## 💥 **8. `panic` and `recover` (Rare Use)**

Go also has `panic()` and `recover()` for unrecoverable issues — like nil dereference or failed assertions.

### A. `panic()` → Immediately stops execution and unwinds the stack

```go
panic("something went terribly wrong")
```

### B. `recover()` → Used to catch a panic **only inside a `defer`**

```go
func mayPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("Boom 💥")
}
```

✅ Only use `panic` for unrecoverable errors (e.g., corrupted memory or developer mistakes).

---

## 🧠 Best Practices in Go Error Handling

| Practice            | Description                           |
| ------------------- | ------------------------------------- |
| `if err != nil`     | Always check errors right away        |
| `return nil`        | Return nil if no error occurred       |
| `fmt.Errorf`        | For contextual messages               |
| `errors.Is/As`      | For error comparison and inspection   |
| Don’t use panic     | Except for truly unexpected errors    |
| Don’t ignore errors | Compiler may allow it, but it’s risky |

---

## ✅ Example Recap

```go
func readFile(name string) ([]byte, error) {
    data, err := os.ReadFile(name)
    if err != nil {
        return nil, fmt.Errorf("failed to read %s: %w", name, err)
    }
    return data, nil
}

func main() {
    content, err := readFile("file.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("File content:", string(content))
}
```

---

---

## 🔍 **Difference Between `panic` and `error` in Go**

| Feature                  | `error`                                                   | `panic`                                                            |
| ------------------------ | --------------------------------------------------------- | ------------------------------------------------------------------ |
| 🧠 Concept               | Expected issues                                           | Unexpected, critical failures                                      |
| 🛠️ Used for             | Recoverable problems (e.g. file not found, invalid input) | Unrecoverable problems (e.g. array out-of-bounds, nil dereference) |
| 🔁 How it works          | Returned as a value                                       | Stops normal flow, starts unwinding                                |
| 🎯 Goal                  | Let the caller decide what to do                          | Exit the program or recover via `recover()`                        |
| 🧪 Should be tested      | Yes — check `err != nil`                                  | Rarely — usually you avoid triggering `panic` in tests             |
| 🧘 Recoverable?          | Yes                                                       | Only inside a `defer` using `recover()`                            |
| 🔥 Used like exceptions? | No                                                        | Somewhat, but only for **severe** issues                           |

---

## 🧵 `error`: Normal Flow of Handling Issues

* **Idiomatic Go** always prefers **returning errors** instead of throwing them.
* You get the error as a value and decide what to do with it.

### ✅ Example: Using `error`

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Handled error:", err)
		return
	}
	fmt.Println("Result:", result)
}
```

---

## 💥 `panic`: For Critical Failures

* A `panic` **immediately stops execution** and **starts unwinding the call stack**.
* You can **recover** using `recover()` inside a `defer` block, but this is rare and mainly used in **library code**, **servers**, or **frameworks**.

### ❌ Example: Using `panic` (Not idiomatic in most cases)

```go
func divide(a, b float64) float64 {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func main() {
	fmt.Println("Result:", divide(10, 0)) // 🔴 program will crash here
}
```

Output:

```
panic: division by zero
```

---

## 🧘‍♂️ `recover`: Catching Panics

You can **recover** from a panic to keep your app alive:

```go
func safeDivide(a, b float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	fmt.Println("Result:", a/b)
}

func main() {
	safeDivide(10, 0)
	fmt.Println("App is still running...")
}
```

---

## 🧠 When to Use What?

| Use Case                                      | Use `error` ✅ | Use `panic` ⚠️ |
| --------------------------------------------- | ------------- | -------------- |
| File not found                                | ✅             | ❌              |
| User input invalid                            | ✅             | ❌              |
| Failed to connect to DB (known issue)         | ✅             | ❌              |
| Nil pointer dereference (bug in code)         | ❌             | ✅              |
| Array out-of-bound (programmer mistake)       | ❌             | ✅              |
| Unexpected internal corruption or assert fail | ❌             | ✅              |
| Library should force crash on misuse          | ❌             | ✅              |
| Critical init failure (e.g. missing ENV)      | ❌             | ✅ (sometimes)  |

---

## 🔐 TL;DR

* ✅ Use `error` for **normal, expected** failures.
* ⚠️ Use `panic` for **unexpected, unrecoverable** issues (like bugs).
* 🛑 Never use `panic` to control logic or skip normal error handling.
* 🧘 Only use `recover()` in very controlled situations — e.g., web servers to prevent crashing the entire app.

---

---

## 🔐 Linux File Permissions: Overview

In Linux, every file and directory has **permissions** associated with it that define **who can read, write, or execute** them. These permissions are set for three kinds of users:

### 👥 User Types

| Symbol | Description                                 |
| ------ | ------------------------------------------- |
| `u`    | User (Owner) — the person who owns the file |
| `g`    | Group — a set of users in the same group    |
| `o`    | Others — everyone else on the system        |
| `a`    | All (user + group + others)                 |

---

## 🔤 Permission Types

Each user type can have three kinds of permissions:

| Symbol | Name    | Meaning on File            | Meaning on Directory        |
| ------ | ------- | -------------------------- | --------------------------- |
| `r`    | Read    | Can view file contents     | Can list directory contents |
| `w`    | Write   | Can modify file contents   | Can add/remove files in dir |
| `x`    | Execute | Can run the file as script | Can `cd` into the directory |

---

## 🧾 Viewing File Permissions (`ls -l`)

```bash
ls -l myfile.txt
```

Output example:

```
-rwxr-xr--
```

Breakdown:

* `-` : Type (dash = regular file, `d` = directory)
* `rwx` : Owner has **read, write, execute**
* `r-x` : Group has **read, execute**
* `r--` : Others have **read only**

---

## 🔢 Numeric (Octal) Representation

Each permission is represented by a number:

* `r = 4`
* `w = 2`
* `x = 1`

Sum them up for each user group:

| Permissions | Value |
| ----------- | ----- |
| `rwx`       | 7     |
| `rw-`       | 6     |
| `r--`       | 4     |
| `---`       | 0     |

So:

```bash
chmod 755 myfile.sh
```

Means:

* Owner: `7` → `rwx`
* Group: `5` → `r-x`
* Others: `5` → `r-x`

---

## 🛠️ Changing Permissions

### `chmod` – Change mode

```bash
chmod 644 file.txt   # rw-r--r--
chmod +x script.sh   # adds execute permission
chmod u+x script.sh  # adds execute only for owner
```

### `chown` – Change ownership

```bash
chown skyy:developers file.txt
# User = skyy, Group = developers
```

---

## 🏷 Special Permissions (Advanced)

| Symbol | Name          | Description                                      |
| ------ | ------------- | ------------------------------------------------ |
| `s`    | setuid/setgid | Executes with owner/group privileges             |
| `t`    | sticky bit    | Only owner can delete files (commonly on `/tmp`) |

Example: `chmod 4755 file` sets **setuid** (4 before permissions).

---

## 🔐 Typical Use-Cases

| Use Case                    | Command                 |
| --------------------------- | ----------------------- |
| Make a script executable    | `chmod +x script.sh`    |
| Give read/write to everyone | `chmod 666 file.txt`    |
| Secure a private file       | `chmod 600 secrets.txt` |
| Make directory accessible   | `chmod 755 /myfolder`   |
| Make file owner-only        | `chmod 700 private.sh`  |

---

## 🧪 Test It Out!

Create a file:

```bash
touch test.txt
ls -l test.txt
chmod 600 test.txt
ls -l test.txt
```

---

Now.. Let's walk through **Linux file permissions in the context of deploying a Go (Golang) app** on a Linux server — this is *super important* for security and functionality.

---

## 🧑‍💻 Scenario: Deploying a Go Web App on a Linux Server

Let's say we have a Go app with the following structure:

```
/home/skyy/my-go-app/
├── main.go
├── app        # compiled binary
├── config.env
└── static/    # static files (HTML, CSS, JS)
```

We want:

* Only **Skyy** (the owner) to run and modify the app.
* The server (e.g., Nginx or systemd) to **execute** the Go binary.
* The public to access only what's needed (like `static/` files if served directly).
* **No one else** (other users) to view configs or tamper with code.

---

## ✅ Step-by-Step File Permission Setup

### 1. **Binary (`app`)**

```bash
chmod 755 app
```

Means:

* `7` for owner (read/write/execute)
* `5` for group (read/execute)
* `5` for others (read/execute)

> ✅ Allows executing the app, but **only the owner can modify** it.

### 2. **Main Go Source File (`main.go`)**

```bash
chmod 644 main.go
```

* Owner: read/write
* Group/Others: read

> ✅ Useful during dev; in production you might remove it or set it to `600`.

### 3. **Environment Config (`config.env`)**

```bash
chmod 600 config.env
```

> 🔒 **Very important**: Only the owner can read/write it. Prevents exposing secrets.

### 4. **Static Assets (`static/`)**

```bash
chmod -R 755 static/
```

* `-R`: recursively sets permissions

> ✅ Allows public read access for web server to serve files, but no write access.

---

## 🔐 Ownership Control (Optional)

If your Go app runs as a different Linux user (e.g., via a systemd service), you might:

```bash
chown root:www-data app
chmod 750 app
```

> This allows the app to be run by `www-data` (like Nginx), but **not by everyone**.

---

## 🔄 Sample: Full Command Setup

```bash
cd /home/skyy/my-go-app

chmod 755 app             # executable binary
chmod 600 config.env      # protect secrets
chmod -R 755 static/      # public static files
chmod 644 main.go         # source file (for development only)
```

---

## 🧪 Quick Test

To confirm permission levels:

```bash
ls -l
# OR to view recursively
ls -lR
```

---

## ✅ Pro Tips

| Tip                                   | Why                                                                        |
| ------------------------------------- | -------------------------------------------------------------------------- |
| Use `.env` file permissions carefully | They often store API keys or DB credentials                                |
| Don’t make everything 777             | That’s a security disaster — gives **everyone full control**               |
| Use `sudo` only when needed           | You don’t want to accidentally give root-level access to wrong files       |
| Make logs readonly                    | e.g., `chmod 644 logs.txt` to allow viewing but not editing logs by others |

---

