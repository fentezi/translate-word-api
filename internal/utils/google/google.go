package google

import (
	"fmt"

	gtranslate "github.com/gilang-as/google-translate"
)

func Translate(word string) (string, error) {
	fmt.Println(word)
	value := gtranslate.Translate{
		Text: word,
		From: "en",
		To:   "ru",
	}

	translated, err := gtranslate.Translator(value)
	if err != nil {
		return "", err
	}

	return translated.Text, nil
}
