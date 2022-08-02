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

func deliverQuestion(quiz problem, answerCH chan string) {
	var answer string
	fmt.Printf("Q: %s? ", quiz.question)
	fmt.Scanln(&answer)
	answerCH <- answer
}

func runQuiz(quiz []problem, timeLimit int) int {
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	var score int = 0
	for _, question := range quiz {
		answerCH := make(chan string)
		go deliverQuestion(question, answerCH)

		select {
		case <-timer.C:
			fmt.Printf("\nTime Expired\n")
			return score
		case answer := <-answerCH:
			if strings.EqualFold(answer, question.answer) {
				score++
			}
		}
	}
	return score
}

func main() {
	var filename string
	var shuffle bool
	var timeLimit int

	flag.StringVar(&filename, "file", "problems.csv", "Problems CSV file, Default: problems.csv")
	flag.BoolVar(&shuffle, "shuffle", false, "Shuffle questions, Default:false")
	flag.IntVar(&timeLimit, "limit", 30, "time limit for the quiz in seconds, Default:30")

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

	score := runQuiz(quiz, timeLimit)

	fmt.Println("Quiz Complete")
	fmt.Printf("Final Score: %d/%d\n", score, len(quiz))

}
