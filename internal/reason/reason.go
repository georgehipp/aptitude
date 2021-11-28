package reason

import (
	"aptitude/internal/utility"
	"fmt"
	"math"
	"os"
	"time"
)

func Reason(numQues int) {
	var currQues int = 1
	var userInput int
	var rightAns int
	var wrongAns int
	var totalTime time.Duration
	var responseTime time.Duration
	var names = [6]string{"George", "Tracy", "Wasabi", "Serafina", "Clouso", "Dexi"}

	var adjectives = [5][4]string{
		{"lighter", "heavier", "light", "heavy"},
		{"weaker", "stronger", "weak", "strong"},
		{"duller", "brighter", "dull", "bright"},
		{"smaller", "larger", "small", "large"},
		{"cooler", "hotter", "cold", "hot"},
	}

	var determiner = [1][3]string{
		{"less ", "more ", "not as "},
	}

	instructions := "Select the proper person to answer the question by inputting the proper number:"

	utility.PrintQuestionResults(numQues, currQues, instructions, "-\n", "-\n")

	for userInput != 100 && currQues != numQues+1 {
		selectedNames := twoRndNames(names[:])
		var details string
		var question string
		var result int
		adjRow := utility.RndNum(len(adjectives))
		adjCol := utility.RndNum(2)
		detCol := utility.RndNum(len(determiner[0]))
		adjIni := adjCol % 2
		detIni := 0
		if utility.RndBool() {
			adj := adjectives[adjRow][adjCol]
			details = fmt.Sprint(selectedNames[0], " is ", adj, " than ", selectedNames[1], ".")
		} else {
			adj := adjectives[adjRow][adjCol+2]
			det := determiner[0][detCol]
			detp := " than "
			if det == determiner[0][2] {
				detp = " as "
			}
			if det != determiner[0][1] {
				detIni = 1
			}
			details = fmt.Sprint(selectedNames[0], " is ", det, adj, detp, selectedNames[1], ".")
		}

		fmt.Println("Type 0 and Return / Enter to Continue.")
		fmt.Fprint(os.Stdout, details)
		fmt.Scan(&userInput)
		adjCol = utility.RndNum(2)
		detCol = utility.RndNum(len(determiner[0]))
		adjFin := adjCol % 2
		detFin := 0
		if utility.RndBool() {
			adj := adjectives[adjRow][adjCol]
			question = fmt.Sprint("Who is ", adj, "?")
		} else {
			adj := adjectives[adjRow][adjCol+2]
			det := determiner[0][detCol]
			if det != determiner[0][1] {
				detFin = 1
			}
			question = fmt.Sprint("Who is ", det, adj, "?")
		}

		if adjFin != adjIni && detIni == detFin {
			result = 1
		}
		if adjFin == adjIni && detIni == detFin {
			result = 0
		}
		if adjFin != adjIni && detIni != detFin {
			result = 0
		}
		if adjFin == adjIni && detIni != detFin {
			result = 1
		}

		utility.PrintQuestionResults(numQues, currQues, instructions, "-\n", "-\n")
		fmt.Println("0 -", selectedNames[0], " |  1 -", selectedNames[1])

		fmt.Fprint(os.Stdout, question+"\n")
		startTime := time.Now()
		fmt.Scan(&userInput)
		responseTime = time.Since(startTime)
		totalTime = totalTime + responseTime
		currQues = currQues + 1
		details = question + "\n" + details
		if userInput != 100 && currQues != numQues+2 {
			//result := 0
			switch true {
			case userInput == result:
				result := fmt.Sprintln(userInput, " is Correct! -", math.Round(responseTime.Seconds()), "Seconds Taken")
				rightAns = rightAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			case userInput != result:
				result := fmt.Sprintln(userInput, "is Incorrect! - ", math.Round(responseTime.Seconds()), "Seconds Taken")
				wrongAns = wrongAns + 1
				utility.PrintQuestionResults(numQues, currQues, instructions, result, details)
			}
		}
	}
	if currQues != 2 {
		utility.PrintFinalResults("Reason", rightAns, wrongAns, totalTime, responseTime)
	} else {
		fmt.Println("No Answers")
	}

}

func twoRndNames(names []string) [2]string {
	var selectedNames = [2]string{names[utility.RndNum(5)], names[utility.RndNum(5)]}
	for selectedNames[0] == selectedNames[1] {
		selectedNames[1] = names[utility.RndNum(5)]
	}
	return selectedNames
}
