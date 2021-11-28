package number

import (
	"aptitude/internal/utility"
	"fmt"
	"math"
	"sort"
	"time"
)

func Number(numQues int, rngEnd int) {
	utility.ClearTerm()
	var currQues int = 1
	var userInput int
	var rightAns int
	var wrongAns int
	var totalTime time.Duration
	var responseTime time.Duration
	instructions := "When sorted in order, which Number is furthest from the Middle Number:"

	utility.PrintQuestionResults(numQues, currQues, instructions, "-\n", "-\n")

	for userInput != 100 && currQues != numQues+1 {
		rndNums := rndNumAry(rngEnd)
		fmt.Println("Three Numbers:", rndNums)
		startTime := time.Now()
		fmt.Scan(&userInput)
		responseTime = time.Since(startTime)
		totalTime = totalTime + responseTime
		currQues = currQues + 1
		if userInput != 100 && currQues != numQues+2 {
			result := correctAnswer(rndNums)
			switch true {
			case userInput == result:
				result := fmt.Sprintln(userInput, " is Correct! -", math.Round(responseTime.Seconds()), "Seconds Taken")
				details := fmt.Sprintln(rndNums[0], "+", rndNums[1]-rndNums[0], "=", rndNums[1], "+", rndNums[2]-rndNums[1], "=", rndNums[2])
				rightAns = rightAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			case userInput != result:
				result := fmt.Sprintln(userInput, "is Incorrect! - ", math.Round(responseTime.Seconds()), "Seconds Taken")
				details := fmt.Sprintln(rndNums[0], "+", rndNums[1]-rndNums[0], "=", rndNums[1], "+", rndNums[2]-rndNums[1], "=", rndNums[2])
				wrongAns = wrongAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			}
		}
	}
	if currQues != 2 {
		utility.PrintFinalResults("Numbers", rightAns, wrongAns, totalTime, responseTime)
	} else {
		fmt.Println("No Answers")
	}
}

func containsDuplicate(rndNums []int) bool {
	record := make(map[int]bool, len(rndNums))
	for _, n := range rndNums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
	}
	return false
}

func equalDistance(rndNums []int) bool {
	crndNums := make([]int, len(rndNums))
	copy(crndNums, rndNums)
	sort.Ints(crndNums)
	return crndNums[1]-crndNums[0] == crndNums[2]-crndNums[1]
}

func rndNumAry(rngEnd int) []int {
	rndNums := []int{utility.RndNum(rngEnd), utility.RndNum(rngEnd), utility.RndNum(rngEnd)}
	for containsDuplicate(rndNums) || equalDistance(rndNums) {
		rndNums = []int{utility.RndNum(rngEnd), utility.RndNum(rngEnd), utility.RndNum(rngEnd)}
	}
	return rndNums
}

func correctAnswer(rndNums []int) int {
	var answer int
	sort.Ints(rndNums)
	switch true {
	case rndNums[1]-rndNums[0] > rndNums[2]-rndNums[1]:
		answer = rndNums[0]
	case rndNums[1]-rndNums[0] < rndNums[2]-rndNums[1]:
		answer = rndNums[2]
	}
	return answer
}
