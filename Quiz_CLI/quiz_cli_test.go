package main

import (
	"reflect"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestParseCSV(t *testing.T) {

	var testProblems = []struct {
		file     string
		expected []problem
	}{
		{
			file:     "mock.csv",
			expected: []problem{{"1+1", "2"}, {"2+2", "4"}, {"3+3", "6"}},
		},
		{
			file:     "mock1.csv",
			expected: []problem{{"What is the capital of India?", "Delhi"}, {"When did India gain Independence?", "1947"}},
		},
		{
			file:     "nonexistent.csv",
			expected: nil,
		},
	}

	for _, v := range testProblems {
		got, _ := parseCSV(v.file)
		if !reflect.DeepEqual(got, v.expected) {
			t.Fail()
		}
	}
}

func testEachQuestion(t *testing.T) {
	ticker := time.NewTicker(time.Duration(2) * time.Second).C
	input := make(chan string)
	answer := "2"
	var ans int
	var err error
	allDone := make(chan bool)
	go func() {
		ans, err = eachQuestion(answer, ticker, input)
		allDone <- true
	}()
	input <- "2"

	<-allDone
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, ans, 1)
}

func TestEachQuestion(t *testing.T) {
	t.Run("Test each question ", testEachQuestion)
}
