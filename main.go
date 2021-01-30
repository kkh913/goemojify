package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	flag "github.com/spf13/pflag"
	"goemojify/emojidb"
	"os"
	"regexp"
	"strings"
)

var GitTag string

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

	showList := flag.BoolP("list", "l", false, "show the list of supported emojies")
	showVersion := flag.BoolP("version", "v", false, "version")
	flag.Parse()

	if *showVersion {
		fmt.Println(GitTag)
		return
	}

	re := regexp.MustCompile(`\w+|([+-]\d)`)

	bytejson, _ := emojidb.Asset("emoji.json")

	var emoji_data []EmojiDataBase

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.Unmarshal(bytejson, &emoji_data); err != nil {
		panic(err)
	}

	if *showList {
		for _, item := range emoji_data {
			for _, v := range item.Aliases {
				fmt.Printf(":%s: %s\n", v, item.Emoji)
			}
		}
		return
	}

	emoji_sentence := args[1]
	parts := strings.Split(emoji_sentence, ":")

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

			if Contains(item.Aliases, source) {
				return item.Emoji
			}
		}

		return source
	}

	var ret string
	prev_state := false

	for _, part := range parts {
		if len(strings.Fields(part)) == 1 && re.MatchString(part) {
			toEmoji := ToEmoji(part)
			if toEmoji != part {
				ret = ret + toEmoji
				prev_state = true
			} else {
				if prev_state {
					ret = ret + part
				} else {
					ret = ret + ":" + part
				}
				prev_state = false
			}
		} else {
			if prev_state {
				ret = ret + part
			} else {
				ret = ret + ":" + part
			}
			prev_state = false
		}
	}

	fmt.Println(ret[1:])
}
