package writer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/joaofnfernandes/md2pdf/reader"
)

const jekyllFrontMatter = `---
title: Glacier Protocol
---
`

// TODO: add jekyll front matter
func Write(jekyllArticles []reader.ProcessedArticle, outputPath string) {
	var buf bytes.Buffer
	for _, article := range jekyllArticles {
		buf.WriteString(fmt.Sprintf("%s\n\n", article.NewBody))
	}
	err := ioutil.WriteFile(outputPath, buf.Bytes(), 0644)
	if err != nil {
		log.Fatalf("Failed to write %s: %s", outputPath, err)
	}
}
