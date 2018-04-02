package article

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/gernest/front"
)

type frontMatter struct {
	Title       string
	Description string
}

type article struct {
	FrontMatter frontMatter
	Body        string
}

func New(markdown string) *article {
	m := front.NewMatter()
	m.Handle("---", front.YAMLHandler)
	fm, body, err := m.Parse(strings.NewReader(markdown))
	if err != nil {
		log.Fatal(err)
	}

	title := fmt.Sprintf("%s", fm["title"])
	description := fmt.Sprintf("%s", fm["description"])
	if title == "" {
		log.Fatal("Markdown file doesn't have title")
	}
	return &article{frontMatter{title, description}, body}
}

func (a *article) String() string {
	const mkdown = `# {{ .FrontMatter.Title }}

{{ .Body }}`

	t := template.Must(template.New("mkdown").Parse(mkdown))
	buffer := bytes.NewBufferString("")
	err := t.Execute(buffer, a)
	if err != nil {
		log.Fatal(err)
	}
	return buffer.String()
}
