package reader

import (
	"testing"
)

const (
	testTocFile = "testdata/toc.yaml"
	baseDir     = "_docs"
	extension   = ".md"
)

var expected = []TocSection{
	TocSection{
		"Glacier overview",
		[]string{"_docs/overview.md", "_docs/overview/key-concepts.md"},
	},
	TocSection{
		"Before you start",
		[]string{},
	},
	TocSection{
		"Setup",
		[]string{"_docs/setup/verify.md", "_docs/setup/non-quarantined-hardware.md", "_docs/setup/quarantined-hardware.md"},
	},
}

func TestNewToc(t *testing.T) {
	const errorMsg = "Failed to parse order file:\nExpected: %v\nGot: %v"

	tocGot := newToc(testTocFile, baseDir, extension)
	if len(tocGot) != len(expected) {
		t.Errorf(errorMsg, expected, tocGot)
	}
	for i := 0; i < len(tocGot); i++ {
		if tocGot[i].Title != expected[i].Title {
			t.Errorf(errorMsg, expected, tocGot)
		}
		for j := 0; j < len(tocGot[i].Path); j++ {
			if tocGot[i].Path[j] != expected[i].Path[j] {
				t.Errorf(errorMsg, expected, tocGot)
			}
		}
	}

}
