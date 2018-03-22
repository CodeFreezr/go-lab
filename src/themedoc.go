package main

import (
	"os"
	"path/filepath"
	"strconv"
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
	searchDir := "D:/dbt/01/gitlab/graphugovizart/themes/blackburn" //os.Args[1]
	//ext := ".go$|.png$"

	//fileList := []gofi{}
	graph := "graph theme {\n rankdir=LR\n node [shape=MRecord style=rounded]\n"
	graph = graph + " graph [bgcolor=grey15 color=black style=filled;style=rounded;]\n"

	gcount := 0
	fcount := 0
	lastdir := false
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fcount++
			//reg, err := regexp.MatchString(ext, f.Name())
			//if err == nil && reg {
			//}

			//p := strings.Replace(path, searchDir, "", 1)

			println(fcount, ". file: ", f.Name())
			graph = graph + "\"" + f.Name() + "\" "
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

		} else {
			gcount++
			if gcount > 1 && !lastdir {
				graph = graph + "}"
			}
			graph = graph + "\n"
			print(gcount, " - subgraph: ")
			graph = graph + "\tsubgraph cluster_" + strconv.Itoa(gcount) + "{label=\"" + f.Name() + "\" "
			println(f.Name())

			//fmt.Println(f.Name(), path)
		}
		lastdir = f.IsDir()

		return err
	})

	if err != nil {
		panic(err)
	}
	graph = graph + "}\n}"
	println("---------------------------------------------------------------")
	println(graph)
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
