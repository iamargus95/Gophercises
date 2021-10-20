package cyoa

import (
	"net/http"
	"text/template"
)

func StoryPage(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("./default.html"))
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
}
