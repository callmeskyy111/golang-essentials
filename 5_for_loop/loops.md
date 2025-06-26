In **Go (Golang)**, **loops** are used to execute a block of code repeatedly. Go keeps loops **simple** and provides **only one loop keyword**: `for`. But this single `for` is very versatile and covers **all looping needs** (like `while`, `do-while`, and `for` in other languages).

---

## 🔁 Types of Loops in Go

### ✅ 1. **Basic `for` Loop (Like C-style)**

This loop has **three parts**:

* **initialization**: executed once
* **condition**: checked before each iteration
* **post**: executed after each iteration

### 🔹 Syntax:

```go
for init; condition; post {
    // code block
}
```

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println("i =", i)
    }
}
```

---

### ✅ 2. **`for` as a `while` Loop**

If we remove the initialization and post, it behaves like a `while` loop.

### 🔹 Syntax:

```go
for condition {
    // code block
}
```

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    i := 1
    for i <= 5 {
        fmt.Println("i =", i)
        i++
    }
}
```

---

### ✅ 3. **Infinite Loop**

When we write `for` with **no condition**, it becomes an **infinite loop** (like `while(true)`).

### 🔹 Syntax:

```go
for {
    // code that runs forever unless we break
}
```

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    count := 0
    for {
        fmt.Println("Running...")
        count++
        if count == 3 {
            break
        }
    }
}
```

---

### ✅ 4. **Using `break` and `continue`**

* `break`: exits the loop immediately.
* `continue`: skips the rest of the code and moves to the next iteration.

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        if i == 3 {
            continue // skips 3
        }
        if i == 5 {
            break // stops at 4
        }
        fmt.Println("i =", i)
    }
}
```

---

### ✅ 5. **Looping Over Arrays, Slices, Maps (using `range`)**

`range` returns index/key and value when looping over collections.

### 🔹 Example with a Slice:

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30}
    for index, value := range nums {
        fmt.Println("Index:", index, "Value:", value)
    }
}
```

### 🔹 Example with a Map:

```go
package main

import "fmt"

func main() {
    colors := map[string]string{"A": "Red", "B": "Blue"}
    for key, value := range colors {
        fmt.Println(key, "=>", value)
    }
}
```

---

## 🧠 Summary Table

| Type              | Syntax Example              | Use Case                           |
| ----------------- | --------------------------- | ---------------------------------- |
| C-style `for`     | `for i := 0; i < 5; i++ {}` | Traditional counting loop          |
| `for` as `while`  | `for i < 10 {}`             | Condition-based loop               |
| Infinite loop     | `for {}`                    | Runs until manually broken         |
| Loop with `range` | `for k, v := range map {}`  | Iterating over slices/maps/strings |
| Loop control      | `break`, `continue`         | Control loop execution             |

---
