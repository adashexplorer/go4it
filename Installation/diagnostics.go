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
