package main

import (
	"log"

	"github.com/dpnetca/exercise/chooseAdventure/v1/cyoa"
)

func main() {
	file := "story.json"

	story, err := cyoa.LoadStoryFromJSON(file)
	if err != nil {
		log.Fatalln(err)
	}

	cyoa.StoryCLI(story)
}
