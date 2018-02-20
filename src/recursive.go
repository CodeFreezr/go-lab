package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() ([]string, error) {
    searchDir := "c:/path/to/dir"

    fileList := []string{}
    err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
        fileList = append(fileList, path)
        return nil
    })

    for _, file := range fileList {
        fmt.Println(file)
    }
}