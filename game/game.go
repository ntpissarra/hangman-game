package game

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type WordsContainer struct {
	Words []string `json:"words"`
}

var game = []string{
	`
 ____
|    |
|    o   
|   /|\   
|    |   
|   / \
|
---------
`,
	` 
 ____
|    |
|    o   
|   /|\   
|    |   
|   
|
---------
`,
	`
 ____
|    |
|    o   
|   /|\   
|       
|   
|
---------
`,
	`
 ____
|    |
|    o   
|    |   
|       
|   
|
---------
`,
	` 
 ____
|    |
|    o   
|     
|      
|   
|
---------
`,
	`
 ____
|    |
|      
|     
|      
|   
|
---------
`,
}

var secret string = "gato"

func ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func PrintIntro() {
	fmt.Println(`
Bem vindo da hangman
--------------------

Objetivo: Adivinhar uma palavra secreta escolhida pelo sistema.

Regras:
* O sistema escolhe uma palavra secreta de um banco de palavras.
* O sistema desenha um esboço da forca na tela do console.
* O jogador deve adivinhar uma letra de cada vez ou a palavra inteira.
* Se a letra estiver na palavra secreta, o sistema desenha a letra.
* Se a letra não estiver na palavra secreta, o sistema desenha uma parte do corpo 
  do homem no esboço da forca.
* O jogo termina quando o jogador adivinha a palavra secreta ou quando o desenho 
  da forca estiver completo.

Vencedor:
* O jogador vence se adivinhar a palavra secreta antes do desenho da 
forca estar completo.
* O sistema vence se o desenho da forca estiver completo antes do jogador 
adivinhar a palavra secreta.`)
}

func FirstQuestion() {
	fmt.Print(`
Para iniciar o jogo digite Y e para sair digite N: `)
}

func StartGame(word string) {
	lives := len(game) - 1
	secretWord := word
	hiddenWord := strings.Repeat("*", len(secretWord))
	var history []string

	for {
		ClearScreen()
		fmt.Printf(`
%d Vidas restantes
%s
%s   
%s
`, lives, game[lives], history, hiddenWord)

		var guess string

		if lives != 0 {
			fmt.Scan(&guess)
			if contains(history, guess) {
				continue
			}
			if guess == secretWord {
				hiddenWord = secretWord
				fmt.Println()
				fmt.Println("Parabéns você acertou a palavra secreta!")
				break
			} else if len(guess) == 1 {
				if !contains(history, guess) {
					history = append(history, guess)
				}
				hasLetter := false
				for i := 0; i < len(secretWord); i++ {
					if guess == string(secretWord[i]) {
						hiddenWord = hiddenWord[:i] + guess + hiddenWord[i+1:]
						hasLetter = true
					}
				}
				if !hasLetter {
					lives--
				}
				if !strings.Contains(hiddenWord, "*") {
					hiddenWord = secretWord
					ClearScreen()
					fmt.Printf(`
%s   
%s
`, game[lives], hiddenWord)
					fmt.Println()
					fmt.Println("Parabéns você acertou a palavra secreta!")
					break
				}

			} else {
				lives--
			}
		} else {
			fmt.Printf("Acabaram as chances, você perdeu! A palavra era %s", secretWord)
			break
		}
	}

}

func GetWord(w *string) {
	var container WordsContainer

	jsonFile, err := os.ReadFile("./words.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(jsonFile, &container)
	if err != nil {
		log.Fatal(err)
	}

	wordslen := len(container.Words) - 1
	*w = container.Words[rand.Intn(wordslen)]
}

func contains(a []string, p string) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == p {
			return true
		}
	}
	return false
}
