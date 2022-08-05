package main

import (
	"log"
	"net/http"

	"github.com/dpnetca/exercise/chooseAdventure/v2/cyoa"
)

func main() {
	file := "story.json"

	story, err := cyoa.LoadStoryFromJSON(file)
	if err != nil {
		log.Fatalln(err)
	}

	// cyoa.StoryCLI(story)
	storyHandler, err := cyoa.StoryWeb(story)
	if err != nil {
		log.Fatalln(err)
	}

	http.ListenAndServe(":5000", storyHandler)
}
