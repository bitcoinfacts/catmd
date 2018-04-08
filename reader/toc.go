package reader

import (
	"io/ioutil"
	"log"
	"path"

	"gopkg.in/yaml.v2"
)

const MarkdownExtension = ".md"

type docsToc []tocSection

type tocSection struct {
	Title string   `yaml:"title"`
	Path  []string `yaml:"docs"`
}

// ReadTocFile reads the yaml file used in the glacier docs to find out
// which markdown files to explore and in what order
func readTocFile(filepath, baseDir, extension string) []string {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	toc := docsToc{}
	err = yaml.Unmarshal(b, &toc)
	if err != nil {
		log.Fatalf("Failed to unmarshal order file %s: %s", filepath, err)
	}
	paths := []string{}
	for _, section := range toc {
		for _, p := range section.Path {
			paths = append(paths, path.Join(baseDir, p+extension))
		}
	}
	return paths
}
