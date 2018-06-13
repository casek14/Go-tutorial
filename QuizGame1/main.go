package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFilename := flag.String("csv", "problem.csv", "a csv file in format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file %s \n", *csvFilename))

	}

	_ = file

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided file")
	}
	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Question %v. %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You answered %v out of %v correct.\n", correct, len(problems))
}

//problem struct definition
type problem struct {
	//q = question
	q string
	//a = answer
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{q: line[0], a: line[1]}
	}
	return ret
}

//print error message and exit program
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)

}
