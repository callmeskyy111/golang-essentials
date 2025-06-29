Let's dive deep into the `**range**` keyword in **Go (Golang)** — one of the most powerful tools in the language for looping over **slices**, **arrays**, **maps**, **strings**, and **channels**.

---

## 🔑 What is `range` in Go?

`range` is a **Go keyword** used in `for` loops to **iterate over**:

* Slices and arrays
* Maps
* Strings
* Channels

It returns one or two values depending on the type of the object being ranged over.

---

## 🧪 1. Range Over **Slice / Array**

```go
nums := []int{10, 20, 30}

for index, value := range nums {
    fmt.Println(index, value)
}
```

### ✅ Output:

```
0 10
1 20
2 30
```

* First value: **index**
* Second value: **element value**

#### 💡 Ignore index if not needed:

```go
for _, value := range nums {
    fmt.Println(value)
}
```

---

## 🗺️ 2. Range Over **Map**

```go
m := map[string]int{"a": 1, "b": 2}

for key, value := range m {
    fmt.Println(key, value)
}
```

* First value: **key**
* Second value: **value**

⚠️ Map iteration order is **not guaranteed** to be the same every time.

---

## 🔤 3. Range Over **String**

```go
s := "gölang"

for i, ch := range s {
    fmt.Printf("%d -> %c\n", i, ch)
}
```

### ✅ Output:

```
0 -> g
1 -> ö
3 -> l
4 -> a
5 -> n
6 -> g
```

* First value: **byte index**
* Second value: **rune (Unicode code point)**

### 💡 Note:

* `ö` takes **2 bytes** in UTF-8 (which is why `l` starts at index 3).

---

## 📡 4. Range Over **Channel**

```go
ch := make(chan int)

go func() {
    for i := 1; i <= 3; i++ {
        ch <- i
    }
    close(ch)
}()

for val := range ch {
    fmt.Println(val)
}
```

* Iterates over values **received from the channel** until it's closed.

---

## 🧠 Internal Mechanics of `range`

### When used with a **slice or array**:

```go
for i, v := range mySlice
```

Behind the scenes:

* A **copy** of the slice is made (not elements, just header).
* Index is incremented each time.
* Value is retrieved via `slice[i]`.

### When used with a **map**:

* Go internally uses a **randomized iteration order** to help prevent developers from relying on order.
* Iterates until all keys are exhausted.

### When used with a **string**:

* Decodes **UTF-8 runes**.
* Each rune is returned with its **starting byte index**.

---

## ⚠️ Common Mistakes

### 1. **Modifying a slice while ranging over it** can lead to bugs:

```go
for i, v := range slice {
    slice = append(slice, v*2) // may cause infinite loop
}
```

### 2. **Loop variables are reused**, so closures capture the **same reference**:

```go
func makeFuncs() []func() {
    var funcs []func()
    for i, v := range []string{"a", "b", "c"} {
        funcs = append(funcs, func() {
            fmt.Println(i, v)
        })
    }
    return funcs
}
```

All functions will print the **same values** of `i` and `v`.

✅ Fix:

```go
for i, v := range []string{"a", "b", "c"} {
    i, v := i, v // create new copies
    funcs = append(funcs, func() {
        fmt.Println(i, v)
    })
}
```

---

## 🧪 Real-World Examples

### 🛒 Total price from a slice of product prices:

```go
prices := []float64{100.0, 250.0, 399.99}
var total float64

for _, price := range prices {
    total += price
}
```

---

### 📈 Counting characters in a string:

```go
text := "hello world"
freq := make(map[rune]int)

for _, ch := range text {
    freq[ch]++
}
```

---

## 📌 Summary Table

| Data Type   | First Value | Second Value        |
| ----------- | ----------- | ------------------- |
| Slice/Array | Index       | Element value       |
| Map         | Key         | Value               |
| String      | Byte index  | Rune (Unicode char) |
| Channel     | Value       | —                   |

---

## 🔍 Bonus: Zero-value iteration

You can use only one value:

```go
for i := range arr {}  // index only
for k := range myMap {} // keys only
```

Or ignore both:

```go
for range [5]int{} {
    fmt.Println("hello")
}
```
---


Let's dive deep into the `**range**` keyword in **Go (Golang)** — one of the most powerful tools in the language for looping over **slices**, **arrays**, **maps**, **strings**, and **channels**.

---

## 🔑 What is `range` in Go?

`range` is a **Go keyword** used in `for` loops to **iterate over**:

* Slices and arrays
* Maps
* Strings
* Channels

It returns one or two values depending on the type of the object being ranged over.

---

## 🧪 1. Range Over **Slice / Array**

```go
nums := []int{10, 20, 30}

for index, value := range nums {
    fmt.Println(index, value)
}
```

### ✅ Output:

```
0 10
1 20
2 30
```

* First value: **index**
* Second value: **element value**

#### 💡 Ignore index if not needed:

```go
for _, value := range nums {
    fmt.Println(value)
}
```

---

## 🗺️ 2. Range Over **Map**

```go
m := map[string]int{"a": 1, "b": 2}

for key, value := range m {
    fmt.Println(key, value)
}
```

* First value: **key**
* Second value: **value**

⚠️ Map iteration order is **not guaranteed** to be the same every time.

---

## 🔤 3. Range Over **String**

```go
s := "gölang"

for i, ch := range s {
    fmt.Printf("%d -> %c\n", i, ch)
}
```

### ✅ Output:

```
0 -> g
1 -> ö
3 -> l
4 -> a
5 -> n
6 -> g
```

* First value: **byte index**
* Second value: **rune (Unicode code point)**

### 💡 Note:

* `ö` takes **2 bytes** in UTF-8 (which is why `l` starts at index 3).

---

## 📡 4. Range Over **Channel**

```go
ch := make(chan int)

go func() {
    for i := 1; i <= 3; i++ {
        ch <- i
    }
    close(ch)
}()

for val := range ch {
    fmt.Println(val)
}
```

* Iterates over values **received from the channel** until it's closed.

---

## 🧠 Internal Mechanics of `range`

### When used with a **slice or array**:

```go
for i, v := range mySlice
```

Behind the scenes:

* A **copy** of the slice is made (not elements, just header).
* Index is incremented each time.
* Value is retrieved via `slice[i]`.

### When used with a **map**:

* Go internally uses a **randomized iteration order** to help prevent developers from relying on order.
* Iterates until all keys are exhausted.

### When used with a **string**:

* Decodes **UTF-8 runes**.
* Each rune is returned with its **starting byte index**.

---

## ⚠️ Common Mistakes

### 1. **Modifying a slice while ranging over it** can lead to bugs:

```go
for i, v := range slice {
    slice = append(slice, v*2) // may cause infinite loop
}
```

### 2. **Loop variables are reused**, so closures capture the **same reference**:

```go
func makeFuncs() []func() {
    var funcs []func()
    for i, v := range []string{"a", "b", "c"} {
        funcs = append(funcs, func() {
            fmt.Println(i, v)
        })
    }
    return funcs
}
```

All functions will print the **same values** of `i` and `v`.

✅ Fix:

```go
for i, v := range []string{"a", "b", "c"} {
    i, v := i, v // create new copies
    funcs = append(funcs, func() {
        fmt.Println(i, v)
    })
}
```

---

## 🧪 Real-World Examples

### 🛒 Total price from a slice of product prices:

```go
prices := []float64{100.0, 250.0, 399.99}
var total float64

for _, price := range prices {
    total += price
}
```

---

### 📈 Counting characters in a string:

```go
text := "hello world"
freq := make(map[rune]int)

for _, ch := range text {
    freq[ch]++
}
```

---

## 📌 Summary Table

| Data Type   | First Value | Second Value        |
| ----------- | ----------- | ------------------- |
| Slice/Array | Index       | Element value       |
| Map         | Key         | Value               |
| String      | Byte index  | Rune (Unicode char) |
| Channel     | Value       | —                   |

---

## 🔍 Bonus: Zero-value iteration

You can use only one value:

```go
for i := range arr {}  // index only
for k := range myMap {} // keys only
```

Or ignore both:

```go
for range [5]int{} {
    fmt.Println("hello")
}
```

---

