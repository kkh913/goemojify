package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cobra"
	"goemojify/emojidb"
	"os"
)

var gitTag string

type EmojiDataBase struct {
	Emoji          string   `json:"emoji"`
	Description    string   `json:"description"`
	Category       string   `json:"category"`
	Aliases        []string `json:"aliases"`
	Tags           []string `json:"tags"`
	UnicodeVersion string   `json:"unicode_version"`
	IOSVersion     string   `json:"ios_version"`
}

var emojiData []EmojiDataBase
var emojiAliases string

var rootCmd = &cobra.Command{
	Use:   "goemojify",
	Short: "Convert the aliases to emojy raw characters",
	Long: `Golang port of the original emojify(https://github.com/mrowa44/emojify).
all the glories should belong to mrowa44(https://github.com/mrowa44).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			emojiAliases = args[0]
		}
		return runCommand()
	},
}

var flags struct {
	filepath string
	version  bool
	list     bool
}

var flagsName = struct {
	file, fileShort       string
	version, versionShort string
	list, listShort       string
}{
	"file", "f",
	"version", "v",
	"list", "l",
}

func main() {
	rootCmd.Flags().StringVarP(
		&flags.filepath,
		flagsName.file,
		flagsName.fileShort,
		"", "path to the file")
	rootCmd.PersistentFlags().BoolVarP(
		&flags.version,
		flagsName.version,
		flagsName.versionShort,
		false, "version number")
	rootCmd.PersistentFlags().BoolVarP(
		&flags.list,
		flagsName.list,
		flagsName.listShort,
		false, "show the list of supported emojies")

	bytejson, _ := emojidb.Asset("emoji.json")

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if err := json.Unmarshal(bytejson, &emojiData); err != nil {
		panic(err)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
