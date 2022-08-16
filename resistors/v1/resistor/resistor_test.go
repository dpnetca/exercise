package resistor_test

import (
	"fmt"
	"testing"

	"github.com/dpnetca/exercise/resistors/resistor"
)

const ohm = 8486

func TestCalculate(t *testing.T) {
	tests := []struct {
		input []resistor.Color
		want  int
	}{
		{[]resistor.Color{resistor.Black, resistor.Brown, resistor.Black}, 1},
		{[]resistor.Color{resistor.Black, resistor.Red, resistor.Brown}, 20},
		{[]resistor.Color{resistor.Black, resistor.Orange, resistor.Red}, 300},
		{[]resistor.Color{resistor.Black, resistor.Yellow, resistor.Orange}, 4000},
		{[]resistor.Color{resistor.Black, resistor.Green, resistor.Yellow}, 50000},
		{[]resistor.Color{resistor.Black, resistor.Blue, resistor.Green}, 600000},
		{[]resistor.Color{resistor.Black, resistor.Violet, resistor.Blue}, 7000000},
		{[]resistor.Color{resistor.Black, resistor.Grey, resistor.Violet}, 80000000},
		{[]resistor.Color{resistor.Black, resistor.White, resistor.Grey}, 900000000},
		{[]resistor.Color{resistor.White, resistor.White, resistor.White}, 99000000000},
	}
	for i, tc := range tests {
		var bands string
		for _, band := range tc.input {
			bands = bands + band.Name + " "
		}
		bands = bands[:len(bands)-1]
		t.Run(fmt.Sprintf("Caclulate %d-(%s)", i, bands), func(t *testing.T) {
			got, _ := resistor.Calculate(tc.input)
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success!")
			}
		})
	}

}

func TestToString(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1, fmt.Sprintf("1 %c", ohm)},
		{20, fmt.Sprintf("20 %c", ohm)},
		{300, fmt.Sprintf("300 %c", ohm)},
		{4000, fmt.Sprintf("4 k%c", ohm)},
		{5600, fmt.Sprintf("5.6 k%c", ohm)},
		{70000, fmt.Sprintf("70 k%c", ohm)},
		{800000, fmt.Sprintf("800 k%c", ohm)},
		{9000000, fmt.Sprintf("9 M%c", ohm)},
		{1200000, fmt.Sprintf("1.2 M%c", ohm)},
		{30000000, fmt.Sprintf("30 M%c", ohm)},
		{400000000, fmt.Sprintf("400 M%c", ohm)},
		{5000000000, fmt.Sprintf("5 G%c", ohm)},
		{6700000000, fmt.Sprintf("6.7 G%c", ohm)},
		{80000000000, fmt.Sprintf("80 G%c", ohm)},
		{900000000000, fmt.Sprintf("900 G%c", ohm)},
	}
	for i, tc := range tests {
		t.Run(fmt.Sprintf("Caclulate %d-(%d)", i, tc.input), func(t *testing.T) {
			got := resistor.ToString(tc.input)
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success!")
			}
		})
	}

}
