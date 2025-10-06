# 🧭 Go4it — Complete Go Setup (macOS)

Welcome to **Go4it** — your step-by-step guide to installing **Go (Golang)**, configuring your environment on macOS. 🚀

---

## 🧱 1️⃣ Check if Go is Already Installed

Open **Terminal** and run:
```bash
 go version
```
If you see:
```bash
go version go1.23.2 darwin/arm64
```
✅ Go is already installed.

If not, follow the steps below 👇

## 💿 2️⃣ Install Go on macOS
### 🪄 Option 1 — Official Installer (Recommended) <br>
	1.	Visit https://go.dev/dl/ <br>
	2.	Download the macOS .pkg installer <br>
	    •	Choose ARM64 for Apple Silicon (M1/M2/M3) <br>
	    •	Choose AMD64 for Intel Macs <br>
	3.	Run the installer — Go will be installed at:
  ```bash
    /usr/local/go
  ```
  Verify installation:
  ```bash
    go version
  ```

### 🧰 Option 2 — Install via Homebrew
If you prefer Homebrew, run:
```bash
   brew install go
```
Verify:
```bash
   go version
```

## ⚙️ 3️⃣ Configure Environment Variables
### Step 1 — Check Go Installation Path
```GoLang
   which go
```
Expected:
```bash
   /usr/local/go/bin/go
```

### Step 2 — Add Environment Variables
🧩 For Zsh (default on macOS)
Edit your .zshrc:
```bash
   nano ~/.zshrc
```
Add:
```bash
   # Go environment setup
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
Save (Ctrl + O, Enter, Ctrl + X) and reload:
```bash
   source ~/.zshrc
```

🧩 For Bash
Edit:
```bash
   nano ~/.bash_profile
```
Add:
```bash
   export GOPATH=$HOME/go
   export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```
Reload:
```bash
   source ~/.bash_profile
```
## 🧪 4️⃣ Verify Go Environment Setup
Check your Go environment:
```bash
   go env
```
Expected values:
```bash
   GOPATH="/Users/yourname/go"
   GOROOT="/usr/local/go"
   PATH="...:/usr/local/go/bin:/Users/yourname/go/bin"
```

Diagnostic script to check the Go environment:
1. Create a script to print Go environment info:
```bash
   nano diagnostics.go
```
2. Add:
```bash
   package main

import (
    "fmt"
    "os"
    "runtime"
)

func main() {
    fmt.Println("Go Version:", runtime.Version())
    fmt.Println("GOOS:", runtime.GOOS)
    fmt.Println("GOARCH:", runtime.GOARCH)
    fmt.Println("GOPATH:", os.Getenv("GOPATH"))
    fmt.Println("GOROOT:", os.Getenv("GOROOT"))
    fmt.Println("PATH:", os.Getenv("PATH"))
}
```
3. Run:
```bash
   go run diagnostics.go
```
