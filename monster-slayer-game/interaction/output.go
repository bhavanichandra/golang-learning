package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

func PrintGreeting() {
	//fmt.Println("MONSTER SLAYER")
	asciiFigure := figure.NewColorFigure("MONSTER SLAYER", "", "green", true)
	asciiFigure.Print()

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
	asciiFigure := figure.NewColorFigure("Game Over.", "", "red", true)
	asciiFigure.Print()
	fmt.Println("---------------------------")
	fmt.Printf("%v won!\n", winner)
}

func WriteLog(rounds *[]RoundData) {
	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Saving a logfile fail. Exiting..!")
		return
	}

	executablePath := filepath.Dir(executable)

	file, err := os.Create(executablePath + "/game-log.txt")
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

	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Saving a logfile fail. Exiting..!")
		return
	}

	executablePath := filepath.Dir(executable)
	file, err := os.Create(executablePath + "/game-log-map.txt")
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
