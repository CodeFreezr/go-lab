package main

import (
	"fmt"
	"os"

	"gopkg.in/kyokomi/emoji.v1"
)

func main() {
	m := emoji.CodeMap()
	arg := os.Args[1]

	if len(arg) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a unicode-alias or use -all ")
		os.Exit(1)
	}

	if arg == "-all" {
		for index, key := range sortedKeys {
			fmt.Printf("%v %s %s \n", index, key, m[key])
		}
	} else {
		key := ":" + arg + ":"
		fmt.Printf(m[key])
	}

}
