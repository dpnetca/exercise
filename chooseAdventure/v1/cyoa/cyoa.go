package cyoa

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type arc struct {
	Title   string
	Story   []string
	Options []option
}

type option struct {
	Text string
	Arc  string
}

func LoadStoryFromJSON(jsonFile string) (map[string]arc, error) {
	var story map[string]arc

	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("unable to load story from json: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read data: %v", err)
	}

	err = json.Unmarshal(data, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func StoryCLI(story map[string]arc) {
	section := story["intro"]
	for {
		fmt.Printf("%v\n*****************************************\n\n", section.Title)
		for _, p := range section.Story {
			fmt.Printf("  %v\n\n", p)
		}
		if len(section.Options) > 0 {
			nextArc, err := getNextArc(section.Options)
			if err != nil {
				panic(err)
			}
			section = story[nextArc]
		} else {
			fmt.Println("The End")
			return
		}
	}
}

func getNextArc(options []option) (string, error) {
	fmt.Print("\nOptions:\n--------\n")
	for i, opt := range options {
		fmt.Printf("%d - %v\n", i+1, opt.Text)
	}
	var choice int
	fmt.Print("Choose: ")
	for {
		fmt.Scanln(&choice)
		if choice < 1 || choice > len(options) {
			fmt.Printf("Invalid Choice (%v), please choose again: ", choice)
		} else {
			break
		}

	}
	fmt.Print("\n\n")
	return options[choice-1].Arc, nil
}
