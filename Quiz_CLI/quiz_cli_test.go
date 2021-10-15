package main

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestParseCSV(t *testing.T) {

	var testProblems = []struct {
		file     string
		expected []problem
		err      error
	}{
		{
			file:     "mock.csv",
			expected: []problem{{"1+1", "2"}, {"2+2", "4"}, {"3+3", "6"}},
			err:      nil,
		},
		{
			file:     "mock1.csv",
			expected: []problem{{"What is the capital of India?", "Delhi"}, {"When did India gain Independence?", "1947"}},
			err:      nil,
		},
		{
			file:     "nonexistent.csv",
			expected: nil,
			err:      errors.New("could not read file"),
		},
	}

	for _, v := range testProblems {
		got, err := parseCSV(v.file)
		if !reflect.DeepEqual(err, v.err) {
			t.Fail()
		}
		if !reflect.DeepEqual(got, v.expected) {
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
