package main

import (
	"fmt"
	"strings"

	"github.com/ntpissarra/hangman/utils"
)

func main() {
	var input string

	utils.ClearScreen()
	utils.PrintIntro()

	for {
		var word string

		utils.FirstQuestion()
		fmt.Scan(&input)
		input = strings.ToLower(input)

		if input == "y" {
			utils.GetWord(&word)
			utils.StartGame(word)
		} else if input == "n" {
			break
		} else {
			utils.ClearScreen()
		}
	}

}
