package jekyll

import (
	"testing"
)

const testfile = "testdata/about.md"

func TestNewJekyllArticle(t *testing.T) {
	article, err := New(testfile)
	if err != nil {
		t.Errorf("Failed to create jekyll article from %s", testfile)
	}
	if !sameArticle.CompareTo(&article) {
		t.Errorf("Failed to import jekyll article.\nExpected: %v\nGot: %v\n", sameArticle, article)
	}
	for _, differentArticle := range differentArticles() {
		if differentArticle.CompareTo(&article) {
			t.Errorf("Failed to import jekyll article.\nExpected: %v\nGot: %v\n", sameArticle, article)
		}
	}
}

func differentArticles() []Article {
	return []Article{
		articleDifferentPath,
		articleDifferentTitle,
		articleDifferentDescription,
		articleDifferentBody}
}

// TODO: this should be extracted since the jekyll package should
// import content without changing it
// func TestTransformUrls(t *testing.T) {
// 	t.Skip("Skipping transform url testing since it needs refactoring")
// 	for testfile, expectedfile := range testfiles {
// 		testfile = path.Join(testDataPath, testfile)
// 		expectedfile = path.Join(testDataPath, expectedfile)
//
// 		buf, err := ioutil.ReadFile(expectedfile)
// 		if err != nil {
// 			t.Error(err)
// 		}
// 		expectedMarkdown := string(buf)
//
// 		article, err := New(testfile)
// 		if err != nil {
// 			t.Errorf("Failed to create jekyll article from %s", testfile)
// 		}
// 		gotMarkdown := article.TransformUrls(identity)
// 		if expectedMarkdown != gotMarkdown {
// 			t.Errorf(`Failed to transform urls.
// Expected: %v
// Got: %v`, gotMarkdown, expectedMarkdown)
// 		}
// 	}
// }
//
// func identity(str string) string {
// 	return str
// }
