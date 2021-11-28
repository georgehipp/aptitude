package main

import (
	"aptitude/internal/number"
	"aptitude/internal/perception"
	"aptitude/internal/reason"
	"aptitude/internal/spatial"
	"aptitude/internal/utility"
	"aptitude/internal/word"
	"fmt"
)

func main() {
	utility.ClearTerm()
	var userInput int
	for userInput != 100 {
		fmt.Println("Please Input Selection:")
		fmt.Println("1 - Numbers")
		fmt.Println("2 - Perception")
		fmt.Println("3 - Reason")
		fmt.Println("4 - Spatial")
		fmt.Println("5 - Words")
		//fmt.Println("6 - All Tests")
		fmt.Println("100 - Quit")
		fmt.Scan(&userInput)
		if userInput != 0 {
			switch true {
			case userInput == 1:
				number.Number(5, 20)
			case userInput == 2:
				perception.Perception(5)
			case userInput == 3:
				reason.Reason(5)
			case userInput == 4:
				spatial.Spatial(5)
			case userInput == 5:
				word.Word(5)
			case userInput == 6:
				fmt.Println("All Tests - Not Enabled")
			case userInput == 100:
				fmt.Println("Goodbye!")
			default:
				fmt.Println("Goodbye!")
			}
		}
	}
}
