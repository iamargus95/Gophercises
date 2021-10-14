package main

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

var testProblems = []struct {
	question string
	answer   string
}{
	{
		question: "1+1",
		answer:   "2",
	},
	{
		question: "2+2",
		answer:   "4",
	},
	{
		question: "3+3",
		answer:   "6",
	},
}

func TestReadCSV(t *testing.T) {

	got, err := readCSV("mock.csv")
	if err != nil {
		t.Error(err)
	}

	for i, v := range got {
		if testProblems[i].question != v[0] {
			t.Fail()
		}
		if testProblems[i].answer != v[1] {
			t.Fail()
		}
	}
}

func TestParseCSV(t *testing.T) {

	input, _ := readCSV("mock.csv")
	got := parseLines(input)

	for i, v := range got {
		if testProblems[i].question != v.question {
			t.Fail()
		}
		if testProblems[i].answer != v.answer {
			t.Fail()
		}
	}
}

func testEachQuestion(t *testing.T) {
	timer := time.NewTimer(time.Duration(2) * time.Second).C
	done := make(chan string)
	answer := "2"
	var ans int
	var err error
	allDone := make(chan bool)
	go func() {
		ans, err = eachQuestion(answer, timer, done)
		allDone <- true
	}()
	done <- "2"

	<-allDone
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, ans, 1)
}

func TestEachQuestion(t *testing.T) {
	t.Run("Test each question ", testEachQuestion)
}
