package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"text/template"

	cyoa "github.com/hamdyjs/go_cyoa"
)

var webTmpl = `
<body>
	<h1>{{.Title}}</h1>
	{{range .Paragraphs}}
		<p>{{.}}</p>
	{{end}}
	<ul>
		{{range .Options}}
			<li><a href="{{.Arc}}">{{.Text}}</a></li>
		{{end}}
	</ul>
</body>`

func main() {
	fileFlag := flag.String("file", "gopher.json", "The json file containing the story")
	flag.Parse()

	fmt.Println("Running story from file:", *fileFlag)
	file, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("%+v", story)
}

func newStoryHandler(story cyoa.Story) http.Handler {
	return storyHandler{story}
}

type storyHandler struct {
	story cyoa.Story
}

func (h storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Parse(webTmpl))
	tmpl.Execute(w, h.story)
}
