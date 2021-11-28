package utility

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

/*
type ApTest struct {
	Name     string
	RightAns int
	WrongAns int
	avgTime  int
	perRight float64
}
*/
func PrintFinalResults(name string, rightAns int, wrongAns int, totalTime time.Duration, responseTime time.Duration) {
	ClearTerm()
	fmt.Println("---", name, "Final Results ---")
	totalTime = totalTime - responseTime
	perRight := (float64(rightAns) / (float64(rightAns) + float64(wrongAns))) * 100
	fmt.Println(perRight, "% Correct")
	fmt.Println("Correct:", rightAns)
	fmt.Println("Wrong:", wrongAns)
	fmt.Println("Avg Time Per Response:", int(math.Round(totalTime.Seconds()))/(rightAns+wrongAns), "Seconds")
	fmt.Println("----------------------")
}

func PrintQuestionResults(numQues int, currQues int, instructions string, result string, details string) {
	ClearTerm()
	fmt.Println("--- Input 100 to Exit -", numQues, "Questions Total ---")
	fmt.Println("--- Results ---")
	fmt.Fprint(os.Stdout, result)
	fmt.Fprint(os.Stdout, details)
	fmt.Println("")
	fmt.Println("---", currQues, "---")
	fmt.Println(instructions)
}

func ClearTerm() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func RndNum(rngEnd int) int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(rngEnd)
}

func RndBool() bool {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Float32() < 0.5
}

func UserInput(userInput *int) time.Duration {
	startTime := time.Now()
	fmt.Scan(&userInput)
	return time.Since(startTime)
}
