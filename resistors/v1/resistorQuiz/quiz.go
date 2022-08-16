package resistorquiz

import (
	"math/rand"
	"time"

	"github.com/dpnetca/exercise/resistors/resistor"
)

type Resistor struct {
	Bands      []resistor.Color
	Resistance int
}

type Quiz struct {
	Resistors []Resistor
	Correct   int
}

func NewQuiz(num int) (Quiz, error) {
	var q Quiz
	err := q.Generate(num)
	return q, err
}

func (q *Quiz) Generate(num int) error {
	rand.Seed(time.Now().UnixNano())
	q.Correct = 0
	var err error

	for i := 0; i < num; i++ {
		q.Resistors = append(q.Resistors, Resistor{})
		// TODO seperate question generation to seperate function?
		// TODO add variable number of bands, tolerance, etc.
		for j := 0; j < 3; j++ {
			n := rand.Intn(len(resistor.Bands))
			q.Resistors[i].Bands = append(q.Resistors[i].Bands, *resistor.Bands[n])
			//TODO validate question is unique
		}
		q.Resistors[i].Resistance, err = resistor.Calculate(q.Resistors[i].Bands)
		if err != nil {
			return err
		}
	}
	return nil
}
