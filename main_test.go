package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func HandleInputWithoutSingleEmoji(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"no emoji :(", "no emoji :("},
	}

	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandleInputWithSingleEmoji(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"an emoji :grin:", "an emoji 😁"},
	}

	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandleInputWithLotOfEmojis(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"emojis :grin::grin: :tada:yay:champagne:", "emojis 😁😁 🎉yay🍾"},
	}

	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandleEmojisWithUnderScoresAndNumbers(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"this is perfect :100: :1st_place_medal:", "this is perfect 💯 🥇"},
	}

	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandleEmojisWithPlusMinus(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"great :+1::+1::-1:", "great 👍👍👎"},
	}

	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandleRightHandSideEmojis(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{":not_an_emoji:point_right:", ":not_an_emoji👉"},
		{"::::point_right:", ":::👉"},
	}

	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandlePunctuationsJustAfterAliases(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"Enter the :airplane:!", "Enter the ✈️!"},
	}
	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func IgnoreExistingUnicodeEmoji(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"🐛 leave the emojis alone!!", "🐛 leave the emojis alone!!"},
	}

	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandleMultipleSpacesAfterEmoji(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{":sparkles:   Three spaces", "✨   Three spaces"},
		{":sparkles:     Five spaces", "✨     Five spaces"},
		{":sparkles: One space", "✨ One space"},
	}
	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func HandleExampleReadme(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"Hey, I just :raising_hand: you, and this is :scream: , but here's my :calling: , so :telephone_receiver: me, maybe?", "Hey, I just 🙋 you, and this is 😱 , but here's my 📲 , so 📞 me, maybe?"},
		{"To :bee: , or not to :bee: : that is the question... To take :muscle: against a :ocean: of troubles, and by opposing, end them?", "To 🐝 , or not to 🐝 : that is the question... To take 💪 against a 🌊 of troubles, and by opposing, end them?"},
	}
	for _, tc := range testCases {
		assert.Equal(t, processConvert(tc.input), tc.output)
	}
}

func TestGoemojifyTestSuite(t *testing.T) {

	GetEmojiDB()

	t.Run("handles an input without a single emoji 😿", HandleInputWithoutSingleEmoji)
	t.Run("handles an input with a single emoji 😹", HandleInputWithSingleEmoji)
	t.Run("handles an input with a lot of emojis 😻", HandleInputWithLotOfEmojis)
	t.Run("handles emojis with underscores and numbers 💯", HandleEmojisWithUnderScoresAndNumbers)
	t.Run("handles emojis with + and - 👍", HandleEmojisWithPlusMinus)
	t.Run("handles right-hand side emojis 👉", HandleRightHandSideEmojis)
	t.Run("handles punctuations just after aliases", HandlePunctuationsJustAfterAliases)
	t.Run("ignores existing unicode emoji characters", IgnoreExistingUnicodeEmoji)
	t.Run("handles multiple spaces after an emoji", HandleMultipleSpacesAfterEmoji)
	t.Run("handles the examples from the readme 😉", HandleExampleReadme)
}
