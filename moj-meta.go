package main

import (
	"fmt"
)

type EmojiMeta struct {
	No          int
	Emoji       string
	Category    string
	SubCategory string
	Unicode     string
	Name        string
	Tags        string
	Shortcode   string
}

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Hello World")
	/*
		people := []Person{
			{"Bob", 31},
			{"John", 42},
			{"Michael", 17},
			{"Jenny", 26},
		}
		fmt.Println(people)
	*/
	EmojiFull := []EmojiMeta{
		{1, "😀", "Smiley & Peoples", "face-positive", "U+1F600", "grinning face", "face | grin | grinning face", ":grinning:"},
		{2, "😁", "Smiley & Peoples", "face-positive", "U+1F601", "beaming face with smiling eyes", "beaming face with smiling eyes | eye | face | grin | smile", ":grin:"},
		{3, "😂", "Smiley & Peoples", "face-positive", "U+1F602", "face with tears of joy", "face | face with tears of joy | joy | laugh | tear", ":joy:"},
		{4, "🤣", "Smiley & Peoples", "face-positive", "U+1F923", "rolling on the floor laughing", "face | floor | laugh | rolling | rolling on the floor laughing", ":rofl:"},
		{5, "😃", "Smiley & Peoples", "face-positive", "U+1F603", "grinning face with big eyes", "face | grinning face with big eyes | mouth | open | smile", ":smiley:"},
		{6, "😄", "Smiley & Peoples", "face-positive", "U+1F604", "grinning face with smiling eyes", "eye | face | grinning face with smiling eyes | mouth | open | smile", ":grin:"},
		{7, "😅", "Smiley & Peoples", "face-positive", "U+1F605", "grinning face with sweat", "cold | face | grinning face with sweat | open | smile | sweat", ":sweat_smile:"},
		{8, "😆", "Smiley & Peoples", "face-positive", "U+1F606", "grinning squinting face", "face | grinning squinting face | laugh | mouth | open | satisfied | smile", ":laughing:"},
		{9, "😉", "Smiley & Peoples", "face-positive", "U+1F609", "winking face", "face | wink | winking face", ":wink:"},
		{10, "😊", "Smiley & Peoples", "face-positive", "U+1F60A", "smiling face with smiling eyes", "blush | eye | face | smile | smiling face with smiling eyes", ":blush:"},
		{2622, "🏴󠁧󠁢󠁳󠁣󠁴󠁿", "flags", "subdivision-flag", "U+1F3F4 U+E0067 U+E0062 U+E0073 U+E0063 U+E0074 U+E007F", "Scotland", "flag", ":scotland:"},
		{2623, "🏴󠁧󠁢󠁷󠁬󠁳󠁿", "flags", "subdivision-flag", "U+1F3F4 U+E0067 U+E0062 U+E0077 U+E006C U+E0073 U+E007F", "Wales", "flag", ":wales:"}}

	//fmt.Println(EmojiFull)

	var s []EmojiMeta = EmojiFull[1:4]
	fmt.Println(len(s))

	fmt.Println(s[0].Shortcode)
}
