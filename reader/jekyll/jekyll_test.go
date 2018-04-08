package jekyll

import (
	"io/ioutil"
	"path"
	"testing"
)

const testDataPath = "testdata"

var testfiles = map[string]string{
	"index-input.md": "index-expected.md",
}

func TestNewJekyllArticle(t *testing.T) {
	for testfile, expectedfile := range testfiles {
		testfile = path.Join(testDataPath, testfile)
		expectedfile = path.Join(testDataPath, expectedfile)

		buf, err := ioutil.ReadFile(expectedfile)
		if err != nil {
			t.Error(err)
		}
		expectedMarkdown := string(buf)

		article, err := New(testfile)
		if err != nil {
			t.Errorf("Failed to create jekyll article from %s", testfile)
		}

		gotMarkdown := article.RawBody

		if expectedMarkdown != gotMarkdown {
			t.Errorf(`Failed to remove frontmatter from markdown file.
Expected: %v
Got: %v`, gotMarkdown, expectedMarkdown)
		}
	}
}

func TestTransformUrls(t *testing.T) {
	for testfile, expectedfile := range testfiles {
		testfile = path.Join(testDataPath, testfile)
		expectedfile = path.Join(testDataPath, expectedfile)

		buf, err := ioutil.ReadFile(expectedfile)
		if err != nil {
			t.Error(err)
		}
		expectedMarkdown := string(buf)

		article, err := New(testfile)
		if err != nil {
			t.Errorf("Failed to create jekyll article from %s", testfile)
		}
		gotMarkdown := article.TransformUrls(identity)
		if expectedMarkdown != gotMarkdown {
			t.Errorf(`Failed to transform urls.
Expected: %v
Got: %v`, gotMarkdown, expectedMarkdown)
		}
	}
}

func identity(str string) string {
	return str
}
