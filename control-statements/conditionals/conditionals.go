package conditionals

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	userAge, err := getUserAge()

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	//isOldEnough := userAge >= 18

	if (userAge >= 30 && userAge < 50) || userAge >= 60 {
		fmt.Println("You are eligible for our senior job")
	} else if userAge >= 50 {
		fmt.Println("Best Age")
	} else if userAge >= 18 && userAge < 30 {
		fmt.Println("Welcome to the club")
	} else {
		fmt.Println("Sorry you are not old enough")
	}

	fmt.Println("Goodbye!")

}

func getUserAge() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your age: ")
	userAgeInput, _ := reader.ReadString('\n')
	userAgeInput = strings.Replace(userAgeInput, "\n", "", -1)
	userAge, err := strconv.ParseInt(userAgeInput, 0, 64)
	return int(userAge), err
}
