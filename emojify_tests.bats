#!/usr/bin/env bats

@test "handles an input without a single emoji 😿" {
  result=$(./build/goemojify "no emoji :(")
  [ "$result" = "no emoji :(" ]
}

@test "handles an input with a single emoji 😹" {
  result=$(./build/goemojify "an emoji :grin:")
  [ "$result" = "an emoji 😁" ]
}

@test "handles an input with a lot of emojis 😻" {
  result=$(./build/goemojify "emojis :grin::grin: :tada:yay:champagne:")
  [ "$result" = "emojis 😁😁 🎉yay🍾" ]
}

@test "handles emojis with underscores and numbers 💯" {
  result=$(./build/goemojify "this is perfect :100: :1st_place_medal:")
  [ "$result" = "this is perfect 💯 🥇" ]
}

@test "handles emojis with + and - 👍" {
  result=$(./build/goemojify "great :+1::+1::-1:")
  [ "$result" = "great 👍👍👎" ]
}

@test "handles right-hand side emojis 👉" {
  result=$(./build/goemojify ":not_an_emoji:point_right:")
  [ "$result" = ":not_an_emoji👉" ]
  result=$(./build/goemojify "::::point_right:")
  [ "$result" = ":::👉" ]
}

@test "handles punctuations just after aliases" {
  result=$(./build/goemojify "Enter the :airplane:!")
  [ "$result" = "Enter the ✈️!" ]
}

@test "ignores existing unicode emoji characters" {
  result=$(./build/goemojify "🐛 leave the emojis alone!!")
  [ "$result" = "🐛 leave the emojis alone!!" ]
}

@test "handles multiple spaces after an emoji" {
  result=$(./build/goemojify ":sparkles:   Three spaces")
  [ "$result" = "✨   Three spaces" ]
  result=$(./build/goemojify ":sparkles:     Five spaces")
  [ "$result" = "✨     Five spaces" ]
  result=$(./build/goemojify ":sparkles: One space")
  [ "$result" = "✨ One space" ]
}

@test "handles the examples from the readme 😉" {
  result=$(./build/goemojify "Hey, I just :raising_hand: you, and this is :scream: , but here's my :calling: , so :telephone_receiver: me, maybe?")
  [ "$result" = "Hey, I just 🙋 you, and this is 😱 , but here's my 📲 , so 📞 me, maybe?" ]
  result=$(./build/goemojify "To :bee: , or not to :bee: : that is the question... To take :muscle: against a :ocean: of troubles, and by opposing, end them?")
  [ "$result" = "To 🐝 , or not to 🐝 : that is the question... To take 💪 against a 🌊 of troubles, and by opposing, end them?" ]
}

@test "handles the list option" {
  emojis=$(./build/goemojify -l | grep "2nd_place_medal")
  [ "$emojis" == ":2nd_place_medal: 🥈" ]
  emojis=$(./build/goemojify --list | grep "2nd_place_medal")
  [ "$emojis" == ":2nd_place_medal: 🥈" ]
}

@test "handles the version option" {
  version=$(./build/goemojify -v)
  [[ "$version" =~ ^[0-9]+.[0-9]+.[0-9]+$ ]]
  version=$(./build/goemojify --version)
  [[ "$version" =~ ^[0-9]+.[0-9]+.[0-9]+$ ]]
}
