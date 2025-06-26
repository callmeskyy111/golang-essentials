In Go (Golang), **conditionals** are used to control the flow of a program based on whether a condition is true or false. Let's go over each type in detail with syntax, explanation, and examples.

---

## ✅ 1. `if` Statement

### 🔹 Syntax:

```go
if condition {
    // code block to execute if condition is true
}
```

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    age := 20
    if age >= 18 {
        fmt.Println("You are an adult.")
    }
}
```

---

## ✅ 2. `if-else` Statement

### 🔹 Syntax:

```go
if condition {
    // executes if condition is true
} else {
    // executes if condition is false
}
```

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    number := 5
    if number%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }
}
```

---

## ✅ 3. `if-else if-else` Ladder

### 🔹 Syntax:

```go
if condition1 {
    // executes if condition1 is true
} else if condition2 {
    // executes if condition2 is true
} else {
    // executes if all conditions are false
}
```

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    score := 75

    if score >= 90 {
        fmt.Println("Grade: A")
    } else if score >= 70 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C")
    }
}
```

---

## ✅ 4. Short Statement in `if`

Go allows declaring variables inside the `if` condition itself.

### 🔹 Syntax:

```go
if variable := expression; condition {
    // code block
}
```

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    if num := 9; num < 10 {
        fmt.Println("Single-digit number")
    }
}
```

---

## ✅ 5. Logical Operators in Conditions

You can use `&&` (AND), `||` (OR), and `!` (NOT) inside conditionals.

### 🔹 Example:

```go
package main

import "fmt"

func main() {
    age := 25
    hasLicense := true

    if age >= 18 && hasLicense {
        fmt.Println("You can drive.")
    } else {
        fmt.Println("You can't drive.")
    }
}
```

---

## ✅ Best Practices

* Always use curly braces `{}` even for one-line `if` blocks (Go enforces this).
* Declare short variables inside conditionals when they're only needed inside that block.
* Conditions must return a boolean value (no implicit type conversions like in some other languages).

---