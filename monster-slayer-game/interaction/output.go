package interaction

import (
	"fmt"
	"os"
)

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game..!")
	fmt.Println("Good Luck!")
}

func ShowAvailableActions(isSpecialRound bool) {
	fmt.Println("Please choose your action")
	fmt.Println("---------------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal Yourself")
	if isSpecialRound {
		fmt.Println("(3) Do a Special Attack")
	}
}

func DeclareWinner(winner string) {
	fmt.Println("---------------------------")
	fmt.Println("Game Over.")
	fmt.Println("---------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLog(rounds *[]RoundData) {
	file, err := os.Create("game-log.txt")
	if err != nil {
		fmt.Println("Saving a logfile fail. Exiting..!")
		return
	}
	for _, data := range *rounds {
		_, err := file.WriteString(data.LogInfo())
		if err != nil {
			fmt.Println("Writing a logfile fail. Exiting..!")
			continue
		}
	}
	_ = file.Close()

}

func WriteLogMap(rounds *[]RoundData) {
	file, err := os.Create("game-log-map.txt")
	if err != nil {
		fmt.Println("Saving a logfile fail. Exiting..!")
		return
	}
	for index, data := range *rounds {
		_, err := file.WriteString(data.LogMap(index))
		if err != nil {
			fmt.Println("Writing a logfile fail. Ignoring..!")
			continue
		}
	}
	_ = file.Close()

}
