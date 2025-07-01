In Go, there is **no built-in `enum` keyword** like in other languages (such as `enum` in TypeScript, C, Java, etc.). However, we **simulate enums** using **constant blocks** and **typed constants**. Despite this workaround, Go provides powerful mechanisms (like `iota`, `stringer`, etc.) to implement **type-safe enumerations** effectively.

Let’s break it down **step-by-step** and go deep.

---

## 🔢 What Are Enums?

> Enums (enumerations) are a way to give **names** to a set of **related constant values**, often to improve code clarity and restrict input to a valid set of choices.

---

## ✅ Simulating Enums in Go

### 💡 Using `const` and `iota`

```go
type Status int

const (
    Pending Status = iota  // 0
    Approved               // 1
    Rejected               // 2
)
```

* `iota` is reset to 0 each time a `const` block begins.
* Each subsequent constant increments `iota` automatically.

📌 Behind the scenes:

```go
const (
    Pending Status = 0
    Approved Status = 1
    Rejected Status = 2
)
```

---

## 🔍 Why Use a Named Type (`Status`)?

If we just used `int`, **any integer could be passed**, which is dangerous.

```go
func setStatus(s Status) { ... } // ✅ Type-safe

func setStatus(s int) { ... }    // ❌ Not restricted to allowed constants
```

---

## 🧠 Using `iota` with Expressions

```go
type Power int

const (
    Off Power = 1 << iota // 1
    Low                   // 2
    Medium                // 4
    High                  // 8
)
```

* This simulates **bitflags** using iota and bit shifting.

---

## 💬 Converting Enums to String (Manually)

By default, Go doesn’t convert enums to strings like `"Pending"` — instead, it just prints the numeric value.

To make it readable:

```go
func (s Status) String() string {
    return [...]string{"Pending", "Approved", "Rejected"}[s]
}
```

Usage:

```go
fmt.Println(Pending) // Output: Pending
```

---

## ⚙️ Automating String Methods with `stringer`

Go provides a tool called `stringer`:

```bash
go install golang.org/x/tools/cmd/stringer@latest
```

Add this comment above your enum block:

```go
//go:generate stringer -type=Status
type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)
```

Then run:

```bash
go generate
```

This creates a file like `status_string.go` with a `String()` method implemented.

---

## 🧪 Enum Example: User Roles

```go
type Role int

const (
    Admin Role = iota
    User
    Guest
)

func (r Role) String() string {
    return [...]string{"Admin", "User", "Guest"}[r]
}

func isAuthorized(r Role) bool {
    return r == Admin || r == User
}
```

---

## 🧱 Enum Validations

You can restrict inputs:

```go
func IsValidStatus(s Status) bool {
    return s >= Pending && s <= Rejected
}
```

Or use a map:

```go
var validRoles = map[Role]bool{
    Admin: true,
    User: true,
    Guest: true,
}
```

---

## 🔐 Enums in JSON / APIs

By default, Go marshals enums as numbers in JSON. To use strings:

### 🧩 Custom MarshalJSON:

```go
func (s Status) MarshalJSON() ([]byte, error) {
    return json.Marshal(s.String())
}
```

### 📦 Or use `stringer` and tags:

```go
type Status int

const (
    Active Status = iota
    Inactive
)

func (s Status) String() string {
    return [...]string{"Active", "Inactive"}[s]
}
```

---

## 🧾 Summary

| Feature                    | Go’s Way                                  |
| -------------------------- | ----------------------------------------- |
| Native `enum` support      | ❌ Not built-in                            |
| Simulated enum             | ✅ `const` block with `iota`               |
| Type-safe enums            | ✅ Define custom type like `type Role int` |
| Enum string representation | ✅ Manually or via `stringer` tool         |
| Bitflag enums              | ✅ With `1 << iota` pattern                |
| JSON encoding              | ✅ Custom `MarshalJSON()`                  |

---

## 🧠 Best Practices

1. **Always define a custom type** (e.g., `type Status int`) for type safety.
2. Use `iota` to auto-increment values.
3. Implement the `String()` method for readability/logging.
4. Use the `stringer` tool in production for cleaner code.
5. Validate enum values before use.

---
# CODE-EXPLANATION

Let’s walk through our code piece by piece. We're simulating **enums** in Go using **typed constants**, and it’s a great example showing both `int`-based and `string`-based enums in action — with a Pokémon theme! 🟡⚡

---

## 🔢 Section 1: Enum-Like Constants (Integer Type)

```go
type PikachuAttack int

const (
	ThunderBolt PikachuAttack = iota
	Thunder
	Agility
	QuickAttack
	DoubleTeam
)
```

### ✅ What's Happening:

* `type PikachuAttack int` defines a **new named type** based on `int`. This adds **type safety** — you can’t pass any plain `int` to functions expecting `PikachuAttack`.
* The constants are defined using `iota` which:

  * Starts from `0` and increments automatically.
  * So effectively you have:

```go
const (
	ThunderBolt  = 0
	Thunder      = 1
	Agility      = 2
	QuickAttack  = 3
	DoubleTeam   = 4
)
```

All of type `PikachuAttack`.

### 📦 Purpose:

This mimics an enum where each attack is represented by a number. Go prefers this over real enum constructs.

---

## 🔡 Section 2: Enum-Like Constants (String Type)

```go
type MyFavPokemons string

const (
	Noctowl MyFavPokemons = "Noctowl"
	Umbreon MyFavPokemons = "Umbreon"
	Rapidash MyFavPokemons = "Rapidash"
)
```

### ✅ What's Happening:

* `MyFavPokemons` is a new type based on `string`.
* You’re assigning literal strings as values.
* This approach is useful when:

  * You want **readable output**
  * Or need to **marshal**/serialize enums into JSON for APIs

So:

```go
Noctowl  = "Noctowl"
Umbreon  = "Umbreon"
Rapidash = "Rapidash"
```

All of type `MyFavPokemons`.

---

## ⚙️ Section 3: Functions

```go
func changePikachuAttack(attack PikachuAttack){
	fmt.Println("Pika Pika:", attack)
}
```

* Accepts a `PikachuAttack` type.
* Since the enum is int-based, printing `attack` will show the **underlying numeric value**, not the name like "Thunderbolt" (unless you implement `.String()` method).

---

```go
func whatsMyFavPokemon(pokemon MyFavPokemons){
	fmt.Println(pokemon, " .. I choose you!")
}
```

* Accepts a string-based enum.
* Prints readable strings like `"Umbreon"` since they are already string-typed.

---

## ▶️ main()

```go
func main(){
	changePikachuAttack(ThunderBolt) // Output: Pika Pika: 0
	changePikachuAttack(Agility)     // Output: Pika Pika: 2

	whatsMyFavPokemon(Noctowl)       // Output: Noctowl  .. I choose you!
	whatsMyFavPokemon(Umbreon)       // Output: Umbreon  .. I choose you!
}
```

### 📌 Output:

```txt
Pika Pika: 0
Pika Pika: 2
Noctowl  .. I choose you!
Umbreon  .. I choose you!
```

---

## 🧠 Improvements You Can Try

### 1. Add `String()` method to PikachuAttack for better readability:

```go
func (p PikachuAttack) String() string {
    return [...]string{"ThunderBolt", "Thunder", "Agility", "QuickAttack", "DoubleTeam"}[p]
}
```

Now:

```go
changePikachuAttack(Agility) // Output: Pika Pika: Agility
```

---

## ✅ Summary

| Concept                | Description                                      |
| ---------------------- | ------------------------------------------------ |
| `iota` enum simulation | Auto-incrementing `int` constants                |
| Type safety            | Created via `type PikachuAttack int`             |
| String enums           | Useful for readability and APIs                  |
| Enums in functions     | Restricted to valid defined constants            |
| Output behavior        | `int` enums show numbers unless `String()` added |

---

