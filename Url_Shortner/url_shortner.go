package main

import (
	"fmt"
	"net/http"
)

func main() {

	pathsToUrl := map[string]string{
		"/github":   "https://github.com/iamargus95",
		"/personal": "https://iamargus95.github.io/",
		"/resume":   "https://iamargus95.github.io/assets/documents/Suraj_Kamath.pdf",
	}

	mapHandler := mapHandler(pathsToUrl, defaultMux())
	fmt.Println("Server up and running at http://localhost:8080")
	http.ListenAndServe(":8080", mapHandler)
}

func mapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if url, ok := pathsToUrls[path]; ok {
			http.Redirect(rw, r, url, http.StatusFound)
			return
		}
		fallback.ServeHTTP(rw, r)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!!")
}
