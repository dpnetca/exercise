package main

import (
	"github.com/dpnetca/exercise/blackjack/v1/blackjack"
)

func main() {
	var players = []blackjack.Player{
		{Name: "Human"},
		{Name: "CPU"},
	}
	blackjack.Play(&players)

}
