package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xpath"
	"github.com/antchfx/xquery/html"
)

func main() {
	//define inline XML
	htmlstr := `<?xml version="1.0" ?>
    <html>
    <head>
     <title>this is a title</title>
    </head>
    <body>Hello,World</body>
    </html>`

	root, err := htmlquery.Parse(strings.NewReader(htmlstr))
	if err != nil {
		panic(err)
	}

	exprA := xpath.MustCompile("//title/text()")
	iterA := exprA.Evaluate(htmlquery.CreateXPathNavigator(root)).(*xpath.NodeIterator)
	for iterA.MoveNext() {
		fmt.Printf("%s \n", iterA.Current().Value()) // output href
	}

	htmlcat := `<html>
<head>
    <title>Test-Title</title>
</head>
<body>
    Hello World!
    <a href=""></a>
<!--    <div id="abc" class="article"><a href="http://www.web.de">web.de</a></div>
    <div id="def" class="article"><a href="http://www.github.com">git</a></div>
    <div id="ghi" class="article"><a href="http://www.twitter.com">twitter</a></div> -->
    <maincategory catname="Smileys-People">😀  😁  😂  😃  😄  😅  😆  😋  😘  😗  😙  😚  🤗  🤨  😑  😶  🙄  😮  😯  😛  🤤  😓  😕  😲  ☹</maincategory>
    <maincategory catname="Animals-Nature">🦍&nbsp;🐶&nbsp;🐕&nbsp;🦊&nbsp;🐱&nbsp;🐈🦁🐆🐴🐎🦌🐮🐄</maincategory>
    <maincategory catname="Food-Drink">🍇🍋🍌🍏🍒🥝🥥🥑🍆🥕🌽🌶🥒🥦🌰🍞🥐🥖🧀🥩🥓🍔</maincategory>
	<maincategory catname="Activity">🎃🎄🎆🎈🎊🎎🎏🎟🥇🥈🥉⚾🏀</maincategory>
	<maincategory catname="Activity">😶&nbsp;🙄&nbsp;😮&nbsp;😯&nbsp;😛&nbsp;🤤&nbsp;😓&nbsp;😕&nbsp;😲</maincategory>
    <flunder class="main">Travel & Places</flunder>
    <flunder class="main">Objects</flunder>
    <flunder class="main">Symbols</flunder>
    <flunder class="main">Flags</flunder>
</body>
</html>`

	doc, err := htmlquery.Parse(strings.NewReader(htmlcat))
	if err != nil {
		panic(err)
	}

	/*
		//Now Per File
		// Load HTML file.
		f, err := os.Open(`test.html`)
		if err != nil {
			panic(err)
		}
		// Parse HTML document.
		doc, err := htmlquery.Parse(f)
		if err != nil {
			panic(err)
		}
	*/

	// Option 1: using xpath's expr to matches nodes.
	//	expr := xpath.MustCompile("count(//div[@class='article'])")
	//	fmt.Printf("%f \n", expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(float64))

	expr := xpath.MustCompile("//a/@href")
	iter := expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(*xpath.NodeIterator)
	for iter.MoveNext() {
		fmt.Printf("%s \n", iter.Current().Value()) // output href
	}

	// Option 2: using build-in functions Find() to matches nodes.
	for _, n := range htmlquery.Find(doc, "//a/@href") {
		fmt.Printf("%s \n", htmlquery.SelectAttr(n, "href")) // output href
	}

	expr = xpath.MustCompile("//maincategory/@catname")
	iter = expr.Evaluate(htmlquery.CreateXPathNavigator(doc)).(*xpath.NodeIterator)
	for iter.MoveNext() {
		mainCateg := iter.Current().Value()
		fmt.Printf("%s: ", mainCateg) // output href

		mypath := "//maincategory[@catname='" + mainCateg + "']/text()"

		fmt.Printf("%s: \n", mypath)
		//xPrint ("",doc)
		exprInner := xpath.MustCompile(mypath)

		iterInner := exprInner.Evaluate(htmlquery.CreateXPathNavigator(doc)).(*xpath.NodeIterator)
		for iterInner.MoveNext() {
			fmt.Printf("%s: \n", iterInner.Current().Value())
		}

	}
}

func xPrint() {}
