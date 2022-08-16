package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	resistorquiz "github.com/dpnetca/exercise/resistors/resistorQuiz"
)

func main() {
	var numQuestions int

	flag.IntVar(&numQuestions, "n", 5, "Number of questions to generate. Default: 5")

	flag.Parse()

	quiz, err := resistorquiz.NewQuiz(numQuestions)
	if err != nil {
		log.Fatalf("Error generating quiz: %v", err)
	}

	// TODO seperate quiz presentation into seperate function, perhaps as part of package
	var answer string
	for i, r := range quiz.Resistors {
		fmt.Printf("%d) ", i+1)
		for _, b := range r.Bands {
			fmt.Printf("%s ", b.Name)
		}
		fmt.Printf("\nResistance: ")
		fmt.Scanln(&answer)

		// TODO allow answer to use (k or K),M,G  i.e. 1000000 = 1000K = 1M
		answerInt, _ := strconv.Atoi(answer)
		if answerInt == r.Resistance {
			fmt.Println("Correct")
			quiz.Correct++
		} else {
			fmt.Printf("Wrong, answer is: %v\n", r.Resistance)
		}
	}
	fmt.Printf("Final Score: %d / %d\n", quiz.Correct, len(quiz.Resistors))

}
