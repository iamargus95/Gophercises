package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday " + name + "! You are now " + strconv.Itoa(age) + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	res := ("Welcome to my party, " + name + "!\nYou have been assigned to table " + fmt.Sprintf("%X", table) + ". Your table is " + direction + ", exactly " + fmt.Sprintf("%.1F", distance) + " meters from here.\nYou will be sitting next to " + neighbor + ".")
	return res
}
