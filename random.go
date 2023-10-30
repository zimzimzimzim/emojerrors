package emojerrors

import (
	"math/rand"
	"time"
)

// NewRandom returns an error with a semi-random emoji
// The possible emojis were chosen by yours truly, and you will accept them.
// All random emojis are 100% gluten free.
func NewRandom(text string) error {
	return New(text, getRandomEmoji())
}

func getRandomEmoji() string {
	return possibleRandomEmojis[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(possibleRandomEmojis))]
}

var possibleRandomEmojis = []string{
	"\U0001f966",
	"\U0001f95d",
	"\U0001f341",
	"\U0001f41e",
	"\U0001f438",
	"\U0001f427",
	"\U0001f43f\ufe0f",
	"\U0001f993",
	"\U0001f415",
	"\U0001f408",
	"\U0001f420",
	"\U0001f33d",
	"\U0001f954",
	"\U0001f35f",
	"\U0001f37f",
	"\U0001f363",
	"\U0001f9cb",
	"\U0001f944",
	"\U0001fab5",
	"\U0001f3e9",
	"\U0001f684",
}
