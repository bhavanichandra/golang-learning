package interaction

import (
	"bufio"
	"fmt"
	"github.com/bhavanichandra/golang-learning/monsterslayer/actions"
	"os"
	"runtime"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(isSpecialRound bool) string {
	for {
		playerChoice, _ := getPlayerInput()
		if isSpecialRound && playerChoice == "3" {
			return actions.SPECIAL_ATTACK_ACTION
		} else if playerChoice == "1" {
			return actions.ATTACK_ACTION
		} else if playerChoice == "2" {
			return actions.HEAL_ACTION
		}
		fmt.Println("Fetching the user input failed. Please try again!")
	}
}

// getPlayerInput
func getPlayerInput() (string, error) {
	fmt.Print("Your choice: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	cleanedInput := cleanUserInput(input)
	return cleanedInput, nil

}

// cleanUserInput function cleans player input
func cleanUserInput(input string) string {
	if runtime.GOOS == "windows" {
		input = strings.Replace(input, "\r\n", "", -1)
	} else {
		input = strings.Replace(input, "\n", "", -1)
	}
	return input
}
