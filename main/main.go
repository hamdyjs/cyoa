package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/hamdyjs/cyoa"
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
	portFlag := flag.Int("port", 3000, "The port to serve the web application on")
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

	sh := storyHandler{story}
	fmt.Println("Listening on port:", *portFlag)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portFlag), sh))
}

type storyHandler struct {
	story cyoa.Story
}

func (h storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	arcKey := r.URL.Path[1:]
	if arcKey == "" || arcKey == "/" {
		arcKey = "intro"
	}

	if arc, ok := h.story[arcKey]; ok {
		tmpl := template.Must(template.New("").Parse(webTmpl))
		err := tmpl.Execute(w, arc)
		if err != nil {
			http.Error(w, "Something went wrong!", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Arc not found!", http.StatusNotFound)
	}
}
