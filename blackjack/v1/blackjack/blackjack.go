package blackjack

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var Deck = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var CardValues = map[string]int{
	"a":  1,
	"A":  11,
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  10,
	"Q":  10,
	"K":  10,
}

type Player struct {
	Name      string
	Cards     []string
	HandValue int
}

func sumHand(player *Player) {
	player.HandValue = 0
	for _, card := range player.Cards {
		player.HandValue = player.HandValue + CardValues[card]
	}
	if player.HandValue > 21 {
		if updateAce(&player.Cards) {
			sumHand(player)
		}
	}
}

func updateAce(cards *[]string) bool {
	for i := range *cards {
		if (*cards)[i] == "A" {
			(*cards)[i] = "a"
			return true
		}
	}
	return false
}

func DealCard(player *Player) {
	rand.Seed(time.Now().UnixNano())
	card := Deck[rand.Intn(len(Deck))]
	player.Cards = append(player.Cards, card)
	sumHand(player)
}

func DealHands(players *[]Player, CardsToDeal int) {

	for i := 0; i < CardsToDeal; i++ {
		for n := range *players {
			DealCard(&(*players)[n])
		}
	}
}

func Play(players *[]Player) {
	DealHands(players, 2)
	gameOver := false

	var choice string
	for {
		fmt.Printf("CPU Display Card: %v\n", (*players)[1].Cards[1])
		fmt.Printf("Your Hand (%d): %s\n", (*players)[0].HandValue, strings.Join((*players)[0].Cards, ","))

		fmt.Printf("[H]it or [S]tand: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Printf("Error reading input: %v.  Please try again.\n", err)
			continue
		}
		if strings.ToUpper(choice) == "H" {
			DealCard(&(*players)[0])
		} else if strings.ToUpper(choice) == "S" {
			break
		} else {
			fmt.Printf("(%v) is not a valid selection, please try again\n\n", choice)
			continue
		}
		if (*players)[0].HandValue > 21 {
			gameOver = true
			break
		}
	}
	if !gameOver {
		for {
			if (*players)[1].HandValue >= 17 {
				break
			}
			DealCard(&(*players)[1])
		}
	}

	fmt.Printf("CPU Hand (%d): %s\n", (*players)[1].HandValue, strings.Join((*players)[1].Cards, ","))
	fmt.Printf("Your Hand (%d): %s\n", (*players)[0].HandValue, strings.Join((*players)[0].Cards, ","))

	if (*players)[0].HandValue > 21 {
		fmt.Println("You busted YOU LOSE")
	} else if (*players)[1].HandValue > 21 {
		fmt.Println("CPU busted YOU WIN!!!")
	} else if (*players)[1].HandValue == (*players)[0].HandValue {
		fmt.Println("It's a Draw")
	} else if (*players)[0].HandValue > (*players)[1].HandValue {
		fmt.Println("You Win")
	} else {
		fmt.Println("You Lose")
	}

}
