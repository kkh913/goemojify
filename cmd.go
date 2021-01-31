package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"regexp"
	"strings"
)

func runCommand() error {
	if flags.version {
		fmt.Println(gitTag)
		return nil
	} else if flags.list {
		for _, item := range emojiData {
			for _, v := range item.Aliases {
				fmt.Printf(":%s: %s\n", v, item.Emoji)
			}
		}
		return nil
	} else if isInputFromPipe() {
		return toEmoji(os.Stdin, os.Stdout)
	} else if emojiAliases != "" {
		return toEmoji(strings.NewReader(emojiAliases), os.Stdout)
	} else {
		file, e := getFile()
		if e != nil {
			return e
		}
		defer file.Close()
		return toEmoji(file, os.Stdout)
	}
}

func isInputFromPipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}

func getFile() (*os.File, error) {
	if flags.filepath == "" {
		return nil, errors.New("please input a file")
	}
	if !fileExists(flags.filepath) {
		return nil, errors.New("the file provided does not exist")
	}
	file, e := os.Open(flags.filepath)
	if e != nil {
		return nil, errors.Wrapf(e,
			"unable to read the file %s", flags.filepath)
	}
	return file, nil
}

func processConvert(s string) string {
	re := regexp.MustCompile(`\w+|([+-]\d)`)

	parts := strings.Split(s, ":")

	Contains := func(s []string, substr string) bool {
		for _, v := range s {
			if v == substr {
				return true
			}
		}
		return false
	}

	ToEmoji := func(source string) string {
		for _, item := range emojiData {
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

	return ret[1:]
}

func toEmoji(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		_, e := fmt.Fprintln(
			w, processConvert(scanner.Text()))
		if e != nil {
			return e
		}
	}
	return nil
}

func fileExists(filepath string) bool {
	info, e := os.Stat(filepath)
	if os.IsNotExist(e) {
		return false
	}
	return !info.IsDir()
}
