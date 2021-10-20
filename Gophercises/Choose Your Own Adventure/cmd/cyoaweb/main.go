package main

import (
	"encoding/json"
	"flag"
	"fmt"
	cyoa "iamargus95/Gophercises/CYOA"
	"os"
)

func main() {

	file := flag.String("file", "gopher.json", "The name of the JSON file with the CYOA story.")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)
	content := parseJson(*file)
	fmt.Printf("%+v\n", content)
}

func parseJson(filename string) cyoa.Story {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data := json.NewDecoder(file)
	var story cyoa.Story
	if err := data.Decode(&story); err != nil {
		panic(err)
	}
	return story
}
