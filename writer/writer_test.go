package writer

import (
	"testing"

	"catmd/reader/jekyll"
)

func TestAddTitleToArticleBody(t *testing.T) {
	article := jekyll.Article{
		Path:        "path/test.md",
		Title:       "Title",
		Description: "Description",
		RawBody: `A paragraph:
* list item 1
* list item 2`,
	}
	const expected = `# Title

A paragraph:
* list item 1
* list item 2`

	article = addTitleToArticleBody(article)
	if article.RawBody != expected {
		t.Errorf("Failed to add title to article body.\nExpected: %s\nGot: %s", expected, article.RawBody)
	}
}

func TestIncrementArticleHeadings(t *testing.T) {
	article := jekyll.Article{
		Path:        "path/test.md",
		Title:       "Title",
		Description: "Description",
		RawBody: `# Title

A paragraph:
* list item 1
* list item 2
## Heading 2
### This is another heading

And another paragraph`,
	}
	const expected = `## Title

A paragraph:
* list item 1
* list item 2
### Heading 2
#### This is another heading

And another paragraph`

	article = incrementArticleHeadings(article)
	if article.RawBody != expected {
		t.Errorf("Failed to increment headings for article.\nExpected: %s\nGot: %s", expected, article.RawBody)
	}
}
