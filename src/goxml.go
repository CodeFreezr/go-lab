package main

import (
	"fmt"
	"strings"
)

func main() {
	htmlstr := `<?xml version="1.0" ?>
    <html>
    <head>
     <title>this is a title</title>
    </head>
    <body>Hello,World</body>
    </html>`
	root, err := xmlquery.Parse(strings.NewReader(htmlstr))
	if err != nil {
		panic(err)
	}
	title := xmlquery.FindOne(root, "//title")
	fmt.Println(title.InnerText())
}
