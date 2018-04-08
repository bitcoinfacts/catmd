package urlparser

import (
	"fmt"
	"regexp"
	"testing"
)

var (
	testUrls = map[string]string{
		// External urls should be left unchanged
		"http://example.com":     "http://example.com",     // external http url
		"https://example.com":    "https://example.com",    // external https url
		"http://example.com/":    "http://example.com/",    // external url, trailing slash
		"http://org.example.com": "http://org.example.com", // external url, multi-domain
		"http://example.com/a/b": "http://example.com/a/b", // external url, multi-path
		"http://example.com?a=b": "http://example.com?a=b", // external url, query string
		"http://example.com#abc": "http://example.com#abc", // external url, with fragment

		// Relative urls with fragments, so be turned into fragments
		"path.md#example":      "#example", // relative url, current path, fragment
		"../path.md#example":   "#example", // relative url, backwards, fragment
		"path/path.md#example": "#example", // relative url, forward, fragment
		"#example":             "#example", // just fragment

		// Relative urls w/o fragments should be turned to fragment, based on article title
		"example":      "#a-title", // relative url, current path
		"path/example": "#b-title", // relative url, multipath
		//TODO: we should be permissive and remove extra slash at the end
		//"path/example/" "#c-title"

		// These are invalid urls since they might point to a resource outside glacier
		"example.md":      "", // relative url, current path
		"../example.md":   "", // relative url, current path
		"path/example.md": "", // relative url, forward
	}

	articleTitle = map[string]string{
		"example":      "a title",
		"path/example": "b title",
	}
)

func TestTransformUrl(t *testing.T) {
	for testUrl, expected := range testUrls {
		got, _ := TransformUrl(testUrl, articleTitle)
		if got != expected {
			t.Errorf("Failed to clean up url %s\nExpected: %s\n Got: %s", testUrl, expected, got)
		}
	}
}

const sampleMarkdown = `# heading

This article has a body with some text
and more text:

* a list item
* and another

a [link](/path) goes here
and finishes
[link](/path)[link](/path)[link](/path)
and more
[link](/path)
end`

func TestBuffers(t *testing.T) {
	parsed := ""
	toParse := sampleMarkdown

	const markdownUrlRegex = `\[(.*?)\]\((.*?)\)`
	re := regexp.MustCompile(markdownUrlRegex)

	for toParse != "" {
		loc := re.FindStringSubmatchIndex(toParse)
		t.Logf("loc: %v\n", loc)
		if loc == nil {
			parsed = parsed + toParse
			break
		}
		if len(loc) < 6 {
			// matched md url but not submatches. TODO: log this
			parsed = parsed + toParse[loc[0]:loc[1]]
			toParse = toParse[loc[1]:]
			continue
		}
		// matched url and submatches
		beforeMatch := toParse[:loc[0]]
		linkText := toParse[loc[2]:loc[3]]
		linkUrl := toParse[loc[4]:loc[5]]

		parsed = parsed + fmt.Sprintf("%s[%s](%s)", beforeMatch, linkText, linkUrl)
		toParse = toParse[loc[1]:]
	}
	if parsed != sampleMarkdown {
		t.Errorf("Failed to parse md. Got: %s\nExpected:%s", parsed, sampleMarkdown)
	}
}
