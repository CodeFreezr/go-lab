package main

type EmojiMeta []struct {
	No          int    `json:"No"`
	Emoji       string `json:"Emoji"`
	Category    string `json:"Category"`
	SubCategory string `json:"SubCategory"`
	Unicode     string `json:"Unicode"`
	Name        string `json:"Name"`
	Tags        string `json:"Tags"`
	Shortcode   string `json:"Shortcode"`
}

type Person struct {
	Name string
	Age  int
}

people := []Person{
	{"Bob", 31},
	{"John", 42},
	{"Michael", 17},
	{"Jenny", 26},
}

