In Go (Golang), `switch-case` is a **control flow statement** used to compare a value against multiple conditions, just like in many other programming languages. But Go’s `switch` is **more powerful and flexible** than many other languages like C, Java, or JavaScript.

---

### 🧠 Basic Syntax:

```go
switch variable {
case value1:
    // code block
case value2:
    // code block
default:
    // default block (optional)
}
```

---

### ✅ Example:

```go
package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend")
	}
}
```

🟢 **Output:**

```
Wednesday
```

---

### 🧩 Key Features of Switch in Go:

#### 1. **No `break` needed**

* Go automatically breaks after the first matching case.
* You **don’t need to write `break`** manually.

#### 2. **Multiple values in a case**

```go
switch day {
case 1, 2, 3, 4, 5:
	fmt.Println("Weekday")
case 6, 7:
	fmt.Println("Weekend")
}
```

#### 3. **Switch without an expression**

You can omit the value after `switch` and write conditions directly in `case`.

```go
x := 42

switch {
case x < 0:
	fmt.Println("Negative")
case x == 0:
	fmt.Println("Zero")
case x > 0:
	fmt.Println("Positive")
}
```

#### 4. **Fallthrough (optional keyword)**

If you want to **continue to the next case**, even if it matches, use `fallthrough`.

```go
switch 1 {
case 1:
	fmt.Println("One")
	fallthrough
case 2:
	fmt.Println("Two")
}
```

🟢 **Output:**

```
One
Two
```

---

### 📌 Summary:

| Feature             | Behavior in GoLang                |
| ------------------- | --------------------------------- |
| `break`             | Not required                      |
| `fallthrough`       | Optional, must be used explicitly |
| `default` case      | Optional fallback                 |
| Multiple values     | Supported                         |
| Condition-only mode | Supported (no expression)         |

---

A **real-world calculator example using `switch-case`** in Go. This program takes two numbers and an operator from the user, performs the operation, and prints the result.

---

### 🧮 **Calculator Example in Go using switch-case**

```go
package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	// Input from the user
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	// Switch-case to determine the operation
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		} else {
			fmt.Println("Error: Division by zero")
		}
	default:
		fmt.Println("Invalid operator")
	}
}
```

---

### 🟢 Sample Run:

```
Enter first number: 20
Enter an operator (+, -, *, /): /
Enter second number: 5
20.00 / 5.00 = 4.00
```

---

### 🧠 Breakdown:

| Code Part                 | Purpose                              |
| ------------------------- | ------------------------------------ |
| `fmt.Scanln(&num1)`       | Reads the first number from the user |
| `switch operator { ... }` | Controls which operation to execute  |
| `case "+": ...`           | Executes addition                    |
| `default:`                | Catches unsupported operators        |
| `if num2 != 0`            | Prevents division by zero            |

---


