package jekyll

// Used for testing
var (
	sameArticle Article = Article{
		Path:        "testdata/about.md",
		Title:       "About Glacier",
		Description: "Glacier description",
		RawBody: `Glacier is a step-by-step protocol for storing bitcoins in a highly secure
manner. It is intended for:

* [test url](example)
* [one more](http://example.org)

And a paragraph [containing another url](/example) in the middle of a sentence.

## Background

### Self-Managed Storage vs. Online

Let's start by assessing whether Glacier is right for you.`,
	}

	articleDifferentPath Article = Article{
		Path:        "different-path.md",
		Title:       "About Glacier",
		Description: "Glacier description",
		RawBody: `Glacier is a step-by-step protocol for storing bitcoins in a highly secure
manner. It is intended for:

* [test url](example)
* [one more](http://example.org)

And a paragraph [containing another url](/example) in the middle of a sentence.

## Background

### Self-Managed Storage vs. Online

Let's start by assessing whether Glacier is right for you.`,
	}

	articleDifferentTitle Article = Article{
		Path:        "testdata/about.md",
		Title:       "A different title",
		Description: "Glacier description",
		RawBody: `Glacier is a step-by-step protocol for storing bitcoins in a highly secure
manner. It is intended for:

* [test url](example)
* [one more](http://example.org)

And a paragraph [containing another url](/example) in the middle of a sentence.

## Background

### Self-Managed Storage vs. Online

Let's start by assessing whether Glacier is right for you.`,
	}

	articleDifferentDescription Article = Article{
		Path:        "testdata/about.md",
		Title:       "About Glacier",
		Description: "A different description",
		RawBody: `Glacier is a step-by-step protocol for storing bitcoins in a highly secure
manner. It is intended for:

* [test url](example)
* [one more](http://example.org)

And a paragraph [containing another url](/example) in the middle of a sentence.

## Background

### Self-Managed Storage vs. Online

Let's start by assessing whether Glacier is right for you.`,
	}
	articleDifferentBody Article = Article{
		Path:        "testdata/about.md",
		Title:       "About Glacier",
		Description: "Glacier description",
		RawBody:     `This content should be different`,
	}
)
