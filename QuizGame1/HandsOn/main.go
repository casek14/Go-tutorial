package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("csv", "quiz.csv", "A .csv file with question/answer format")
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
	ask(problems)
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

func ask(problems []problem) {
	var correct = 0
	for i, p := range problems {
		fmt.Printf("Question #%d: %s = ", i, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
		}
	}
	fmt.Printf("You answered %d out of %d correct !\nk", correct, len(problems))
}
