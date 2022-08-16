package main

import (
	"fmt"
)

const bg = "\033[48;5;94m"
const bold = "\033[1m"
const reset = "\033[0m"
const red = "\033[38;5;9m"
const blue = "\033[38;5;21m"

func main() {
	// test printing ansi to screen.works in WSL Terminal, but not in vscode terminal
	fmt.Println(bold + bg + red + " |" + blue + "| " + reset)

}
