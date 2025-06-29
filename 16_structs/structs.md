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
