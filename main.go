package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var (
	inp = bufio.NewReader(os.Stdin)
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dictionarySlice, _ := loadDictionary()
	keyWord := getKeyword(dictionarySlice)
	hangmanState := 0
	guessedLetters := initializeGuessedLetterd(keyWord)
	fmt.Println(keyWord)
	for !wordGussedStatus(keyWord, guessedLetters) && !isHangmanComplete(hangmanState) {
		printCurrentState(keyWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input")
			continue
		}
		letter := rune(input[0])
		if isCorrect(keyWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}

	}
	fmt.Println("...Game over...")
	if wordGussedStatus(keyWord, guessedLetters) {
		fmt.Println("You won")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("you loose")
	} else {
		panic("Invalid state")
	}

}
func wordGussedStatus(keyWord string, guessedLetters map[rune]bool) bool {

	for _, ch := range keyWord {

		if !guessedLetters[unicode.ToLower(ch)] {

			return false
		}

	}
	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}
func isCorrect(keyWord string, letter rune) bool {
	return strings.ContainsRune(keyWord, letter)
}

func readInput() string {
	fmt.Print(">> ")
	text, err := inp.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(text)
}

func printHangmanImage(state int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("resources/s%d", state))
	if err != nil {
		panic(err)
	}
	return string(data)

}

func printCurrentState(keyWord string, guessedLetters map[rune]bool, hangManState int) {
	for _, ch := range keyWord {
		if ch == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(ch)] {
			fmt.Printf("%c", ch)
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println()
	fmt.Println(printHangmanImage(hangManState))
}

func initializeGuessedLetterd(keyWord string) map[rune]bool {
	guessedletters := map[rune]bool{}
	guessedletters[unicode.ToLower(rune(keyWord[0]))] = true
	guessedletters[unicode.ToLower(rune(keyWord[len(keyWord)-1]))] = true
	return guessedletters
}

func getKeyword(dictionarySlice []string) string {
	keyWord := dictionarySlice[rand.Intn(len(dictionarySlice))]
	return keyWord
}

func printHangman(keyWord string, triedLetters map[rune]bool) {

}

func loadDictionary() ([]string, error) {
	fileName := "resources/dictionary"

	fileBytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sliceData := strings.Split(string(fileBytes), "\n")
	return sliceData, nil
}
