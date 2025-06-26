Let's dive into **Constants in Go (Golang)** in detail. Constants are an important part of Go, especially when we want to define **fixed values** that should not change during execution.

---

## 🔹 What Is a Constant in Go?

A **constant** is a name given to a value that **cannot be changed** after it's defined.

* Declared using the `const` keyword.
* Must be assigned a **compile-time determinable value** (i.e., not from a function or user input).
* Cannot be reassigned later.

---

## 🔸 Syntax of Declaring Constants

### **1. Basic Declaration**

```go
const pi = 3.14
const appName string = "Prescripto"
```

* Go will **infer the type** if not explicitly stated.
* Once set, the value **cannot be changed**.

---

### **2. Multiple Constants at Once**

```go
const (
    appVersion = "1.0.0"
    maxUsers   = 100
    timeout    = 30
)
```

---

## 🔸 Typed vs Untyped Constants

### **Untyped Constant**

```go
const x = 5
```

* No type is assigned at declaration.
* Type is inferred **only when used**.

### **Typed Constant**

```go
const y int = 5
```

* You restrict how the value is used.
* For example, `y` is explicitly an `int`, so it cannot be used where a `float64` is expected.

---

## 🔸 Valid Constant Types

Only certain **basic types** can be used:

* `bool`
* `string`
* `numeric` types: `int`, `float64`, `complex`, etc.

✅ Allowed:

```go
const name = "Skyy"
const age = 29
const pi = 3.1415
const isAdmin = false
```

❌ Not Allowed:

```go
// const user = getUserInput() // ❌ Error: value must be known at compile time
```

---

## 🔸 Constant Expressions

Go allows **compile-time expressions**:

```go
const width, height = 100, 50
const area = width * height // 5000
```

---

## 🔸 `iota` – Auto-Incrementing Constants

`iota` is a special identifier used to create **incrementing constants**. It's reset to 0 for each `const` block.

### Example:

```go
const (
    first = iota  // 0
    second        // 1
    third         // 2
)
```

### Real-world example:

```go
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

You can also use expressions with `iota`:

```go
const (
    KB = 1 << (10 * iota) // 1 KB = 1024
    MB                   // 1 MB = 1024 * 1024
    GB                   // ...
)
```

---

## 🔸 Differences Between `const` and `var`

| Feature            | `const`               | `var`                |
| ------------------ | --------------------- | -------------------- |
| Mutability         | Immutable (read-only) | Mutable (changeable) |
| Value at           | Compile-time          | Runtime              |
| Type required?     | No (optional)         | Optional             |
| Allowed types      | Only basic types      | Any type             |
| Can use functions? | ❌ No                  | ✅ Yes                |

---

## 🔸 Example Program

```go
package main

import "fmt"

const (
    appName = "Prescripto"
    version = "1.0.0"
    isBeta  = true
)

const (
    Sunday = iota
    Monday
    Tuesday
)

func main() {
    fmt.Println("App:", appName, "Version:", version, "Beta:", isBeta)
    fmt.Println("Days - Sunday:", Sunday, "Monday:", Monday, "Tuesday:", Tuesday)
}
```

---

## ✅ Summary

| Concept            | Key Point                              |
| ------------------ | -------------------------------------- |
| `const` keyword    | Declares a read-only value             |
| Compile-time value | Must be known at compile-time          |
| Typed/Untyped      | Go supports both                       |
| Expressions        | Allowed (e.g., `width * height`)       |
| `iota`             | Auto-increments values in const blocks |

---