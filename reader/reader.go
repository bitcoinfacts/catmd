package reader

import (
	"bytes"
	"fmt"
	"strings"

	"catmd/reader/jekyll"
)

// Book represents the Glacier PDF which has multiple sections,
// each section containing multiple articles
type Book []Section

type Section struct {
	Title    string
	Articles []jekyll.Article
}

// Read takes the table of contents yaml file used in the
// Glacier docs, imports and cleans all the markdown files
// it mentions
func Read(tocFilepath, basedir string) Book {
	toc := newToc(tocFilepath, basedir, MarkdownExtension)
	book := importBook(toc)
	return book
}

// importBook imports all markdown articles referenced in
// the table of contents file and returns a Book with them
func importBook(toc Toc) Book {
	book := Book{}
	for _, tocSection := range toc {
		book = append(book, importSection(tocSection))
	}
	return book
}

// importSection imports all markdown articles referenced in
// a section of the table of contents file, and returns a Section
// with them
func importSection(tocSection TocSection) Section {
	section := Section{tocSection.Title, []jekyll.Article{}}
	articles := []jekyll.Article{}

	for _, filepath := range tocSection.Path {
		article, err := jekyll.New(filepath)
		if err == nil {
			articles = append(articles, article)
		} else {
			fmt.Print(err)
		}
	}
	section.Articles = articles
	return section
}

func (s *Section) String() string {
	buf := bytes.Buffer{}
	sectionFormat := "# %s\n\n"

	buf.WriteString(fmt.Sprintf(sectionFormat, s.Title))
	for i, article := range s.Articles {
		buf.WriteString(article.PrintBody())
		if i < len(s.Articles)-1 {
			buf.WriteString("\n\n")
		}
	}
	return buf.String()
}

func (b *Book) String() string {
	const jekyllFrontMatter = `---
layout: pdf
permalink: /pdf.html
---`
	buf := bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("%s\n\n", jekyllFrontMatter))

	buf.WriteString(fmt.Sprintf("%s\n\n", b.CreateToc()))

	for i, section := range *b {
		buf.WriteString(section.String())
		if i < len(*b)-1 {
			buf.WriteString("\n\n")
		}
	}
	return buf.String()
}

func (b *Book) CreateToc() string {
	const sectionLinkFormat = "* [%s](#%s)\n"
	const articleLinkFormat = "  * [%s](#%s)\n"

	buf := bytes.Buffer{}
	for _, section := range *b {
		sectionFragment := strings.Replace(section.Title, " ", "-", -1)
		sectionFragment = strings.ToLower(sectionFragment)
		buf.WriteString(fmt.Sprintf(sectionLinkFormat, section.Title, sectionFragment))

		for _, article := range section.Articles {
			articleFragment := strings.Replace(article.Title, " ", "-", -1)
			articleFragment = strings.ToLower(articleFragment)
			buf.WriteString(fmt.Sprintf(articleLinkFormat, article.Title, articleFragment))
		}
	}
	return buf.String()
}

// MapPathToTitle returns a mapping of filepath of an article to its title
func (b *Book) MapPathToTitle() map[string]string {
	m := map[string]string{}
	for _, section := range *b {
		for _, article := range section.Articles {
			m[article.Path] = article.Title
		}
	}
	return m
}
