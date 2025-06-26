
---

## 🧩 What is an Array in Go?

An **array** in Go is a **fixed-length sequence** of elements **of the same type**. Once declared, its **size cannot be changed**.

### 🧠 Key Points:

* All elements are of the same type.
* Array length is part of its **type** (`[5]int` is different from `[10]int`).
* Arrays are **value types** — assigning an array copies all its elements.

---

## 🔤 Syntax of Declaring Arrays

```go
var arr [5]int // Array of 5 integers, default values (0)
```

```go
arr := [3]string{"Go", "Rust", "Python"} // Short declaration with initialization
```

```go
arr := [...]float64{3.14, 2.71, 1.41} // Let compiler infer length
```

---

## 🧪 Accessing Elements in an Array

```go
fmt.Println(arr[0]) // Access first element
arr[1] = 42         // Set second element to 42
```

---

## 📋 Looping Through an Array

### 🔁 Using for loop with index:

```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}
```

### 🔁 Using `range`:

```go
for index, value := range arr {
    fmt.Printf("Index %d: %v\n", index, value)
}
```

---

## 📦 Example Program

```go
package main

import "fmt"

func main() {
    var numbers = [5]int{10, 20, 30, 40, 50}

    fmt.Println("All elements:")
    for i := 0; i < len(numbers); i++ {
        fmt.Printf("Element at index %d is %d\n", i, numbers[i])
    }

    // Using range
    fmt.Println("\nUsing range:")
    for index, value := range numbers {
        fmt.Printf("Index %d, Value %d\n", index, value)
    }
}
```

---

## 🧠 Array is a Value Type

When we assign one array to another, **a copy is made**, not a reference.

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 100

fmt.Println(a) // Output: [1 2 3]
fmt.Println(b) // Output: [100 2 3]
```

This is **different from slices**, which are reference types.

---

## 🧱 Multidimensional Arrays

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

fmt.Println(matrix[1][2]) // Outputs 6
```

---

## 📌 Summary of Array Features

| Feature              | Description                           |
| -------------------- | ------------------------------------- |
| Fixed size           | Cannot grow or shrink dynamically     |
| Zero-based indexing  | Index starts from 0                   |
| Value type           | Assignment copies the entire array    |
| Homogeneous elements | All elements must be of the same type |
| Range-based loops    | Simplifies iteration                  |

---
