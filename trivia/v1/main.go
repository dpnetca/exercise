package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	ResponseCode int      `json:"response_code"`
	Triva        []Trivia `json:"results"`
}
type Trivia struct {
	Category         string   `json:"category"`
	Type             string   `json:"type"`
	Difficulty       string   `json:"difficulty"`
	Question         string   `json:"question"`
	CorrectAnswer    string   `json:"correct_answer"`
	IncorrectAnswers []string `json:"incorrect_answers"`
}

type AnswerKey struct {
	Answer  string
	Correct bool
}

func getTrivia(url string) ([]Trivia, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("unable to get URL (%v): %v", url, err)
	}
	defer res.Body.Close()

	var response Response
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error parsing json: %v", err)
	}

	return response.Triva, nil
}

func printQuestion(i int, question string) error {
	var q []byte
	q, err := base64.StdEncoding.DecodeString(question)
	if err != nil {
		return fmt.Errorf("unable to decode question: %v", err)
	}
	fmt.Printf("%d) %v\n", i+1, string(q))
	return nil
}

func decodeAnswers(correctAnswer string, incorrectAnswers []string) ([]AnswerKey, error) {
	var allAnswers []AnswerKey

	a, err := base64.StdEncoding.DecodeString(correctAnswer)
	if err != nil {
		return nil, err
	}
	allAnswers = append(allAnswers, AnswerKey{Answer: string(a), Correct: true})

	for _, answer := range incorrectAnswers {
		a, err := base64.StdEncoding.DecodeString(answer)
		if err != nil {
			return nil, err
		}
		allAnswers = append(allAnswers, AnswerKey{Answer: string(a), Correct: false})
	}
	return allAnswers, nil
}

func shuffleAnswers(answers []AnswerKey) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(answers), func(i, j int) {
		answers[i], answers[j] = answers[j], answers[i]
	})
}

func printAnswers(answers []AnswerKey) {
	key := 'A'
	for i, a := range answers {
		fmt.Printf("    %s) %s\n", string(key+rune(i)), a.Answer)

	}
}

func getAnswer(numAnswers int) int {
	var input string
	var answerIndex int
	fmt.Print("Answer: ")
	for {
		fmt.Scanln(&input)
		if len(input) != 1 {
			fmt.Printf("Invalid input \"%v\", please enter a single character: ", input)
			continue
		}
		input = strings.ToUpper(input)
		answerIndex = int(input[0] - 'A')
		if answerIndex < 0 || answerIndex > numAnswers-1 {
			fmt.Printf("\"%v\" is not a valid answer, please try again: ", input)
			continue
		}

		break
	}

	return answerIndex
}

func getAnswers(correctAnswer string, incorrectAnswers []string) (bool, error) {
	allAnswers, err := decodeAnswers(correctAnswer, incorrectAnswers)
	if err != nil {
		return false, err
	}
	shuffleAnswers(allAnswers)
	printAnswers(allAnswers)

	answerIndex := getAnswer(len(allAnswers))
	if allAnswers[answerIndex].Correct {
		return true, nil
	}

	return false, nil
}

func runTrivia(trivia []Trivia) (int, error) {
	correct := 0
	for i, t := range trivia {
		err := printQuestion(i, t.Question)
		if err != nil {
			return 0, err
		}
		c, err := getAnswers(t.CorrectAnswer, t.IncorrectAnswers)
		if err != nil {
			return 0, err
		}
		if c {
			correct++
		}
	}
	return correct, nil
}

func main() {
	url := "https://opentdb.com/api.php?amount=10&category=9&difficulty=easy&type=multiple&encode=base64"
	trivia, err := getTrivia(url)
	if err != nil {
		log.Fatalln(err)
	}
	correct, err := runTrivia(trivia)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("score: %v/%v\n", correct, len(trivia))

}
