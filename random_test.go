package emojerrors

import (
	"testing"
)

// Just make sure that the seed is working
func TestGetRandomEmoji(t *testing.T) {
	var (
		timesToGenerate = 10
		emojisFound     = make(map[string]int, timesToGenerate)
	)

	for i := 0; i < timesToGenerate; i++ {
		emojisFound[getRandomEmoji()]++
	}

	if len(emojisFound) == 1 {
		t.Errorf("really unlikely result, expected close to %d, found 1", timesToGenerate)
	}
}
