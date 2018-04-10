package writer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/joaofnfernandes/catmd/reader"
)

const jekyllFrontMatter = `---
layout: pdf
permalink: /pdf.html
---`

// TODO: add jekyll front matter
func Write(jekyllArticles []reader.ProcessedArticle, outputPath string) {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%s\n\n", jekyllFrontMatter))
	for _, article := range jekyllArticles {
		buf.WriteString(fmt.Sprintf("%s\n\n", article.NewBody))
	}
	err := ioutil.WriteFile(outputPath, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed to write %s: %s", outputPath, err)
	}
}
