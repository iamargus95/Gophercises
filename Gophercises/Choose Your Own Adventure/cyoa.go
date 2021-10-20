package cyoa

import (
	"encoding/json"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func ParseJson(filename string) Story {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	data := json.NewDecoder(file)
	var story Story
	if err := data.Decode(&story); err != nil {
		panic(err)
	}
	return story
}
