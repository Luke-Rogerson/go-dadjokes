package main

import (
	"bufio"
	"fmt"
	"os"
)

func getUserInput() string {
	var input string
	for {
		scanner := bufio.NewReader(os.Stdin)
		input, _ = scanner.ReadString('\n')
		if len(input) > 1 {
			return input
		}
		fmt.Println("You must enter something...")
	}
}

func main() {
	fmt.Println("What do you want to hear a joke about?")
	input := getUserInput()
	fmt.Println("YAY", input)
}
