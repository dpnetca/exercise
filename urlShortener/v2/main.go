package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dpnetca/exercise/urlShortener/v2/urlshort"
)

func main() {
	var yamlFile, jsonFile string
	flag.StringVar(&yamlFile, "yaml", "", "yaml file")
	flag.StringVar(&jsonFile, "json", "", "json file")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlData, err := getYaml(yamlFile)
	if err != nil {
		panic(err)
	}
	yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
	if err != nil {
		panic(err)
	}

	jsonData, err := getJson(jsonFile)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := urlshort.JSONHandler(jsonData, yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func getYaml(yamlFile string) ([]byte, error) {

	var yamlData []byte
	var err error

	if yamlFile != "" {
		yamlData, err = readFile(yamlFile)
		if err != nil {
			return nil, err
		}
	} else {
		yamlData = []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	}
	return []byte(yamlData), nil
}

func getJson(jsonFile string) ([]byte, error) {

	var jsonData []byte
	var err error

	if jsonFile != "" {
		jsonData, err = readFile(jsonFile)
		if err != nil {
			return nil, err
		}
	} else {
		jsonData = []byte(`
[ 
	{
		"path": "/google",
  		"url": "https://google.ca"
	},
	{
		"path": "/duck",
  		"url": "https://duckduckgo.com/"
	}
]
`)
	}
	return []byte(jsonData), nil
}
func readFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
