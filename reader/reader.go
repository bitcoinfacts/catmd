package reader

import (
	"log"

	"github.com/joaofnfernandes/catmd/reader/jekyll"
	"github.com/joaofnfernandes/catmd/reader/urlparser"
)

// catmd docs_toc.yml _docs bin/glacier.md

// Read reads the yaml file used in to render the left nav
// of the glacier docs, parses and cleans all those markdown
// files mention there, and returns the content of all of them
// concatenated
func Read(tocFilepath, basedir string) []ProcessedArticle {
	filesToExplore := readTocFile(tocFilepath, basedir, MarkdownExtension)
	articles := importArticles(filesToExplore)
	return fixUrls(articles)
}

// importArticles imports a markdown article and associates it to
// its path.
func importArticles(filepaths []string) []jekyll.Article {
	articles := []jekyll.Article{}

	for _, filepath := range filepaths {
		article, err := jekyll.New(filepath)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, article)
	}
	return articles
}

type ProcessedArticle struct {
	OriginalArticle jekyll.Article
	NewBody         string
}

// fixUrls replaces urls to other articles to use a url fragment
// so that all the articles can be combined into a single one.
func fixUrls(articles []jekyll.Article) []ProcessedArticle {
	processedArticles := []ProcessedArticle{}
	m := filepathToTitle(articles)

	for _, article := range articles {
		processed := ProcessedArticle{article, ""}
		processed.NewBody = processed.OriginalArticle.TransformUrls(func(url string) string {
			newUrl, err := urlparser.TransformUrl(url, m)
			if err != nil {
				// URL pointing to an internal resource that we can't match.
				log.Printf("Failed to fix url for article at %s: %s", article.Path, err)
				return ""
			}
			return newUrl
		})
		processedArticles = append(processedArticles, processed)
	}
	return processedArticles
}

// filepathToTitle maps an article's filename to its title
func filepathToTitle(articles []jekyll.Article) map[string]string {
	m := map[string]string{}
	for _, article := range articles {
		m[article.Path] = article.Title
	}
	return m
}
