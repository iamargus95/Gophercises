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
	port := flag.String("port", "8080", "Input the port at which you want the service to run.")
	flag.Parse()

	fmt.Printf("Starting server on http://localhost:%s ...\n", *port)

	content := cyoa.ParseJson(*file)
	handler := cyoa.StoryPage(content)

	log.Fatal(http.ListenAndServe(":"+*port, handler))
}
