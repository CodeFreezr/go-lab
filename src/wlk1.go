package main

import (
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type gofi struct {
	Name    string
	Path    string
	User    string
	Repo    string
	Bare    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
}

func main() {
	searchDir := os.Args[1]
	ext := ".go$"

	//fileList := []gofi{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			reg, err := regexp.MatchString(ext, f.Name())
			if err == nil && reg {
				//println(f.Name())
				/*
					user := strings.Split(path, ".")[0]
					bare := strings.Replace(path, user, "", 1)
					bare = strings.Replace(bare, f.Name(), "", -1)
					bare = strings.TrimRight(bare, "\\")
					bare = strings.TrimLeft(bare, ".")
					repo := strings.Split(bare, "\\")[0]
					bare = strings.Replace(bare, repo, "", 1)
					bare = strings.TrimLeft(bare, "\\")
					//fileList = append(fileList, path)
					fileList = append(fileList, gofi{f.Name(), path, user, repo, bare, f.Size(), f.Mode(), f.ModTime(), f.IsDir()})
				*/
			}
		} else {
			println(path)

			//fmt.Println(f.Name(), path)
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	/*
		for _, file := range fileList {
			fmt.Println("----- ")
			fmt.Println("User: ", file.User)
			fmt.Println("Repo: ", file.Repo)
			fmt.Println("File: ", file.Name)
			fmt.Println("Path: ", file.Bare)
			fmt.Println("Full: ", file.Path)

		}
	*/
}
