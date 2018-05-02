package jekyll

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gernest/front"
)

type Article struct {
	Path        string
	Title       string
	Description string
	RawBody     string
}

const frontMatterDelimiter = "---"

func New(filepath string) (Article, error) {
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Article{}, errors.New(fmt.Sprintf("Failed to open jekyll file %s, %s", filepath, err))
	}
	markdown := string(b)
	m := front.NewMatter()
	m.Handle(frontMatterDelimiter, front.YAMLHandler)
	fm, body, err := m.Parse(strings.NewReader(markdown))
	if err != nil {
		return Article{}, errors.New(fmt.Sprintf("Failed to parse jekyll file %s, %s", filepath, err))
	}

	title := fmt.Sprintf("%s", fm["title"])
	if title == "" {
		return Article{}, errors.New(fmt.Sprintf("Jekyll file doesn't have title: %s", filepath))
	}

	description := fmt.Sprintf("%s", fm["description"])
	if description == "" {
		return Article{}, errors.New(fmt.Sprintf("Jekyll file doesn't have description: %s", filepath))
	}

	return Article{filepath, title, description, body}, nil
}

// CompareTo returns true if two articles are equivalent and have the same
// path, title, description, and body. Returns false otherwise.
func (a *Article) CompareTo(b *Article) bool {
	if a.Path != b.Path {
		return false
	}
	if a.Title != b.Title {
		return false
	}
	if a.Description != b.Description {
		return false
	}
	if a.RawBody != b.RawBody {
		return false
	}
	return true
}

func (a *Article) PrintBody() string {
	return a.RawBody
}

// TODO: this can be extracted from here, as the jekyll package should
// just import content without changing it.
// TransformUrls iterates over the article's body in search for urls.
// When a url is found it uses a callback to replace the url.
// func (a Article) TransformUrls(transformFunc func(string) string) string {
// 	parsed, toParse := "", a.RawBody
// 	// Matches a markdown url like [this one](/example)
// 	const markdownUrlRegex = `\[(.*?)\]\((.*?)\)`
// 	re := regexp.MustCompile(markdownUrlRegex)
//
// 	for toParse != "" {
// 		loc := re.FindStringSubmatchIndex(toParse)
// 		if loc == nil {
// 			parsed = parsed + toParse
// 			break
// 		}
// 		if len(loc) < 6 {
// 			// A url match should have submatches, so we'll leave this as is
// 			// and continue parsing since we're not sure what to do with it
// 			parsed = parsed + toParse[loc[0]:loc[1]]
// 			toParse = toParse[loc[1]:]
// 			continue
// 		}
// 		// We've matched and have two submatches
// 		beforeMatch := toParse[:loc[0]]
// 		linkText := toParse[loc[2]:loc[3]]
// 		linkUrl := toParse[loc[4]:loc[5]]
//
// 		transformedUrl := transformFunc(linkUrl)
// 		if transformedUrl == "" {
// 			parsed = parsed + fmt.Sprintf("%s%s", beforeMatch, linkText)
// 		} else {
// 			parsed = parsed + fmt.Sprintf("%s[%s](%s)", beforeMatch, linkText, transformedUrl)
// 		}
// 		toParse = toParse[loc[1]:]
// 	}
// 	return parsed
// }
