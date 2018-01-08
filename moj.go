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

	arg := ""
	if len(os.Args) < 2 {
		arg = "-help"
	} else {
		arg = os.Args[1]
	}

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
			fmt.Printf("(%v: %s -> %v  ) ", c, noColon, m[key])
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
			fmt.Printf(" ")
		}

	case "/?", "/h", "/help", "-help", "-h", "-.help", "--h":
		fmt.Println(`Syntax:
	moj [shortcode] [modifier]
	
	shortcode (optional):
		:smile:	-> Display emoji with shortcode ":smile:"
		:smile	-> List emojis where shortcode >>starts<< with ":smile"
		smile:	-> List emojis where shortcode >>ends<< with "smile:"
		smile	-> List emojis with shortcodes >>contain<< "smile"

	modifier (optional:
		-all	Displays (No.: shortcode -> emoji)
		-codes	Displays only the shortcodes
		-emojis	List all emojis. Can be omitted if shortcode is given
		(-info	Displays further Information for exact one emoji -> tbd)

	moj -v	Gives some informations about this programm and used unicode- and emoji-version
    `)

	case "/v", "/version", "-v", "-version", "--v", "--version":
		fmt.Println(`Version:
	moj version 0.0.9 (CC) by 💡+🤓  CodeFreezr❄️🖖, github/codefreezr/emojo/go
	used emoji version: 3.0, based on unicode 8.0 (C) by unicode.org`)

	default:

		findings := filter(sortedKeys, func(v string) bool { return strings.Contains(v, arg) })

		if len(findings) == 0 {
			fmt.Println("Nothing found, try again.")
		}

		if len(os.Args) < 3 {
			// No modifiers, just emojis
			for _, index := range findings {
				fmt.Printf(m[index])
				fmt.Printf("  ")
			}
		} else {
			arg2 := os.Args[2]
			switch arg2 {
			case "-codes":
				for _, index := range findings {
					noColon := strings.Replace(index, ":", "", -1)
					fmt.Printf(noColon)
					fmt.Printf(" ")
				}
			case "-all":
				d := 0
				for _, key := range findings {
					d++
					noColon := strings.Replace(key, ":", "", -1)
					fmt.Printf("(%v: %s -> %v  ) ", d, noColon, m[key])
				}
			case "-info":
				fmt.Println(`TBD: Some informations on `, arg)
			}
		}
	}
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
