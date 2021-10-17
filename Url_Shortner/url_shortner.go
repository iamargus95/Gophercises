package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	yaml "gopkg.in/yaml.v3"
)

type pathURL struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func main() {

	var yamlFile string
	flag.StringVar(&yamlFile, "yaml", "default.yaml", "path to yaml file.")
	flag.Parse()

	yamlData, err := yamlReader(yamlFile)
	if err != nil {
		panic(err)
	}

	yamlHandler, err := parseYAML([]byte(yamlData), defaultMux())
	if err != nil {
		panic(err)
	}

	fmt.Println("Server up and running at http://localhost:8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func yamlReader(filename string) ([]byte, error) {
	yamlData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read YAML file")
	}
	return yamlData, nil
}

func parseYAML(yamlData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var paths []pathURL
	err := yaml.Unmarshal(yamlData, &paths)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall yaml data")
	}

	pathsToUrls := make(map[string]string)
	for _, v := range paths {
		pathsToUrls[v.Path] = v.Url
	}

	return mapHandler(pathsToUrls, fallback), nil
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
