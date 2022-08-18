package resistor

import "fmt"

const ohm = 8486
const BG = "\033[48;2;210;180;140m"

type Color struct {
	Name       string
	Ansi       string
	Value      float64
	Multiplier float64
}

//TODO Add Tolerance
//TODO Add Multiplyer for Gold, Silver and Pink ()
//TODO add temperature coefficient

var Black = Color{
	Name:       "Black",
	Ansi:       "\033[38;5;0m",
	Value:      0,
	Multiplier: 1,
}
var Brown = Color{
	Name:       "Brown",
	Ansi:       "\033[38;2;140;70;20m",
	Value:      1,
	Multiplier: 10,
}
var Red = Color{
	Name:       "Red",
	Ansi:       "\033[38;5;1m",
	Value:      2,
	Multiplier: 100,
}
var Orange = Color{
	Name:       "Orange",
	Ansi:       "\033[38;5;208m",
	Value:      3,
	Multiplier: 1000,
}
var Yellow = Color{
	Name:       "Yellow",
	Ansi:       "\033[38;5;226m",
	Value:      4,
	Multiplier: 10000,
}
var Green = Color{
	Name:       "Green",
	Ansi:       "\033[38;5;2m",
	Value:      5,
	Multiplier: 100000,
}
var Blue = Color{
	Name:       "Blue",
	Ansi:       "\033[38;5;12m",
	Value:      6,
	Multiplier: 1000000,
}
var Violet = Color{
	Name:       "Violet",
	Ansi:       "\033[38;5;5m",
	Value:      7,
	Multiplier: 10000000,
}
var Grey = Color{
	Name:       "Grey",
	Ansi:       "\033[38;5;8m",
	Value:      8,
	Multiplier: 100000000,
}
var White = Color{
	Name:       "White",
	Ansi:       "\033[38;5;15m",
	Value:      9,
	Multiplier: 1000000000,
}

var Bands = []*Color{&Black, &Brown, &Red, &Orange, &Yellow, &Green, &Blue, &Violet, &Grey, &White}

func Calculate(bands []Color) (float64, error) {
	var resistance float64
	switch len(bands) {
	case 3:
		resistance = (bands[0].Value*10 + bands[1].Value) * bands[2].Multiplier
	case 4:
		resistance = (bands[0].Value*10 + bands[1].Value) * bands[2].Multiplier
	case 5:
		resistance = (bands[0].Value*100 + bands[1].Value*10 + bands[2].Value) * bands[3].Multiplier
	case 6:
		resistance = (bands[0].Value*100 + bands[1].Value*10 + bands[2].Value) * bands[3].Multiplier
	default:
		return 0, fmt.Errorf("invalid number of bands(%v) must include 3-6 bands", len(bands))
	}
	return resistance, nil
}

func ToString(resistance float64) string {

	var res float64
	var unit string
	if resistance < 1000 {
		res = resistance
		unit = ""
	} else if resistance < 1000000 {
		res = resistance / 1000
		unit = "k"
	} else if resistance < 1000000000 {
		res = resistance / 1000000
		unit = "M"
	} else {
		res = resistance / 1000000000
		unit = "G"
	}
	if res == float64(int32(res)) {
		return fmt.Sprintf("%.0f %s%c", res, unit, ohm)
	}
	return fmt.Sprintf("%.1f %s%c", res, unit, ohm)

}
