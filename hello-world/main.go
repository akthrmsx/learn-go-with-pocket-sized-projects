package main

import (
	"flag"
	"fmt"
)

type language string

var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}

func main() {
	var lang string

	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

func greet(lang language) string {
	result, ok := phrasebook[lang]

	if !ok {
		return fmt.Sprintf("unsupported language: %q", lang)
	}

	return result
}
