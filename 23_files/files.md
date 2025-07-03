Let’s dive **in-depth into file handling in Go**, covering everything from reading/writing to advanced operations like buffering, appending, and error handling.

---

## 📁 What Is File Handling in Go?

In Go, **file handling** is mainly done using the `os`, `io`, `bufio`, and `ioutil` (deprecated in Go 1.16) or `io/ioutil` packages.

You can:

* Create files
* Open files for reading or writing
* Read/write line-by-line or byte-by-byte
* Use buffered I/O for efficiency
* Handle file metadata (permissions, size, etc.)
* Close files properly

---

## 🔧 1. Import Packages

```go
import (
	"fmt"
	"os"
	"bufio"
	"io"
)
```

---

## 📂 2. Creating a File

```go
file, err := os.Create("example.txt")
if err != nil {
	fmt.Println("Error creating file:", err)
	return
}
defer file.Close()

file.WriteString("Hello, Skyy!\n")
```

* `os.Create`: Creates the file if it doesn’t exist; **overwrites if it does**
* `file.WriteString`: Writes a string
* `defer file.Close()`: Ensures the file is closed properly

---

## 📖 3. Reading a File

### a) Read entire content

```go
data, err := os.ReadFile("example.txt")
if err != nil {
	fmt.Println("Error reading file:", err)
	return
}
fmt.Println(string(data))
```

> 🔄 Replaces the deprecated `ioutil.ReadFile`

---

### b) Read line-by-line using `bufio.Scanner`

```go
file, err := os.Open("example.txt")
if err != nil {
	fmt.Println("Error opening file:", err)
	return
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
	fmt.Println(scanner.Text()) // each line
}

if err := scanner.Err(); err != nil {
	fmt.Println("Error during scanning:", err)
}
```

✅ Best for **log files** or large files, where reading everything at once is inefficient.

---

### c) Read byte-by-byte

```go
file, _ := os.Open("example.txt")
defer file.Close()

buffer := make([]byte, 10) // read in 10-byte chunks
for {
	n, err := file.Read(buffer)
	if err == io.EOF {
		break
	}
	fmt.Print(string(buffer[:n]))
}
```

✅ Useful for **custom parsers**, streaming, or binary files.

---

## ✍️ 4. Writing to a File

### a) Overwriting content

```go
file, _ := os.Create("example.txt") // overwrites
file.WriteString("Overwritten!\n")
file.Close()
```

### b) Appending content

```go
file, _ := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
defer file.Close()

file.WriteString("Appended line\n")
```

> `os.O_APPEND`: Append mode
> `os.O_WRONLY`: Write-only
> `0644`: File permission (user: read/write, group+others: read)

---

## ⚡ 5. Buffered Writing

```go
file, _ := os.Create("buffered.txt")
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("Buffered write\n")
writer.Flush() // must flush to write to file
```

> ✅ Buffered writing is faster, especially when writing lots of small strings.

---

## 🧼 6. Deleting and Renaming Files

```go
// Delete
err := os.Remove("example.txt")

// Rename
err := os.Rename("old.txt", "new.txt")
```

---

## 📊 7. Getting File Info

```go
info, err := os.Stat("example.txt")
if err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Println("File Name:", info.Name())
fmt.Println("Size:", info.Size())
fmt.Println("Modified:", info.ModTime())
fmt.Println("Permissions:", info.Mode())
```

---

## 💥 8. Error Handling Tips

Always handle:

* `os.IsNotExist(err)` → file not found
* `os.IsPermission(err)` → permission denied
* `scanner.Err()` → scanner failure

Example:

```go
if os.IsNotExist(err) {
	fmt.Println("File does not exist.")
}
```

---

## 🧪 Real-Life Example: Append Logs to File

```go
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	log := fmt.Sprintf("Log at %v\n", time.Now())

	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	file.WriteString(log)
	fmt.Println("Log written.")
}
```

✅ Appends a timestamp to `app.log` every time it runs.

---

## 🔄 File Modes in `os.OpenFile`

| Constant      | Description          |
| ------------- | -------------------- |
| `os.O_RDONLY` | Read-only            |
| `os.O_WRONLY` | Write-only           |
| `os.O_RDWR`   | Read + Write         |
| `os.O_APPEND` | Append mode          |
| `os.O_CREATE` | Create if not exists |
| `os.O_TRUNC`  | Truncate if exists   |

You can combine them using `|` (bitwise OR).

---

## ✅ Summary

| Action          | Method                                   |
| --------------- | ---------------------------------------- |
| Create file     | `os.Create()`                            |
| Open file       | `os.Open()` or `os.OpenFile()`           |
| Read whole file | `os.ReadFile()`                          |
| Read lines      | `bufio.NewScanner()`                     |
| Write/append    | `WriteString()`, `OpenFile()` with flags |
| Delete/rename   | `os.Remove()`, `os.Rename()`             |
| File info       | `os.Stat()`                              |

---
