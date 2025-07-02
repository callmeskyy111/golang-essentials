Let’s explore **Goroutines in GoLang** in **depth** — they're one of the core reasons Go is known for its **concurrent programming capabilities**.

---

## 🔷 What is a Goroutine?

A **goroutine** is a lightweight thread managed by the **Go runtime**, not the operating system.

✅ **Key points:**

* It allows us to **run functions concurrently**.
* They are **cheaper than threads** (in terms of memory and scheduling).
* They’re **non-blocking**, meaning they don’t halt the main execution flow.

> You can think of goroutines as **functions running in the background** or **asynchronously**.

---

## 🔷 How to Start a Goroutine

You start a goroutine by using the keyword `go` before a function call:

```go
go sayHello()
```

---

## 🔷 Example – Basic Goroutine

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from Goroutine!")
}

func main() {
    go sayHello() // This runs concurrently

    fmt.Println("Main function execution")
    time.Sleep(1 * time.Second) // Give the goroutine time to finish
}
```

> 🧠 If we don’t `Sleep` or wait, the goroutine may not run because the main function can exit too fast.

---

## 🔷 Goroutine Characteristics

| Feature       | Description                                                 |
| ------------- | ----------------------------------------------------------- |
| Lightweight   | Uses **2KB** of stack memory initially (OS threads use MBs) |
| Multiplexed   | Many goroutines share fewer OS threads (M\:N model)         |
| Scalable      | You can spawn **thousands** of goroutines easily            |
| Managed by Go | Go scheduler handles their creation, execution, suspension  |

---

## 🔷 Anonymous Function in a Goroutine

You can run an inline (anonymous) function concurrently:

```go
go func() {
    fmt.Println("Running anonymous goroutine")
}()
```

You can even pass arguments:

```go
go func(name string) {
    fmt.Println("Hello", name)
}("Skyy")
```

---

## 🔷 Wait for Goroutines: `sync.WaitGroup`

To wait for goroutines to finish properly, use `sync.WaitGroup`.

### Example:

```go
package main

import (
    "fmt"
    "sync"
)

func printName(wg *sync.WaitGroup, name string) {
    defer wg.Done()
    fmt.Println("Hello", name)
}

func main() {
    var wg sync.WaitGroup

    wg.Add(2) // Expecting 2 goroutines

    go printName(&wg, "Skyy")
    go printName(&wg, "Banerjee")

    wg.Wait() // Wait until both goroutines finish
    fmt.Println("All done!")
}
```

---

## 🔷 Common Use-Cases for Goroutines

* Running **I/O-bound operations** (e.g., HTTP requests, file reads)
* Performing **parallel computations**
* Handling **concurrent tasks** like processing data streams, or background jobs

---

## 🔷 Goroutines vs Threads

| Feature        | Goroutine           | Thread             |
| -------------- | ------------------- | ------------------ |
| Memory usage   | \~2 KB              | \~1 MB             |
| Managed by     | Go runtime          | OS                 |
| Creation speed | Very fast           | Slower             |
| Scheduling     | Cooperative (by Go) | Preemptive (by OS) |
| Scalability    | Thousands possible  | Limited            |

---

## 🔷 Real Example – Fetching Multiple URLs Concurrently

```go
package main

import (
    "fmt"
    "net/http"
    "sync"
)

func fetchURL(wg *sync.WaitGroup, url string) {
    defer wg.Done()
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(url, "status:", resp.Status)
}

func main() {
    var wg sync.WaitGroup
    urls := []string{
        "https://google.com",
        "https://golang.org",
        "https://github.com",
    }

    for _, url := range urls {
        wg.Add(1)
        go fetchURL(&wg, url)
    }

    wg.Wait()
}
```

---

## 🔷 Caveats / Things to Watch Out For

1. **Race conditions**: When multiple goroutines read/write shared data.

   * Use `sync.Mutex` or channels to manage state safely.

2. **Leaking goroutines**: Goroutines that never exit (e.g., blocked on a channel).

   * Always ensure goroutines can finish or timeout.

3. **Avoid starting too many goroutines** without need.

---

## 🔷 Summary

| Concept         | Description                                               |
| --------------- | --------------------------------------------------------- |
| `go func()`     | Runs a function concurrently as a goroutine               |
| Lightweight     | Starts with \~2 KB memory                                 |
| Concurrent      | Multiple can run independently                            |
| Use `WaitGroup` | To wait for them to complete                              |
| Use `channels`  | For communication between them (we can go into this next) |

---