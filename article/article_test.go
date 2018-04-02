package article

import (
	"fmt"
	"testing"
)

const (
	input = `---
title: About Glacier
description: Glacier description
redirect_from:
  - /docs/
---

Glacier is a step-by-step protocol for storing bitcoins in a highly secure
manner. It is intended for:

* **Personal storage**: Glacier does not address institutional security
needs such as internal controls, transparent auditing, and preventing access
to funds by a single individual.`

	expected = `# About Glacier

Glacier is a step-by-step protocol for storing bitcoins in a highly secure
manner. It is intended for:

* **Personal storage**: Glacier does not address institutional security
needs such as internal controls, transparent auditing, and preventing access
to funds by a single individual.`
)

func TestStringer(t *testing.T) {
	a := New(input)
	str := fmt.Sprintf("%v", a)
	if str != expected {
		t.Errorf("Failed to remove frontmatter from markdown file.\n Expected:\n%s\n\nGot:\n%s\n", expected, str)
	}
}
