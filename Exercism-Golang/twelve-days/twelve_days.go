package twelve

// package main
import (
	"fmt"
	"strings"
)

//type
var DAYS = [12][]string{
	0:  {"first", "a Partridge in a Pear Tree"},
	1:  {"second", "two Turtle Doves"},
	2:  {"third", "three French Hens"},
	3:  {"fourth", "four Calling Birds"},
	4:  {"fifth", "five Gold Rings"},
	5:  {"sixth", "six Geese-a-Laying"},
	6:  {"seventh", "seven Swans-a-Swimming"},
	7:  {"eighth", "eight Maids-a-Milking"},
	8:  {"ninth", "nine Ladies Dancing"},
	9:  {"tenth", "ten Lords-a-Leaping"},
	10: {"eleventh", "eleven Pipers Piping"},
	11: {"twelfth", "twelve Drummers Drumming"},
}

func Song() string {
	var result string
	for i := 1; i <= 12; i++ {
		result += Verse(i) + "\n"
	}
	return strings.Trim(result, "\n")
}

func Verse(n int) string {
	n--
	var present string = DAYS[n][1]
	var day string = DAYS[n][0]
	n--
	for i := n; i >= 0; i-- {
		if i == 0 {
			present += ", and " + DAYS[i][1]
		} else {
			present += ", " + DAYS[i][1]
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", day, present)
}
