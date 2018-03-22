package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

//func readit(dirname string) ([]os.FileInfo, error) {
func readit(dirname string) {
	f, err := os.Open(dirname)
	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}
	list, err := f.Readdir(-1)
	println(f.Name())
	f.Close()

	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })

	for _, list := range list {
		if list.IsDir() {
			fmt.Println(list.IsDir(), list.Name())
		}
	}

	//return list, nil
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)

	readit(pwd)
	/*
		if err != nil {
			log.Fatal(err)
		}
		println(files)
	*/
	/*
		for _, file := range files {
			fmt.Println(file.Name())
		}
	*/
}
