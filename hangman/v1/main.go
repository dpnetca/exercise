package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/dpnetca/exercise/hangman/v1/words"
)

const (
	maxWrong = 5
)

func main() {
	rand.Seed(time.Now().UnixNano())

	word := words.List[rand.Intn(len(words.List))]

	var displayWord []string
	for range word {
		displayWord = append(displayWord, "_")
	}

	var guessedLetters = make(map[string]bool)
	var letter string
	var found int
	wrong := 0

	for {
		fmt.Printf("Guess: %s   (%d/%d)\n", strings.Join(displayWord, " "), wrong, maxWrong)

		letter = ""
		fmt.Print("Pick a letter: ")
		_, err := fmt.Scanln(&letter)
		if err != nil {
			log.Printf("Error getting input: %v\n", err)
			continue
		}
		if len(letter) != 1 {
			fmt.Printf("Invalid entry '%v', you must guess a single letter. Please try again.\n", letter)
			continue
		}

		_, exists := guessedLetters[letter]

		if exists {
			fmt.Printf("you have already guessed '%v' try again.\n", letter)
			continue
		}
		guessedLetters[letter] = true
		found = 0
		for i, v := range word {
			if letter == string(v) {
				found++
				displayWord[i] = string(v)
			}
		}

		if strings.Join(displayWord, "") == word {
			fmt.Printf("Congratulations, you win!! The word was '%v'\n", word)
			break
		}

		if found == 0 {
			wrong++
		}

		if wrong >= maxWrong {
			fmt.Printf("Sorry you lose, the answer was '%v'\n", word)
			break
		}
	}
	fmt.Printf("Game Over\n")
}
