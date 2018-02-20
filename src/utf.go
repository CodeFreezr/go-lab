package main

import (
	"fmt"
	"os"
)

func main() {

	arg := os.Args[1]

	if arg == "-all" {

		//sr := '\u212C'
		for i := 160000; i < 500000; i++ {

			sr1 := '\u0001'
			sr := sr1 + rune(i)
			//fmt.Printf("%U: ", sr)  // print proper Formated Codepoint
			fmt.Printf("%c  ", sr) // print the symbal aka charackter
			//if unicode.IsSymbol(sr) {fmt.Printf("%c  ", sr)}

			//Findings:
			//  ...  (57468 ... 58110) is nix of Is...
			// 𐎀 ...
			// ⠁ ... ⣿ Braile
			// 🅰 ... 🧀 most single point emoji (126000 ... 129490)
			/*
				fmt.Printf("\n--------------------------------------\n")
				fmt.Printf("ID: %v\n", i)
				fmt.Printf("The Codepoint: %+q\n", sr) // print back the input
				fmt.Printf("The Symbol: %c\n", sr)     // print the symbal aka charackter
				fmt.Printf("Unicode-Format: %U\n", sr) // print proper Formated Codepoint
			*/
		}

	} else {

		fmt.Printf("plain string: ")
		fmt.Printf("%s", arg)
		fmt.Printf("\n")

		fmt.Printf("quoted string: ")
		fmt.Printf("%+q", arg)
		fmt.Printf("\n")

		/*
			const placeOfInterest = `🍒`
			fmt.Printf("hex bytes: ")
			for i := 0; i < len(arg); i++ {
				fmt.Printf("%x ", arg[i])
			}
		*/

	}

	//fmt.Println(sr) // wrong way to print!

	/*
		src := '你'
		fmt.Printf("%q\n", src)  // print character 你
		fmt.Printf("%+q\n", src) // print unicode codepoint
	*/

}
