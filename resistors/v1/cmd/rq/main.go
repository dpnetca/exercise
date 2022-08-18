package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dpnetca/exercise/resistors/resistor"
	resistorquiz "github.com/dpnetca/exercise/resistors/resistorQuiz"
)

func main() {
	var numQuestions, numBands int
	var mode string

	flag.IntVar(&numQuestions, "n", 5, "Number of questions to generate. Default: 5")
	flag.IntVar(&numBands, "b", 3, "Number of bands per resistor (3-6, or 0 to randomize). Default: 3")
	flag.StringVar(&mode, "m", "text", "Band Display mode (text, color, graphic) Default: text")
	flag.Parse()

	// TODO validate mode
	if numBands != 0 && (numBands < 3 || numBands > 6) {
		log.Fatalf("Invalid number bands, '%d' must be 0,3,4,5,6", numBands)
	}

	quiz, err := resistorquiz.NewQuiz(numQuestions, numBands)
	if err != nil {
		log.Fatalf("Error generating quiz: %v", err)
	}

	// TODO seperate quiz presentation into seperate function, perhaps as part of package
	// TODO clear screen before quiz and between questions (seeing previous questions is easy to cheat)
	var answer string
	for i, r := range quiz.Resistors {
		fmt.Printf("%d) ", i+1)
		if mode != "text" {
			fmt.Print(resistor.BG + " ")
		}
		for _, b := range r.Bands {
			if mode == "color" {
				fmt.Printf("%s%s ", b.Ansi, b.Name)
			} else if mode == "graphic" {
				fmt.Printf("%s┃", b.Ansi)
			} else {
				fmt.Printf("%s ", b.Name)
			}
		}
		fmt.Println(" \033[0m")
		fmt.Printf("Resistance: ")
		fmt.Scanln(&answer)

		correct, err := quiz.ValidateAnswer(i, answer)
		if err != nil {
			fmt.Printf("Wrong, Error validating Answer: %v", err)
		}
		if correct {
			fmt.Println("Correct")
		} else {
			fmt.Printf("Wrong, answer is: %.0f (%s)\n", r.Resistance, resistor.ToString(r.Resistance))
		}
	}
	fmt.Printf("Final Score: %d / %d\n", quiz.Correct, len(quiz.Resistors))

}
