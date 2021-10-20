package main

import (
	"flag"
	"fmt"
	cyoa "iamargus95/Gophercises/CYOA"
	"log"
	"net/http"
)

func main() {

	file := flag.String("file", "gopher.json", "The name of the JSON file with the CYOA story.")
	flag.Parse()
	fmt.Println("Starting server on https://localhost:8080 ...")
	content := cyoa.ParseJson(*file)
	handler := cyoa.StoryPage(content)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
