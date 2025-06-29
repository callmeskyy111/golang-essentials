Let's explore **structs in Go** (Golang) in **deep detail**, since they are essential to building real-world applications in Go — from APIs and systems to backend services.

---

## 🧱 What is a Struct in Go?

A `struct` (short for **structure**) is a **composite type** that groups together variables (called **fields**) under a single name. It's used to model **real-world entities** and **custom data types**.

### 🔹 Think of a struct like an object (without methods) in OOP languages.

---

## ✅ Declaring a Struct

```go
type Person struct {
    Name string
    Age  int
}
```

This defines a `Person` type with two fields:

* `Name` of type `string`
* `Age` of type `int`

---

## 🛠 Creating Struct Instances

```go
p1 := Person{Name: "Skyy", Age: 29}
p2 := Person{"John", 30}
var p3 Person
p3.Name = "Alice"
p3.Age = 25
```

| Syntax        | Description                                 |
| ------------- | ------------------------------------------- |
| `Person{}`    | Struct literal                              |
| `new(Person)` | Returns a pointer to a zero-valued `Person` |
| `&Person{}`   | Also creates a pointer to struct            |

---

## 🧾 Accessing & Modifying Fields

```go
fmt.Println(p1.Name)   // "Skyy"
p1.Age = 30
fmt.Println(p1.Age)    // 30
```

---

## 🪄 Struct with Pointers

```go
p := &Person{Name: "Lily", Age: 22}
fmt.Println(p.Name) // "Lily"
p.Age = 23
```

You can use `.` to access fields even through pointers — Go dereferences automatically.

---

## 🔁 Structs Are Passed by Value (Not by Reference)

```go
func changeName(p Person) {
    p.Name = "Changed"
}

func main() {
    x := Person{Name: "Skyy", Age: 29}
    changeName(x)
    fmt.Println(x.Name) // Skyy (NOT changed)
}
```

🔁 To update it, use a pointer:

```go
func changeName(p *Person) {
    p.Name = "Changed"
}
```

---

## 🧩 Embedded Structs (Composition)

Go doesn't support inheritance but uses **embedding**.

```go
type Address struct {
    City  string
    State string
}

type Employee struct {
    Name string
    Address  // embedded
}
```

Now we can do:

```go
emp := Employee{Name: "Skyy", Address: Address{City: "Kolkata", State: "WB"}}
fmt.Println(emp.City)  // Direct access due to embedding
```

---

## 🧠 Anonymous Fields

Go allows anonymous fields — usually used for composition:

```go
type Job struct {
    string
    int
}

j := Job{"Developer", 1000}
fmt.Println(j.string) // Developer
```

But **not recommended** for clarity.

---

## 🎯 Struct Tags

Used for metadata, often in APIs (e.g., with `json`):

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

When serialized to JSON:

```go
{"name": "Skyy", "email": "skyy@example.com"}
```

---

## 🧪 Comparing Structs

Structs are **comparable** if all their fields are comparable.

```go
a := Person{Name: "Skyy", Age: 29}
b := Person{Name: "Skyy", Age: 29}
fmt.Println(a == b) // true
```

But slices, maps, or functions inside struct make it **non-comparable**.

---

## 📦 Zero Value of Struct

When a struct is declared without initialization, it gets **zero values**:

```go
var p Person
fmt.Println(p.Name) // ""
fmt.Println(p.Age)  // 0
```

---

## 🔃 Struct Methods

Methods can be attached to structs via **receiver functions**.

```go
func (p Person) Greet() {
    fmt.Printf("Hi, I'm %s!\n", p.Name)
}
```

### Pointer vs Value Receivers

```go
func (p *Person) HaveBirthday() {
    p.Age += 1
}
```

Use a **pointer receiver** when:

* You want to mutate the struct
* The struct is large (avoids copy)
* You want consistency across methods

---

## 🔍 Real-life Example

```go
type Product struct {
    ID    string
    Name  string
    Stock int
}

func (p *Product) IsInStock() bool {
    return p.Stock > 0
}

func (p *Product) ReduceStock(qty int) {
    if qty <= p.Stock {
        p.Stock -= qty
    }
}
```

---

## 🧭 Summary Table

| Feature                   | Example                            | Notes                          |
| ------------------------- | ---------------------------------- | ------------------------------ |
| Declare struct            | `type User struct { Name string }` | Named struct                   |
| Create instance           | `u := User{"Skyy"}`                | Struct literal                 |
| Access field              | `u.Name`                           | Dot notation                   |
| Struct pointer            | `p := &User{"Skyy"}`               | Auto dereferencing             |
| Embedded struct           | `type A struct { B }`              | Enables composition            |
| Struct tag                | `json:"name"`                      | Metadata for encoding/decoding |
| Method on struct          | `func (u *User) SayHi()`           | Method with receiver           |
| Value vs pointer receiver | `User` vs `*User`                  | Pointer needed for mutation    |

---

## ✅ Best Practices

* Use `structs` to model real-world entities (User, Product, Order, etc.).
* Use **pointers** when:

  * Modifying struct
  * Large data
* Use **struct tags** for clean serialization (JSON, XML).
* Prefer **composition (embedding)** over inheritance.
* Stick with **explicit field names** (avoid anonymous fields for readability).

---

Absolutely, Skyy! Let's dive deep into **struct embedding in Go**, one of the most important features for **composition over inheritance**.

---

## 🧱 What is Struct Embedding in Go?

**Struct embedding** is a way to achieve **code reuse** and **composition** in Go.

Instead of inheritance (like in OOP languages), Go **embeds** one struct into another, allowing the outer struct to **promote** the fields and methods of the embedded struct.

> 🧠 Think of it as a "has-a" relationship, not "is-a".

---

## 🧪 Basic Syntax

```go
type Address struct {
    City  string
    State string
}

type Person struct {
    Name string
    Age  int
    Address // Embedded (anonymous) field
}
```

✅ Here, `Person` embeds `Address`.

Now, you can access `Address` fields directly via `Person`:

```go
func main() {
    p := Person{
        Name: "Skyy",
        Age:  29,
        Address: Address{
            City:  "Kolkata",
            State: "West Bengal",
        },
    }

    fmt.Println(p.City)  // Access promoted field from Address
    fmt.Println(p.State) // Same
}
```

---

## 🔍 Behind the Scenes

Under the hood, Go composes the fields like this:

```go
p.Address.City
```

But allows you to write it as:

```go
p.City
```

This is called **field promotion**.

---

## 🧑‍🔧 Embedded Struct Methods

If the embedded struct has methods, they are **also promoted**.

```go
type Logger struct{}

func (l Logger) Log(msg string) {
    fmt.Println("LOG:", msg)
}

type Server struct {
    Host string
    Logger
}

func main() {
    s := Server{Host: "localhost"}
    s.Log("Server started") // Promoted method from Logger
}
```

We can still access it explicitly with `s.Logger.Log(...)`.

---

## 🧠 Why Use Struct Embedding?

| Goal                     | Explanation                                         |
| ------------------------ | --------------------------------------------------- |
| Composition              | Reuse logic and data without inheritance            |
| Clean syntax             | Avoid boilerplate delegation                        |
| Promote methods/fields   | Access embedded struct fields/methods directly      |
| Interface implementation | Embedded structs help satisfy interfaces implicitly |

---

## 🔁 Named vs Anonymous Embedding

### ✅ Anonymous Embedding

```go
type Engine struct {
    Power int
}

type Car struct {
    Engine
}
```

We can use `c.Power`.

### ❌ Named Field (No Promotion)

```go
type Car struct {
    Engine Engine
}
```

You must write `c.Engine.Power`.

---

## 🧱 Struct Embedding & Interfaces

Embedding allows structs to **implicitly implement interfaces**.

```go
type Notifier interface {
    Notify()
}

type BaseNotifier struct{}

func (b BaseNotifier) Notify() {
    fmt.Println("Notification sent!")
}

type User struct {
    Name string
    BaseNotifier
}

func main() {
    var n Notifier = User{Name: "Skyy"}
    n.Notify() // "Notification sent!"
}
```

---

## ⚠️ Shadowing with Embedded Fields

If both embedded and outer structs have the same field/method name, the outer struct takes precedence:

```go
type A struct {
    Name string
}

type B struct {
    Name string
    A
}

func main() {
    b := B{ Name: "Outer", A: A{Name: "Inner"} }
    fmt.Println(b.Name)   // Outer
    fmt.Println(b.A.Name) // Inner
}
```

---

## 💡 Tip: Embedding Multiple Structs

```go
type Timestamps struct {
    CreatedAt string
    UpdatedAt string
}

type SoftDelete struct {
    Deleted bool
}

type Product struct {
    Name string
    Timestamps
    SoftDelete
}
```

Now `Product` has all those fields *promoted*:

```go
p := Product{Name: "Book", CreatedAt: "Today"}
fmt.Println(p.CreatedAt)
fmt.Println(p.Deleted)
```

---

## 📌 Summary

| Concept                   | Description                                         |
| ------------------------- | --------------------------------------------------- |
| Anonymous embedding       | Promotes fields and methods to outer struct         |
| Composition > Inheritance | Preferred pattern in Go                             |
| Method promotion          | Allows outer struct to call inner struct methods    |
| Field name conflict       | Outer struct's field overrides                      |
| Interfaces                | Embedded structs help satisfy interfaces implicitly |

---
