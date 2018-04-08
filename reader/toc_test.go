package reader

import (
	"strings"
	"testing"
)

const (
	testTocFile = "testdata/toc.yaml"
	baseDir     = "_docs"
	extension   = ".md"
)

var expected = []string{
	"_docs/overview.md",
	"_docs/overview/key-concepts.md",
	"_docs/setup/verify.md",
	"_docs/setup/non-quarantined-hardware.md",
	"_docs/setup/quarantined-hardware.md",
}

func TestReadTocFile(t *testing.T) {
	const errorMsg = "Failed to parse order file:\nExpected: %v\nGot: %v"

	r := readTocFile(testTocFile, baseDir, extension)
	if len(r) != len(expected) {
		t.Errorf(errorMsg, expected, r)
	}
	for i := 0; i < len(r); i++ {
		if strings.Compare(r[i], expected[i]) != 0 {
			t.Errorf(errorMsg, expected, r)
		}
	}
}
