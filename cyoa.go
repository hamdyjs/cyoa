package cyoa

import (
	"encoding/json"
	"io"
)

// Story is the whole story
type Story map[string]Arc

// Arc is a part of the story
type Arc struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option is what directs the story towards a different arc
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// JSONStory takes a reader and returns a Story
func JSONStory(r io.Reader) (Story, error) {
	var story Story
	d := json.NewDecoder(r)
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
