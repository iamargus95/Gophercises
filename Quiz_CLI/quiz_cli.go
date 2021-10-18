package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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

	qAndA, err := parseCSV(*filename)
	if err != nil {
		log.Fatal(err)
	}
	qz := quiz{problems: qAndA, score: 0}

	score, err := qz.askQuestion(*timelimit)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("You scored " + strconv.Itoa(score) + " out of " + strconv.Itoa(len(qAndA)) + "\n")
	} else {
		fmt.Println("You scored " + strconv.Itoa(score) + " out of " + strconv.Itoa(len(qAndA)) + "\n")
	}
}

func parseCSV(filename string) ([]problem, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)

	lines, _ := r.ReadAll()

	ask := make([]problem, len(lines))
	for i, line := range lines {
		ask[i] = problem{
			question: strings.TrimSpace(line[0]),
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ask, nil
}

func (c *quiz) askQuestion(timeLimit int) (int, error) {
	c.score = 0

	ticker := time.NewTicker(time.Duration(timeLimit) * time.Second)
	input := make(chan string)

	go getInput(input)

	for i, q := range c.problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, q.question)
		output, err := eachQuestion(q.answer, ticker.C, input)
		if err != nil && output == -1 {
			return c.score, err
		}
		if output == 1 {
			c.score++
		}
	}

	return c.score, nil
}

func getInput(input chan string) {

	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input <- result
	}
}

func eachQuestion(answer string, stop <-chan time.Time, input <-chan string) (int, error) {

	select {
	case <-stop:
		return -1, fmt.Errorf("time out")
	case ans := <-input:
		score := 0
		if strings.Compare(strings.Trim(strings.ToLower(ans), "\n"), answer) == 0 {
			score = 1
		} else {
			return 0, nil
		}
		return score, nil
	}
}
