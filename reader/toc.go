package reader

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"gopkg.in/yaml.v2"
)

const MarkdownExtension = ".md"

type Toc []TocSection

type TocSection struct {
	Title string   `yaml:"title"`
	Path  []string `yaml:"docs"`
}

// ReadTocFile unmarshals the table of contents file in yaml format
// TODO: proper error handling
func newToc(filepath, baseDir, extension string) []TocSection {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	toc := Toc{}
	err = yaml.Unmarshal(b, &toc)
	if err != nil {
		log.Fatalf("Failed to unmarshal order file %s: %s", filepath, err)
	}
	for i := 0; i < len(toc); i++ {
		for j := 0; j < len(toc[i].Path); j++ {
			fullPath := fmt.Sprintf("%s%s", path.Join(baseDir, toc[i].Path[j]), extension)
			toc[i].Path[j] = fullPath
		}
	}
	return toc
}
