package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func getRandomNumber(min int, max int) int {

	return rand.Intn(max-min+1) + min
}

func main() {
	fmt.Println(rand.Intn(10-5) + 1)

	scanner := bufio.NewScanner(os.Stdin)
	minNum := 1
	maxNum := 10

	randomNumber := getRandomNumber(1, 10)

	lives := 5
	attemps := 0
	fmt.Println("welcome to our guessing game.\nYou are required to guess the random number.")

	for {
		fmt.Printf("You have %d lives left.\n", lives)
		lives = lives - 1
		attemps = attemps + 1
		if lives < 0 {
			fmt.Printf("You lost!!\nThe number is %d", randomNumber)
			break
		}

		fmt.Printf("Enter a number (%d - %d):", minNum, maxNum)
		scanner.Scan()
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		if num > randomNumber {
			fmt.Println("Your guess is wrong. Please guess lower.")
		} else if num < randomNumber {
			fmt.Println("Your guess is wrong. Please guess higher.")
		} else {
			fmt.Printf("Congratulation!!\nYou won in %d attemps!!!\n", attemps)
			break
		}

		continue
	}

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()

}
