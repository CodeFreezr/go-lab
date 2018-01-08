package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/kyokomi/emoji.v1"
)

func main() {
	m := emoji.CodeMap()

	//divider := "👽🎲💎🥨🛸"
	if len(os.Args) == 0 {
		fmt.Fprintf(os.Stderr, "You must specify a unicode-alias or use -all ")
		os.Exit(1)
	}
	arg := os.Args[1]
	sortedKeys := make([]string, 0, len(m))
	for k := range m {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)

	switch arg {
	case "/all", "-all", "--all":
		c := 0
		for _, key := range sortedKeys {
			c++
			noColon := strings.Replace(key, ":", "", -1)
			fmt.Printf("(%v: %s -> %v  ), ", c, noColon, m[key])
		}

	case "/emojis", "-emojis", "--emojis":
		for _, index := range sortedKeys {
			fmt.Printf(m[index])
			fmt.Printf("  ")
		}

	case "/codes", "-codes", "--codes":
		for _, index := range sortedKeys {
			noColon := strings.Replace(index, ":", "", -1)
			fmt.Printf(noColon)
			fmt.Printf(", ")
		}

	case "/?", "/h", "/help", "-help", "-h", "-.help", "--h":
		fmt.Println(`Main:
	moj [string] [-v] 
	List all emojis with shortcodes containing [string]
	Examples: 
		moj smiley
		moj smiley -v

	moj :[string]: [-v]
	Print exact the emoji from the first shortcode
	Excamples:
		moj :smiley:
		moj :smiley: -v

	Verbose Modes (-v) prints some further informations.

Utilities:
	moj -all: 
	Get all emojis with (No: shortcode -> emoji),
	
	moj -codes 
	List all shortcodes comma seperated (a to z)

	moj -emojis
	List all emojis (a to z of its shortcode) 

	moj -version or moj -v
	Gives some Version-Informations about this programm and used unicode and emoji

    `)

	case "/v", "/version", "-v", "-version", "--v", "--version":
		fmt.Println(`Version:
	moj version 0.0.2 (CC) by CodeFreezr, github/codefreezr/emojo/go
	used emoji version: 3.0, based on unicode 8.0
			`)

	default:

		findings := filter(sortedKeys, func(v string) bool {
			return strings.Contains(v, arg)
		})
		fmt.Println(findings)

		/*
				for _, sortedKeys := range likes[arg] {

					fmt.Println(p.Name, "likes cheese.")
				}


			key := ":" + arg + ":"
			fmt.Printf(m[key])
		*/
	}
	/*
		if arg == "-all" {
			//for _, key := range sortedKeys {fmt.Printf("---  %s  ---|", m[key])}#
			fmt.Printf(" 👽  ")
			for _, key := range sortedKeys {
				fmt.Printf("%s 👽 ", key)
			}
		} else {
			key := ":" + arg + ":"
			fmt.Printf(m[key])
		}
	*/
}

// Returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
