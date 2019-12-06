package link

import (
	"io"
	"strings"

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

	// dfs(doc, "")

	nodes := findLinkNodes(doc)

	var links []Link

	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = parseText(n)

	return ret
}

func parseText(n *html.Node) string {
	if n.Type == html.TextNode {
		return strings.TrimSpace(n.Data)
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += parseText(c) + " "
	}

	return strings.TrimSpace(ret)
}

func findLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, findLinkNodes(c)...)
	}

	return ret
}

// func dfs(n *html.Node, padding string) {
// 	msg := n.Data

// 	if n.Type == html.ElementNode {
// 		msg = "<" + msg + ">"
// 	}

// 	fmt.Println(padding, msg)

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}
// }
