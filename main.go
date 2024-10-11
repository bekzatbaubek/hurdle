package main

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	os.Setenv("CANDIDATES", "HELLO,WORLD,QUITE,FANCY,FRESH,PANIC,CRAZY,BUGGY,SCARE")
	os.Setenv("NUMBER_OF_ROUNDS", "5")
}

func pickAnswer(candidates []string) string {
	index := rand.Intn(len(candidates))
	return candidates[index]
}

func hurdle(answer string, guess string) string {
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

func main() {
	candidate_env := os.Getenv("CANDIDATES")
	candidates := strings.Split(candidate_env, ",")

	numberOfRoundsEnv := os.Getenv("NUMBER_OF_ROUNDS")
	numberOfRounds, err := strconv.Atoi(numberOfRoundsEnv)
	if err != nil {
		// Configuration error
		panic(err)
	}

	answer := pickAnswer(candidates)

	fmt.Println("Wellcome to Hurdle")
	fmt.Printf("You have %d rounds to guess the word\n", numberOfRounds)
	fmt.Println("Hints:")
	fmt.Println("'0' - letter not in word")
	fmt.Println("'1' - letter in word but in different position")
	fmt.Println("'2' - letter in word and in correct position")

	// Game loop
	for i := 0; i < numberOfRounds; i++ {
		var guess string
		fmt.Printf("Round %d\n", i+1)
		fmt.Printf("Enter your guess:\n")
		fmt.Scanln(&guess)

		// Validate guess
		guess = strings.ToUpper(guess)
		if len(guess) != len(answer) || regexp.MustCompile(`[^A-Z]`).MatchString(guess) {
			fmt.Println("Invalid guess format")
			i--
			continue
		}
		// Hurdle the guess
		hint := hurdle(answer, guess)
		fmt.Println(hint)
		if hint == "22222" {
			fmt.Println("You win!")
			return
		}
	}
	fmt.Println("You lose!")
}
