package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dpnetca/exercise/chooseAdventure/v4/cyoa"
)

func main() {
	var file string
	var port int
	flag.IntVar(&port, "port", 5000, "Port to start the server on. Default:5000")
	flag.StringVar(&file, "file", "story.json", "json story file. Default:story.json")

	flag.Parse()

	story, err := cyoa.LoadStoryFromJSON(file)
	if err != nil {
		log.Fatalln(err)
	}

	// using Custom Template:
	// tpl := template.Must(template.New("").Parse("Hello"))
	// storyHandler := cyoa.NewStoryHandler(story, cyoa.WithTemplate(tpl))

	// storyHandler := cyoa.NewStoryHandler(story, cyoa.WithPathFunc(pathFunc))
	tpl := template.Must(template.New("").Parse(storyTpl))
	storyHandler := cyoa.NewStoryHandler(story, cyoa.WithPathFunc(pathFunc), cyoa.WithTemplate(tpl))

	// storyHandler := cyoa.NewStoryHandler(story)
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/story/", storyHandler)

	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting Server on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, mux)) // same as doing err=...if err...log.Fatal(err)

}

func pathFunc(r *http.Request) string {
	path := r.URL.Path
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]

}

const storyTpl = `
<!DOCTYPE html>
<html>
  <head>
    <title> Choose Your Own Adventure </title>
  </head>
  <body>
	<h1>{{.Title}}</h1>
	{{range .Story}}
	<p>{{.}}</p>
	{{end}}
	<ul>
	{{range .Options}}
	  <li> <a href="/story/{{.Arc}}">{{ .Arc}}</a>  - {{.Text}}</li>
	{{end}}
	</ul>
  </body>
`
