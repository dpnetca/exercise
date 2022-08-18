package resistorquiz

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/dpnetca/exercise/resistors/resistor"
)

// TODO Resistor should be part of resistor package but resistor.Resistor is bad format need to rename something...
type Resistor struct {
	Bands      []resistor.Color
	Resistance float64
}

type Quiz struct {
	Resistors []Resistor
	Correct   int
}

func NewQuiz(numQuestions, numBands int) (Quiz, error) {
	var q Quiz
	err := q.Generate(numQuestions, numBands)
	return q, err
}

// TODO should be part of resistor package...
// TODO add rules to ensure resitor is valid (i.e. no black on first band)
func (r *Resistor) Generate(numBands int) error {
	if (numBands < 3 || numBands > 6) && numBands != 0 {
		return fmt.Errorf("invalid number of bands (%d) must 0 or be between 3 and 6", numBands)
	}
	rand.Seed(time.Now().UnixNano())
	if numBands == 0 {
		numBands = rand.Intn(3) + 3
	}
	for i := 0; i < numBands; i++ {
		n := rand.Intn(len(resistor.Bands))
		r.Bands = append(r.Bands, *resistor.Bands[n])
	}
	return nil
}
func (q *Quiz) Generate(numQuestion, numBands int) error {

	q.Correct = 0
	var err error

	for i := 0; i < numQuestion; i++ {
		q.Resistors = append(q.Resistors, Resistor{})
		q.Resistors[i].Generate(numBands)
		//TODO validate resistor is unique, (don't quiz on the same resistor twice)
		q.Resistors[i].Resistance, err = resistor.Calculate(q.Resistors[i].Bands)
		if err != nil {
			return err
		}
	}
	return nil
}

func (q *Quiz) ValidateAnswer(rIndex int, answerStr string) (bool, error) {
	var answer float64
	var err error

	answerStr = strings.TrimSpace(answerStr)

	if unicode.IsDigit([]rune(answerStr[len(answerStr)-1:])[0]) {
		answer, err = strconv.ParseFloat(answerStr, 64)
		if err != nil {
			return false, err
		}
	} else {
		lastChar := answerStr[len(answerStr)-1:]
		answerStr = answerStr[:len(answerStr)-1]
		answer, err = strconv.ParseFloat(answerStr, 64)
		if err != nil {
			return false, err
		}

		switch lastChar {
		case "k", "K":
			answer = answer * 1000

		case "M":
			answer = answer * 1000000

		case "G":
			answer = answer * 1000000000
		default:
			return false, fmt.Errorf("invalid answer '%v' not recognized", lastChar)
		}
	}
	if q.Resistors[rIndex].Resistance == float64(answer) {
		q.Correct++
		return true, nil
	}
	return false, nil
}
