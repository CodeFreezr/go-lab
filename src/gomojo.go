package main

import (
	"fmt"
	"sort"

	"gopkg.in/kyokomi/emoji.v1"
)

func main() {

	m := emoji.CodeMap()

	sortedKeys := make([]string, 0, len(m))

	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	for index, key := range sortedKeys {
		fmt.Printf("%v %s %s \n", index, key, m[key])

	}
	fmt.Printf("-----------------------------------------------------------------\n")
	for _, key := range sortedKeys {
		//fmt.Printf("%s %s \n", key, m[key])
		fmt.Printf("%s ", m[key])
	}

}
