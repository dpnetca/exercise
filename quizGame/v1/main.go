package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func getQuiz(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	quiz, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV file: %v", err)
	}
	return quiz, nil
}

func deliverQuestion(question []string) (bool, error) {
	var answer string
	fmt.Printf("Q: %s? ", question[0])
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false, fmt.Errorf("error getting answer: %v", err)
	}
	if strings.EqualFold(answer, question[1]) {
		return true, nil
	}
	return false, nil

}

func main() {
	var filename string
	var shuffle bool

	flag.StringVar(&filename, "file", "problems.csv", "Problems CSV file, Default: problems.csv")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle questions, Default:false")
	flag.Parse()

	quiz, err := getQuiz(filename)
	if err != nil {
		log.Fatal(err)
	}

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(quiz), func(i, j int) {
			quiz[i], quiz[j] = quiz[j], quiz[i]
		})
	}

	var score int = 0
	for _, question := range quiz {
		correct, err := deliverQuestion(question)
		if err != nil {
			log.Fatal(err)
		}
		if correct {
			score++
		}
	}
	fmt.Println("Quiz Complete")
	fmt.Printf("Final Score: %d/%d\n", score, len(quiz))

}
