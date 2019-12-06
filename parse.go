package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents the link(<a href="..."/>) tag in an
// HTML document
type Link struct {
	Href string
	Text string
}

// Parse wll take an HTML document and return the
// links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data

	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}

	fmt.Println(padding, msg)

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
