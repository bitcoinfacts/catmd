package urlparser

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func TransformUrl(rawurl string, articleTitle map[string]string) (string, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	// External urls should be left unchanged
	if u.IsAbs() {
		return rawurl, nil
	}
	// Relative urls with fragments, should be turned into fragments
	if u.Fragment != "" {
		return fmt.Sprintf("#%s", u.Fragment), nil
	}
	// Relative urls w/o fragments should be turned to fragment, based on article title
	path := strings.TrimPrefix(u.EscapedPath(), "/")
	title := articleTitle[path]
	if title == "" {
		return rawurl, errors.New(fmt.Sprintf("Could not find matching title for %s", rawurl))
	}
	title = strings.Replace(title, " ", "-", -1)
	return fmt.Sprintf("#%s", title), nil
}
