package word

import (
	"aptitude/internal/utility"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Word(numQues int) {

	var oddWordList = [9][3]string{
		{"halt", "stop", "cold"},
		{"up", "down", "street"},
		{"car", "plane", "road"},
		{"below", "under", "letter"},
		{"gold", "silver", "medal"},
		{"cheese", "milk", "dairy"},
		{"table", "chair", "wood"},
		{"talk", "shout", "cough"},
		{"river", "stream", "puddle"},
	}

	var numbers = [3]string{"- 0 -", "- 1 -", "- 2 -"}

	utility.ClearTerm()
	var currQues int = 1
	var userInput int
	var rightAns int
	var wrongAns int
	var totalTime time.Duration
	var responseTime time.Duration
	instructions := "Which word does not belong, input number below word (0-2, left to right):"

	utility.PrintQuestionResults(numQues, currQues, instructions, "-\n", "-\n")

	for userInput != 100 && currQues != numQues+1 {
		a := oddWordList[utility.RndNum(9)]
		ans := a[2]
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		fmt.Println(a)
		fmt.Println(numbers)

		startTime := time.Now()
		fmt.Scan(&userInput)
		utility.ClearTerm()
		responseTime = time.Since(startTime)
		totalTime = totalTime + responseTime
		currQues = currQues + 1

		if userInput != 100 && currQues != numQues+2 {
			switch true {
			case a[userInput] == ans:
				result := fmt.Sprintln(userInput, "is Correct! -", math.Round(responseTime.Seconds()), "Seconds Taken")
				details := fmt.Sprintln(ans, "does not belong.")
				rightAns = rightAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			case a[userInput] != ans:
				result := fmt.Sprintln(userInput, "is Incorrect! - ", math.Round(responseTime.Seconds()), "Seconds Taken")
				details := fmt.Sprintln(ans, "does not belong.")
				wrongAns = wrongAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			}
		}

	}

	utility.ClearTerm()
	if currQues != 2 {
		utility.PrintFinalResults("Words", rightAns, wrongAns, totalTime, responseTime)
	} else {
		fmt.Println("No Answers")
	}
}
