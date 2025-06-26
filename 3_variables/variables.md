## 🔹 What is a Variable in Go? 📦

A **variable** is a named storage that holds a value. Every variable in Go has:

* A **name**
* A **type**
* And optionally, a **value**

Go is a **statically typed** language, which means the type of a variable is known at compile-time.

---

## 🔹 Declaring Variables

There are **3 primary ways** to declare variables in Go:

---

### **1. Using `var` keyword (explicit type)**

```go
var name string
name = "Skyy"
```

* Here, `name` is a variable of type `string`.
* You can assign a value later.

---

### **2. Using `var` with initialization (type inferred)**

```go
var age = 29
```

* The type (`int`) is inferred from the assigned value.

---

### **3. Short declaration with `:=` (only inside functions)**

```go
city := "Kolkata"
```

* This is the most commonly used method inside functions.
* Go infers the type automatically.
* **Cannot** be used outside functions (like at the package level).

---

## 🔹 Multiple Variable Declarations

```go
var x, y, z int = 1, 2, 3
```

or

```go
name, age := "Skyy", 29
```

---

## 🔹 Default Zero Values

If no value is assigned, Go gives a **default zero value** based on the type:

| Type                                   | Zero Value |
| -------------------------------------- | ---------- |
| `int`                                  | `0`        |
| `float64`                              | `0.0`      |
| `string`                               | `""`       |
| `bool`                                 | `false`    |
| pointers, slices, maps, channels, etc. | `nil`      |

Example:

```go
var salary int
fmt.Println(salary) // Output: 0
```

---

## 🔹 Constants vs Variables

Variables can change value, constants **cannot**:

```go
const pi = 3.14 // immutable
var radius = 5  // mutable
```

---

## 🔹 Scope of Variables

1. **Package-level** (declared outside all functions)
2. **Function-level / Local scope**
3. **Block-level** (inside `{}` blocks like loops, if statements)

---

## 🔹 Best Practices

* Use short variable names for small scopes, meaningful names for larger scopes.
* Prefer `:=` inside functions for cleaner code.
* Group related variables:

```go
var (
  name string
  age  int
  city string
)
```

---

## 🔹 Example Program

```go
package main

import "fmt"

func main() {
    var name string = "Skyy"
    age := 29
    var isDeveloper = true

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Developer:", isDeveloper)
}
```
---