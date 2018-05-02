package writer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/joaofnfernandes/catmd/reader"
	"github.com/joaofnfernandes/catmd/reader/jekyll"
	"github.com/joaofnfernandes/catmd/reader/urlparser"
)

func Write(book reader.Book, outputPath string) {
	book = updateArticles(book, addTitleToArticleBody)
	book = updateArticles(book, incrementArticleHeadings)
	book = updateArticles(book, fixUrlsFunc(book))

	err := ioutil.WriteFile(outputPath, []byte(book.String()), 0644)
	if err != nil {
		log.Fatalf("Failed to write %s: %s", outputPath, err)
	}
}

type transformFunc func(jekyll.Article) jekyll.Article

// updateArticles updates each article in the book, with a new article
func updateArticles(book reader.Book, f transformFunc) reader.Book {
	for i := 0; i < len(book); i++ {
		for j := 0; j < len(book[i].Articles); j++ {
			book[i].Articles[j] = f(book[i].Articles[j])
		}
	}
	return book
}

// addTitleToBody updates the body of an article to include its title
func addTitleToArticleBody(article jekyll.Article) jekyll.Article {
	const headingFormat = "# %s\n\n%s"
	article.RawBody = fmt.Sprintf(headingFormat, article.Title, article.RawBody)
	return article
}

// incrementArticleHeadings increments the headings of the article by one.
// H1->H2, H2->H3, and so forward
func incrementArticleHeadings(article jekyll.Article) jekyll.Article {
	// Matches a valid markdown heading, e.g ## Glacier vs. others
	const exp = `#{1,6}\s.*?\n`
	re := regexp.MustCompile(exp)

	toParse := []byte(article.RawBody)
	parsed := bytes.Buffer{}

	for len(toParse) > 0 {
		loc := re.FindIndex(toParse)
		if loc == nil {
			parsed.Write(toParse)
			toParse = []byte{}
		} else {
			parsed.WriteString(fmt.Sprintf("%s#%s", toParse[0:loc[0]], toParse[loc[0]:loc[1]]))
			toParse = toParse[loc[1]:]
		}
	}
	article.RawBody = parsed.String()
	return article
}

// fixUrlsFunc returns a function that converts urls in an article.
// External urls are kept as is, while internal urls are converted to anchors
// so that users can navigate when  all the markdown articles are rendered as a
// single page
func fixUrlsFunc(book reader.Book) transformFunc {
	m := book.MapPathToTitle()

	return func(article jekyll.Article) jekyll.Article {
		// Matches a markdown link, e.g. [example](https://example.org)
		const exp = `\[(.*?)\]\((.*?)\)`
		re := regexp.MustCompile(exp)

		toParse := []byte(article.RawBody)
		parsed := bytes.Buffer{}

		for len(toParse) > 0 {
			loc := re.FindSubmatchIndex(toParse)
			if loc == nil {
				parsed.Write(toParse)
				toParse = []byte{}
			} else {
				beforeLink := string(toParse[:loc[0]])
				linkText := string(toParse[loc[2]:loc[3]])
				linkUrl := string(toParse[loc[4]:loc[5]])

				toParse = toParse[loc[1]:]

				parsed.WriteString(beforeLink)
				transformedUrl, err := urlparser.TransformUrl(string(linkUrl), m)
				if err != nil {
					// Could not transform url, remove it but leave text
					parsed.WriteString(linkText)
				} else {
					parsed.WriteString(fmt.Sprintf("[%s](%s)", linkText, transformedUrl))
				}
			}
		}
		article.RawBody = parsed.String()
		return article
	}
}
