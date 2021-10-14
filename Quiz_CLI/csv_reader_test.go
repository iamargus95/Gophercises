package main

import (
	"testing"
)

var want = []struct {
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
		if want[i].question != v[0] {
			t.Fail()
		}
		if want[i].answer != v[1] {
			t.Fail()
		}
	}
}
