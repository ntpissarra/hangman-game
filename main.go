package main

import (
	"fmt"
	"strings"

	"github.com/ntpissarra/hangman/game"
)

func main() {
	var input string

	game.ClearScreen()
	game.PrintIntro()

	for {
		var word string

		game.FirstQuestion()
		fmt.Scan(&input)
		input = strings.ToLower(input)

		if input == "y" {
			game.GetWord(&word)
			game.StartGame(word)
		} else if input == "n" {
			break
		} else {
			game.ClearScreen()
		}
	}

}
