package main

import (
	"flag"
	"log"

	"github.com/dpnetca/exercise/chooseAdventure/v4/cyoa"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "story.json", "json story file. Default:story.json")

	flag.Parse()

	story, err := cyoa.LoadStoryFromJSON(file)
	if err != nil {
		log.Fatalln(err)
	}

	cyoa.StoryCLI(story)

}
