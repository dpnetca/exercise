package resistor

import "fmt"

const ohm = 8486

type Color struct {
	Name       string
	Ansi       string
	Value      int
	Multiplier int
}

// TODO Add resistor ansi colours codes
var Black = Color{
	Name:       "Black",
	Ansi:       "",
	Value:      0,
	Multiplier: 1,
}
var Brown = Color{
	Name:       "Brown",
	Ansi:       "",
	Value:      1,
	Multiplier: 10,
}
var Red = Color{
	Name:       "Red",
	Ansi:       "",
	Value:      2,
	Multiplier: 100,
}
var Orange = Color{
	Name:       "Orange",
	Ansi:       "",
	Value:      3,
	Multiplier: 1000,
}
var Yellow = Color{
	Name:       "Yellow",
	Ansi:       "",
	Value:      4,
	Multiplier: 10000,
}
var Green = Color{
	Name:       "Green",
	Ansi:       "",
	Value:      5,
	Multiplier: 100000,
}
var Blue = Color{
	Name:       "Blue",
	Ansi:       "",
	Value:      6,
	Multiplier: 1000000,
}
var Violet = Color{
	Name:       "Violet",
	Ansi:       "",
	Value:      7,
	Multiplier: 10000000,
}
var Grey = Color{
	Name:       "Grey",
	Ansi:       "",
	Value:      8,
	Multiplier: 100000000,
}
var White = Color{
	Name:       "White",
	Ansi:       "",
	Value:      9,
	Multiplier: 1000000000,
}

var Bands = []*Color{&Black, &Brown, &Red, &Orange, &Yellow, &Green, &Blue, &Violet, &Grey, &White}

func Calculate(bands []Color) (int, error) {
	var resistance int
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

func ToString(resistance int) string {

	var res float32
	var unit string
	if resistance < 1000 {
		res = float32(resistance)
		unit = ""
	} else if resistance < 1000000 {
		res = float32(resistance) / 1000
		unit = "k"
	} else if resistance < 1000000000 {
		res = float32(resistance) / 1000000
		unit = "M"
	} else {
		res = float32(resistance) / 1000000000
		unit = "G"
	}
	if res == float32(int32(res)) {
		return fmt.Sprintf("%.0f %s%c", res, unit, ohm)
	}
	return fmt.Sprintf("%.1f %s%c", res, unit, ohm)

}
