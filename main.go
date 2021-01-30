package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"goemojify/emojidb"
	"os"
	"regexp"
	"strings"
)

type EmojiDataBase struct {
	Emoji          string   `json:"emoji"`
	Description    string   `json:"description"`
	Category       string   `json:"category"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version"`
	IOSVersion     string   `json:"ios_version"`
}

func main() {

	args := os.Args

	if len(args) == 1 {
		fmt.Fprintf(os.Stderr, "You must input one argument\n")
		os.Exit(1)
	} else if len(args) > 2 {
		fmt.Fprintf(os.Stderr, "Error, only one command line argument allowed\n")
		os.Exit(1)
	}

	bytejson, _ := emojidb.Asset("emoji.json")

	var emoji_data []EmojiDataBase

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.Unmarshal(bytejson, &emoji_data); err != nil {
		panic(err)
	}

	test_sentence := args[1]
	// test_sentence := "kanghee :beer:"

	re := regexp.MustCompile(`:(\w+|([+-]\d)):`)

	Contains := func(s []string, substr string) bool {
		for _, v := range s {
			if v == substr {
				return true
			}
		}
		return false
	}

	ToEmoji := func(source string) string {

		for _, item := range emoji_data {

			word := strings.ReplaceAll(source, ":", "")

			if Contains(item.Aliases, word) {
				return item.Emoji
			}
		}

		return source
	}

	fmt.Println(re.ReplaceAllStringFunc(test_sentence, ToEmoji))

	// output may be...
	// Hey, I just 🙋 you, and this is 😱 , but here's my 📲 , so 📞 me, maybe?
}
