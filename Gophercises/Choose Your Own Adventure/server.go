package cyoa

import (
	"log"
	"net/http"
	"strings"
	"text/template"
)

func StoryPage(s Story) http.Handler {
	return Handler{s}
}

type Handler struct {
	Title Story
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("./default.html"))
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]
	if chapter, ok := h.Title[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Println(err)
			http.Error(w, "Something went wrong ...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
