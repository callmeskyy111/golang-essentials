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

Let’s dive **deep** into **`sync.WaitGroup`** in Go — a **powerful synchronization tool** used with **goroutines** to **wait until a group of concurrent tasks finish**.

---

## 🧠 What is `sync.WaitGroup`?

A `WaitGroup` is a **struct** provided by the `sync` package in Go that helps **wait for a collection of goroutines to finish executing**.

### ✅ Why use it?

When we spawn multiple goroutines, they run **asynchronously**. The main function may finish before the goroutines complete. `WaitGroup` ensures the **main (or parent) goroutine waits** for others.

---

## 🔧 Basic Structure

```go
var wg sync.WaitGroup
```

We typically use **three methods**:

| Method      | Purpose                                        |
| ----------- | ---------------------------------------------- |
| `wg.Add(n)` | Add `n` number of goroutines to wait for       |
| `wg.Done()` | Called inside a goroutine when it finishes     |
| `wg.Wait()` | Block until all `.Done()` calls match `.Add()` |

---

## 🧪 Simple Example

```go
package main

import (
    "fmt"
    "sync"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Tells WaitGroup this goroutine is done
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    wg.Add(3) // We expect 3 goroutines

    go worker(1, &wg)
    go worker(2, &wg)
    go worker(3, &wg)

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All workers finished")
}
```

### 🧵 What happens here:

* We **register** 3 tasks (`Add(3)`)
* Each goroutine **calls `Done()` once**
* `Wait()` blocks until all `Done()` calls are made
* Then, the main function continues

---

## 🔄 How `.Add()` and `.Done()` Work Together

Imagine `WaitGroup` maintains an internal **counter**:

1. `Add(n)` increases the counter by `n`
2. `Done()` decreases it by 1
3. `Wait()` blocks until the counter is 0

So:

```go
wg.Add(3)
wg.Done() // Now 2 remaining
wg.Done() // Now 1 remaining
wg.Done() // Now 0 → wg.Wait() unblocks
```

---

## ⚠️ IMPORTANT: Always call `.Done()` once per `.Add()`

If you forget `.Done()`:

* `Wait()` will **block forever**

If you call too many `.Done()`:

* You’ll get a **panic: negative WaitGroup counter**

💡 **Use `defer wg.Done()`** at the top of your goroutines to avoid forgetting.

---

## 💥 What if you add `.Add()` **inside a loop**?

Yes, you can do this dynamically too!

```go
for _, item := range list {
    wg.Add(1)
    go func(val string) {
        defer wg.Done()
        fmt.Println(val)
    }(item)
}
```

⚠️ **Pass the loop variable as an argument**, or else you might capture the wrong value due to goroutine closure behavior.

---

## ✅ Real-life Use Case Example

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
        fmt.Println("Error fetching:", url)
        return
    }
    fmt.Println(url, "->", resp.Status)
}

func main() {
    var wg sync.WaitGroup
    urls := []string{
        "https://golang.org",
        "https://google.com",
        "https://github.com",
    }

    for _, url := range urls {
        wg.Add(1)
        go fetchURL(&wg, url)
    }

    wg.Wait()
    fmt.Println("All fetches complete")
}
```

---

## 🧭 Summary

| Feature          | Explanation                                     |
| ---------------- | ----------------------------------------------- |
| `sync.WaitGroup` | A struct to wait for multiple goroutines        |
| `Add(n)`         | Increments internal counter by `n`              |
| `Done()`         | Decrements counter by 1                         |
| `Wait()`         | Blocks until counter reaches 0                  |
| Thread-safe      | Can be used concurrently by multiple goroutines |
| Use `defer`      | Helps ensure `Done()` is always called          |

---

## 🧪 Tips & Best Practices

* Always **pass `*sync.WaitGroup` by reference** to avoid copying.
* Use **`defer wg.Done()`** at the **start of each goroutine**.
* Avoid modifying the WaitGroup **after calling `Wait()`** — it's not safe.
* You can use `WaitGroup` along with **channels** for more complex patterns.

---

Let's walk through this Go program **step by step**, explaining every part so it's **crystal clear**, especially for someone new to Go and goroutines.

---

## 🧾 Full Code Recap:

```go
package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Doing Task:", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}
```

---

## 🧠 Step-by-Step Explanation

---

### 📦 `package main`

This tells the Go compiler:

> “This is the main package, and it’s the **entry point** for execution.”

---

### 📚 `import (...)`

We’re importing two packages:

* `fmt`: For printing to the console.
* `sync`: Provides `WaitGroup`, which helps us **wait for goroutines to finish**.

---

### 🧵 `func task(id int, w *sync.WaitGroup)`

This is a **function that runs in a goroutine**. It takes:

* `id int`: Just an integer to identify which task this is.
* `w *sync.WaitGroup`: A pointer to a WaitGroup (used to notify that this task is done).

Inside:

```go
defer w.Done()
```

This **schedules** `w.Done()` to be called **when this function finishes** — a safe way to ensure it’s always called, even if there’s an error or early return.

Then:

```go
fmt.Println("Doing Task:", id)
```

Prints the task number.

---

### 🏁 `func main()`

The main function where execution starts.

---

### 🧰 `var wg sync.WaitGroup`

Here we declare a **WaitGroup**, named `wg`, to wait for our goroutines to finish.

---

### 🔁 `for i := 0; i <= 10; i++ { ... }`

This loop runs **11 times** — from `0` to `10`.
Each time:

1. `wg.Add(1)` — tells the WaitGroup:

   > “I’m starting one more task.”
2. `go task(i, &wg)` — launches a new **goroutine** to run the `task` function with `i`.

### 🧠 What is a **goroutine**?

A **lightweight thread** managed by the Go runtime. It allows us to run code **concurrently**, i.e., multiple things at once.
It’s launched with the keyword `go`:

```go
go task(i, &wg)
```

That line doesn't block — it **starts running in the background**.

---

### ⏳ `wg.Wait()`

This tells the main goroutine:

> “Wait here until the counter in the WaitGroup becomes 0.”

Each `task()` goroutine, when done, calls `wg.Done()`, which decreases the counter by 1.

Once all 11 goroutines have called `Done()`, `wg.Wait()` unblocks, and the program ends.

---

### 📥 Commented Part

```go
// go func(i int) {
//     fmt.Println("Internal GoRoutine: ", i)
// }(i)
```

This is an **anonymous inline goroutine** doing the same thing as `task()`, but without a named function.

Passing `i` as an argument to the anonymous function is important — otherwise, due to closure behavior, all goroutines could capture the **same final value of `i`**.

---

### 🖨 Sample Output

```
Doing Task: 1
Doing Task: 10
Doing Task: 6
Doing Task: 7
Doing Task: 8
Doing Task: 9
Doing Task: 3
Doing Task: 5
Doing Task: 2
Doing Task: 0
Doing Task: 4
```

Notice: **The order is random.**
Why? Because goroutines run **asynchronously** — the Go runtime decides when to execute each one.

---

## 🔍 Why WaitGroup is important?

Without `wg.Wait()`, the **main function may finish early**, ending the program **before** the goroutines print anything.

If we remove `wg.Wait()`:

* The main goroutine would exit
* Other goroutines might not get time to finish

---

## 🛑 Rules to Remember

| Rule                                         | Reason                          |
| -------------------------------------------- | ------------------------------- |
| Use `go` to start a goroutine                | To run concurrently             |
| Use `wg.Add(n)` before starting n goroutines | So we know how many to wait for |
| Use `defer wg.Done()` inside goroutine       | So it always decrements         |
| Use `wg.Wait()` in main or parent goroutine  | To wait for all to finish       |
| Always pass `*sync.WaitGroup`                | It's a pointer, don’t copy it   |

---

## ✅ What We’ve Covered

* What goroutines are
* How to start goroutines in a loop
* What WaitGroup is and how to use it properly
* Why order of output is not guaranteed
* How `defer` works with `wg.Done()`

---

