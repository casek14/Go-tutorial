package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFile := flag.String("csv", "quiz.csv", "A .csv file with question/answer format")
	timeLimit := flag.Int("time-limit",30,"Time limit for quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		printEror(fmt.Sprintf("File %v was not found.\n", *csvFile))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		printEror(fmt.Sprintf("Provided file has invalid format\n"))
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)


	var correct = 0
problemloop:
	for i, p := range problems {
		fmt.Printf("Question #%d: %s = ", i, p.question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You answered %d out of %d correct !\nk", correct, len(problems))
			break problemloop
		case answer :=   <- answerCh:
			if answer == p.answer {
				correct++
			}
		}


	}
	fmt.Printf("\nYou answered %d out of %d correct !\n", correct, len(problems))

}

type problem struct {
	question string
	answer   string
}

//parse csv lines
func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))
	for i, p := range lines {
		ret[i] = problem{question: p[0],
			answer: p[1]}
	}
	return ret
}

//Print error message and exit the program
func printEror(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}



