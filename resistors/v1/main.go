package main

import (
	"fmt"

	"github.com/dpnetca/exercise/resistors/resistor"
)

const bg = "\033[48;2;210;180;140m"
const bold = "\033[1m"
const reset = "\033[0m"

func main() {
	// test printing ansi to screen.works in WSL Terminal, but not in vscode terminal
	fmt.Print(bold + bg + " ")
	for _, b := range resistor.Bands {
		fmt.Print(b.Ansi + "┃")

	}
	fmt.Println(" " + reset)

}
