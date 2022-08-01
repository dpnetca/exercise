package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// random must be seeded to generate a psuedo-random number, seeding with time ensures a different result every time program is run
	// however this is a predictable result and not security safe
	rand.Seed(time.Now().UnixNano())

	// print Title
	fmt.Printf("Rock Paper Scissors\n*******************\n\n")

	rock := "Rock"
	paper := "Paper"
	scissors := "Scissors"

	options := []string{rock, paper, scissors}

	cpuChoice := rand.Intn(len(options))

	var input string
	var userChoice int
	for {
		fmt.Print("What do you choose? 0 for Rock, 1 for Paper or 2 for Scissors: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalf("Error getting input: %v\n", err)
		}
		userChoice, err = strconv.Atoi(input)
		if err != nil {
			log.Fatalf("unable to convert input to integer selection: %v\n", err)
		}
		if userChoice >= 0 && userChoice <= 2 {
			break
		} else {
			fmt.Printf("%v is an invalid entry,value must be between 0 and 2, please select again\n", userChoice)
		}
	}
	fmt.Printf("%-10s %-10s\n", "CPU", "User")
	fmt.Printf("%-10s %-10s\n", options[cpuChoice], options[userChoice])

	var result string
	if userChoice == cpuChoice {
		result = "Draw"
	} else if userChoice == 0 && cpuChoice == 2 {
		result = "CPU Wins"
	} else if userChoice > cpuChoice {
		result = "User Wins"
	} else {
		result = "CPU Wins"
	}
	fmt.Printf("\n*********\nResult: %s\n*********\n", result)
}
