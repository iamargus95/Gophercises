package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	filename := flag.String("csv", "problems.csv", "Use this to add a relative path to your custom quiz.csv")
	flag.Parse()
	parseCSV(*filename)

}

func parseCSV(filename string) {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("clould not read file")
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		fmt.Println("Failed to parse the provided csv file.")
		os.Exit(1)
	}

	fmt.Println(lines)
}
