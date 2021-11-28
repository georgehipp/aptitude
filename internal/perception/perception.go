package perception

import (
	"aptitude/internal/utility"
	"fmt"
	"math"
	"time"
)

func Perception(numQues int) {
	utility.ClearTerm()
	var currQues int = 1
	var userInput int
	var rightAns int
	var wrongAns int
	var totalTime time.Duration
	var responseTime time.Duration
	instructions := "How many columns are the same letter, regardless of case:"

	utility.PrintQuestionResults(numQues, currQues, instructions, "-\n", "-\n")

	topRow := [4]string{}
	bottomRow := [4]string{}
	matchRow := [4]bool{false}
	var matchTotal int

	for userInput != 100 && currQues != numQues+1 {
		matchTotal = 0
		for i := range topRow {
			randomRune := rune(utility.RndNum(26))
			topRow[i] = string('A' + randomRune)
			if utility.RndBool() {
				bottomRow[i] = string('a' + randomRune)
				matchRow[i] = true
				matchTotal++
			} else {
				randomRune = rune(utility.RndNum(26))
				bottomRow[i] = string('a' + randomRune)
				if topRow[i] == string('A'+randomRune) {
					matchRow[i] = true
					matchTotal++
				}
			}
		}

		fmt.Println(topRow)
		fmt.Println(bottomRow)
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
				details := fmt.Sprintln(matchTotal, "is the number of columns with the same letter.")
				rightAns = rightAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			case userInput != matchTotal:
				result := fmt.Sprintln(userInput, "is Incorrect! - ", math.Round(responseTime.Seconds()), "Seconds Taken")
				details := fmt.Sprintln(matchTotal, "is the number of columns with the same letter.")
				wrongAns = wrongAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			}
		}

	}

	utility.ClearTerm()
	if currQues != 2 {
		utility.PrintFinalResults("Perception", rightAns, wrongAns, totalTime, responseTime)
	} else {
		fmt.Println("No Answers")
	}
}
