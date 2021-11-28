package spatial

import (
	"aptitude/internal/utility"
	"fmt"
	"math"
	"time"
)

func Spatial(numQues int) {
	//fmt.Printf("Forward R Symbol: %c\n", '\uA4E3')
	//fmt.Printf("Backward R Symbol: %c\n", '\u042f')
	//fmt.Printf("Upsidedown R Symbol: %c\n", '\uA4E4')

	runes := [3]rune{'\uA4E3', '\uA4E4', '\u042f'}

	utility.ClearTerm()
	var currQues int = 1
	var userInput int
	var rightAns int
	var wrongAns int
	var totalTime time.Duration
	var responseTime time.Duration
	instructions := "What columns contain matching symbols if they can be rotated:"

	utility.PrintQuestionResults(numQues, currQues, instructions, "-\n", "-\n")

	topRow := [2]rune{}
	bottomRow := [2]rune{}
	matchRow := [2]bool{false}
	var matchTotal int

	for userInput != 100 && currQues != numQues+1 {
		matchTotal = 0
		for i := range topRow {
			topRow[i] = runes[utility.RndNum(3)]
			bottomRow[i] = runes[utility.RndNum(3)]
			if topRow[i] == bottomRow[i] || topRow[i] != runes[2] && bottomRow[i] != runes[2] {
				matchRow[i] = true
				matchTotal++
			}
		}

		fmt.Printf("%c  %c\n", topRow[0], topRow[1])
		fmt.Printf("%c  %c\n", bottomRow[0], bottomRow[1])
		startTime := time.Now()
		fmt.Scan(&userInput)
		utility.ClearTerm()
		responseTime = time.Since(startTime)
		totalTime = totalTime + responseTime
		currQues = currQues + 1

		if userInput != 100 && currQues != numQues+2 {
			switch true {
			case userInput == matchTotal:
				result := fmt.Sprintln(userInput, "is Correct! -", math.Round(responseTime.Seconds()), "Seconds Taken")
				details := fmt.Sprintln(matchTotal, "is the number of columns with the same symbol.")
				rightAns = rightAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			case userInput != matchTotal:
				result := fmt.Sprintln(userInput, "is Incorrect! - ", math.Round(responseTime.Seconds()), "Seconds Taken")
				details := fmt.Sprintln(matchTotal, "is the number of columns with the same symbol.")
				wrongAns = wrongAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			}
		}

	}

	utility.ClearTerm()
	if currQues != 2 {
		utility.PrintFinalResults("Spatial", rightAns, wrongAns, totalTime, responseTime)
	} else {
		fmt.Println("No Answers")
	}

}
