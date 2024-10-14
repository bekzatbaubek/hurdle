package hurdle

import (
	"math/rand"
	"regexp"
)

func PickAnswer(candidates []string) string {
	index := rand.Intn(len(candidates))
	return candidates[index]
}

func InvalidGuess(guess string) bool {
	if len(guess) != 5 {
		return true
	}
	if regexp.MustCompile(`[^A-Z]`).MatchString(guess) {
		return true
	}
	return false
}

func Hurdle(answer string, guess string) string {
	occupiedPositions := make([]bool, len(answer))

	hint := make([]rune, len(guess))
	for i := 0; i < len(guess); i++ {
		hint[i] = '0'
	}

	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			hint[i] = '2'
			occupiedPositions[i] = true
		}
	}

	for i := 0; i < len(guess); i++ {
		if occupiedPositions[i] {
			continue
		}
		for j := 0; j < len(answer); j++ {
			if guess[i] == answer[j] {
				if occupiedPositions[j] {
					continue
				} else {
					hint[i] = '1'
					break
				}
			}
		}
	}

	return string(hint)
}
