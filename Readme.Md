## Additional, curated Reading-Material 📖👨🏻‍💻

Let’s dive deep into **Linux file permissions** — a foundational concept every developer or system user should know, especially when working with servers, deployments, or backend development.

---

## 🔐 Linux File Permissions: Overview

In Linux, every file and directory has **permissions** associated with it that define **who can read, write, or execute** them. These permissions are set for three kinds of users:

### 👥 User Types

| Symbol | Description                                 |
| ------ | ------------------------------------------- |
| `u`    | User (Owner) — the person who owns the file |
| `g`    | Group — a set of users in the same group    |
| `o`    | Others — everyone else on the system        |
| `a`    | All (user + group + others)                 |

---

## 🔤 Permission Types

Each user type can have three kinds of permissions:

| Symbol | Name    | Meaning on File            | Meaning on Directory        |
| ------ | ------- | -------------------------- | --------------------------- |
| `r`    | Read    | Can view file contents     | Can list directory contents |
| `w`    | Write   | Can modify file contents   | Can add/remove files in dir |
| `x`    | Execute | Can run the file as script | Can `cd` into the directory |

---

## 🧾 Viewing File Permissions (`ls -l`)

```bash
ls -l myfile.txt
```

Output example:

```
-rwxr-xr--
```

Breakdown:

* `-` : Type (dash = regular file, `d` = directory)
* `rwx` : Owner has **read, write, execute**
* `r-x` : Group has **read, execute**
* `r--` : Others have **read only**

---

## 🔢 Numeric (Octal) Representation

Each permission is represented by a number:

* `r = 4`
* `w = 2`
* `x = 1`

Sum them up for each user group:

| Permissions | Value |
| ----------- | ----- |
| `rwx`       | 7     |
| `rw-`       | 6     |
| `r--`       | 4     |
| `---`       | 0     |

So:

```bash
chmod 755 myfile.sh
```

Means:

* Owner: `7` → `rwx`
* Group: `5` → `r-x`
* Others: `5` → `r-x`

---

## 🛠️ Changing Permissions

### `chmod` – Change mode

```bash
chmod 644 file.txt   # rw-r--r--
chmod +x script.sh   # adds execute permission
chmod u+x script.sh  # adds execute only for owner
```

### `chown` – Change ownership

```bash
chown skyy:developers file.txt
# User = skyy, Group = developers
```

---

## 🏷 Special Permissions (Advanced)

| Symbol | Name          | Description                                      |
| ------ | ------------- | ------------------------------------------------ |
| `s`    | setuid/setgid | Executes with owner/group privileges             |
| `t`    | sticky bit    | Only owner can delete files (commonly on `/tmp`) |

Example: `chmod 4755 file` sets **setuid** (4 before permissions).

---

## 🔐 Typical Use-Cases

| Use Case                    | Command                 |
| --------------------------- | ----------------------- |
| Make a script executable    | `chmod +x script.sh`    |
| Give read/write to everyone | `chmod 666 file.txt`    |
| Secure a private file       | `chmod 600 secrets.txt` |
| Make directory accessible   | `chmod 755 /myfolder`   |
| Make file owner-only        | `chmod 700 private.sh`  |

---

## 🧪 Test It Out!

Create a file:

```bash
touch test.txt
ls -l test.txt
chmod 600 test.txt
ls -l test.txt
```

---

Now.. Let's walk through **Linux file permissions in the context of deploying a Go (Golang) app** on a Linux server — this is *super important* for security and functionality.

---

## 🧑‍💻 Scenario: Deploying a Go Web App on a Linux Server

Let's say we have a Go app with the following structure:

```
/home/skyy/my-go-app/
├── main.go
├── app        # compiled binary
├── config.env
└── static/    # static files (HTML, CSS, JS)
```

We want:

* Only **Skyy** (the owner) to run and modify the app.
* The server (e.g., Nginx or systemd) to **execute** the Go binary.
* The public to access only what's needed (like `static/` files if served directly).
* **No one else** (other users) to view configs or tamper with code.

---

## ✅ Step-by-Step File Permission Setup

### 1. **Binary (`app`)**

```bash
chmod 755 app
```

Means:

* `7` for owner (read/write/execute)
* `5` for group (read/execute)
* `5` for others (read/execute)

> ✅ Allows executing the app, but **only the owner can modify** it.

### 2. **Main Go Source File (`main.go`)**

```bash
chmod 644 main.go
```

* Owner: read/write
* Group/Others: read

> ✅ Useful during dev; in production you might remove it or set it to `600`.

### 3. **Environment Config (`config.env`)**

```bash
chmod 600 config.env
```

> 🔒 **Very important**: Only the owner can read/write it. Prevents exposing secrets.

### 4. **Static Assets (`static/`)**

```bash
chmod -R 755 static/
```

* `-R`: recursively sets permissions

> ✅ Allows public read access for web server to serve files, but no write access.

---

## 🔐 Ownership Control (Optional)

If your Go app runs as a different Linux user (e.g., via a systemd service), you might:

```bash
chown root:www-data app
chmod 750 app
```

> This allows the app to be run by `www-data` (like Nginx), but **not by everyone**.

---

## 🔄 Sample: Full Command Setup

```bash
cd /home/skyy/my-go-app

chmod 755 app             # executable binary
chmod 600 config.env      # protect secrets
chmod -R 755 static/      # public static files
chmod 644 main.go         # source file (for development only)
```

---

## 🧪 Quick Test

To confirm permission levels:

```bash
ls -l
# OR to view recursively
ls -lR
```

---

## ✅ Pro Tips

| Tip                                   | Why                                                                        |
| ------------------------------------- | -------------------------------------------------------------------------- |
| Use `.env` file permissions carefully | They often store API keys or DB credentials                                |
| Don’t make everything 777             | That’s a security disaster — gives **everyone full control**               |
| Use `sudo` only when needed           | You don’t want to accidentally give root-level access to wrong files       |
| Make logs readonly                    | e.g., `chmod 644 logs.txt` to allow viewing but not editing logs by others |

---

