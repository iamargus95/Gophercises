package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type problem struct {
	question string
	answer   string
}

type quiz struct {
	problems []problem
	score    int
}

func main() {

	filename := flag.String("csv", "problems.csv", "Use this to add a relative path to your custom quiz.csv")
	timelimit := flag.Int("timer", 30, "Use this flag to set a custom timer.")
	flag.Parse()
	lines, _ := readCSV(*filename)
	qAndA := parseLines(lines)
	qz := quiz{problems: qAndA, score: 0}
	fmt.Println(qz.askQuestion(*timelimit))
}

func readCSV(filename string) ([][]string, error) {

	file, err := os.Open(filename)
	if err != nil {
		err := fmt.Errorf("could not read file")
		return nil, err
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()
	if err != nil {
		err := fmt.Errorf("failed to parse the provided csv file")
		return nil, err
	}

	return lines, nil
}

func parseLines(lines [][]string) []problem {
	ask := make([]problem, len(lines))
	for i, line := range lines {
		ask[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ask
}

func (c *quiz) askQuestion(timeLimit int) string {
	c.score = 0
	// timer := time.NewTimer(time.Duration(timeLimit) * time.Second)

	for i, q := range c.problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, q.question)
		var input string
		fmt.Scanf("%s\n", &input)
		if input == q.answer {
			c.score++
		}
	}

	return "You scored " + strconv.Itoa(c.score) + " out of " + strconv.Itoa(len(c.problems))
}
