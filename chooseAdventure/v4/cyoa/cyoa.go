package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func init() {
	tpl = template.Must(template.New("").Parse(StoryTpl))
}

var tpl *template.Template

type Story map[string]arc

type arc struct {
	Title   string
	Story   []string
	Options []option
}

type option struct {
	Text string
	Arc  string
}

type storyHandler struct {
	s      Story
	t      *template.Template
	Pathfn func(r *http.Request) string
}

type HandlerOptions func(h *storyHandler)

func WithTemplate(t *template.Template) HandlerOptions {
	return func(h *storyHandler) {
		h.t = t
	}
}

func WithPathFunc(fn func(r *http.Request) string) HandlerOptions {
	return func(h *storyHandler) {
		h.Pathfn = fn
	}
}

func defaultPathFn(r *http.Request) string {
	query := r.URL.Query()
	arc, ok := query["arc"]
	if !ok {
		arc = []string{"intro"}
	}
	return arc[0]
}

func NewStoryHandler(s Story, opts ...HandlerOptions) http.Handler {
	h := storyHandler{s, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}
func LoadStoryFromJSON(jsonFile string) (Story, error) {
	var story Story

	file, err := os.Open(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("unable to load story from json: %v", err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&story)
	if err != nil {
		return nil, err
	}

	return story, nil
}

func StoryCLI(story Story) {
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

func (h storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arc := h.Pathfn(r)
	if s, ok := h.s[arc]; ok {
		err := h.t.Execute(w, s)
		if err != nil {
			log.Println(err)
			http.Error(w, "unable to execute template", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Story arc not found", http.StatusNotFound)
		return
	}

}
