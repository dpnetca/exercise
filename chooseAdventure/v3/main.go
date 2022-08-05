package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/dpnetca/exercise/chooseAdventure/v3/cyoa"
)

func main() {
	var method, file string
	flag.StringVar(&method, "method", "web", "play story via 'CLI' or 'WEB' inerface. Default:WEB")
	flag.StringVar(&file, "file", "story.json", "json story file. Default:story.json")

	flag.Parse()

	story, err := cyoa.LoadStoryFromJSON(file)
	if err != nil {
		log.Fatalln(err)
	}

	switch strings.ToLower(method) {
	case "cli":
		cyoa.StoryCLI(story)
	case "web":
		storyHandler, err := cyoa.StoryWeb(story)
		if err != nil {
			log.Fatalln(err)
		}

		http.ListenAndServe(":5000", storyHandler)

	default:
		log.Fatalln("unkown method")
	}

}
