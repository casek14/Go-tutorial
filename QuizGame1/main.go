package main

import (
	"flag"
	"os"
	"fmt"
	"encoding/csv"
)

func main() {

	csvFilename := flag.String("csv","problem.csv","a csv file in format 'question,answer'")
	flag.Parse()

	file,err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file %s \n",*csvFilename))

	}

	_ = file

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
		exit("Failed to parse the provided file")
	}

	fmt.Println(lines)

}


func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)

}
