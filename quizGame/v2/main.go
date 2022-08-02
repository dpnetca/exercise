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

type problem struct {
	question string
	answer   string
}

func getQuiz(filename string) ([]problem, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("unable to read CSV file: %v", err)
	}

	quiz := make([]problem, len(data))
	for i, line := range data {
		quiz[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return quiz, nil
}

func deliverQuestion(quiz problem) (bool, error) {
	var answer string
	fmt.Printf("Q: %s? ", quiz.question)
	_, err := fmt.Scanln(&answer)
	if err != nil {
		return false, fmt.Errorf("error getting answer: %v", err)
	}
	if strings.EqualFold(answer, quiz.answer) {
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
