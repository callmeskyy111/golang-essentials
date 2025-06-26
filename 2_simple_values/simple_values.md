### All the basic data types in Go (Golang) 💻

## ✅ 1. **Boolean Type (`bool`)**

* Represents a truth value: `true` or `false`
* Default value: `false`

```go
var isActive bool = true
var isLoggedIn = false // inferred type
```

> Use in logical expressions, conditions, or control flow (e.g., `if`, `for`).

---

## ✅ 2. **Numeric Types**

### A. **Integer Types**

Go supports both **signed** and **unsigned** integers.

#### 🔹 Signed Integers:

| Type    | Size                                      | Range                           |
| ------- | ----------------------------------------- | ------------------------------- |
| `int8`  | 8-bit                                     | -128 to 127                     |
| `int16` | 16-bit                                    | -32,768 to 32,767               |
| `int32` | 32-bit                                    | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64-bit                                    | ±9.22×10^18                     |
| `int`   | 32 or 64-bit (depends on OS/architecture) |                                 |

#### 🔹 Unsigned Integers:

| Type     | Size                      | Range                     |
| -------- | ------------------------- | ------------------------- |
| `uint8`  | 8-bit                     | 0 to 255                  |
| `uint16` | 16-bit                    | 0 to 65,535               |
| `uint32` | 32-bit                    | 0 to 4,294,967,295        |
| `uint64` | 64-bit                    | 0 to 18,446,744,073×10^18 |
| `uint`   | 32 or 64-bit (like `int`) |                           |

#### Example:

```go
var a int8 = -100
var b uint16 = 65000
```

> ⚠️ Always choose the smallest type that fits your range to optimize memory.

---

### B. **Floating Point Types**

Used for **decimal or real numbers**.

| Type      | Size   | Precision           |
| --------- | ------ | ------------------- |
| `float32` | 32-bit | \~6 decimal places  |
| `float64` | 64-bit | \~15 decimal places |

#### Example:

```go
var pi float64 = 3.14159265358979
var price float32 = 19.99
```

> Use `float64` by default for better precision.

---

### C. **Complex Numbers**

Go supports complex numbers with real and imaginary parts.

| Type         | Description                      |
| ------------ | -------------------------------- |
| `complex64`  | float32 real and imaginary parts |
| `complex128` | float64 real and imaginary parts |

#### Example:

```go
var c1 complex64 = 1 + 2i
var c2 = complex(5.5, 3.3) // real: 5.5, imag: 3.3
```

> Rarely used, but handy in scientific computing.

---

## ✅ 3. **String Type**

A **string** is a sequence of bytes (UTF-8 encoded characters).

* Immutable
* Can be indexed and sliced

```go
var name string = "Skyy"
var greeting = "Hello, " + name
```

#### String Operations:

```go
fmt.Println(len(name))      // Length
fmt.Println(name[0])        // First byte
fmt.Println(name[0:2])      // Slice "Sk"
```

> Strings are byte-based; to work with Unicode characters (like emojis), use `rune`.

---

## ✅ 4. **Rune Type (`rune`)**

* Alias for `int32`
* Represents a **Unicode code point**
* Useful for **multibyte characters**

```go
var ch rune = 'あ' // Japanese character
fmt.Println(ch)     // Unicode point
```

> Runes allow working with characters, not just bytes.

---

## ✅ 5. **Byte Type (`byte`)**

* Alias for `uint8`
* Used for raw binary data or ASCII characters

```go
var b byte = 'A'
```

> Often used in file I/O, buffers, and byte slices (`[]byte`).

---

## 🧠 Type Inference

Go uses type inference with `:=`:

```go
age := 30           // inferred as int
pi := 3.14          // inferred as float64
flag := true        // inferred as bool
```

---

## 🗂 Default Zero Values

| Type      | Default Value |
| --------- | ------------- |
| `bool`    | `false`       |
| `int`     | `0`           |
| `float64` | `0.0`         |
| `string`  | `""` (empty)  |
| `pointer` | `nil`         |

---

## ⚙️ Best Practices

* Use `int` unless a specific size is required (e.g., for network data).
* Use `float64` for decimals unless performance/memory optimization is critical.
* Use `string` for text, `[]byte` for binary data.
* Use type aliases like `byte` and `rune` when appropriate for clarity.

---