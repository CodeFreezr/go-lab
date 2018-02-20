package main

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(filepath.Join("a", "b", "c"))
	fmt.Println(filepath.Join("a", "b/c"))
	fmt.Println(filepath.Join("a/b", "c"))
	fmt.Println(filepath.Join("a/b", "/c"))
}
