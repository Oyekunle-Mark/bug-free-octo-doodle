package link

import "io"

// Link represents the link(<a href="..."/>) tag in an
// HTML document
type Link struct {
	Href string
	Text string
}

// Parse wll take an HTML document and return the
// links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
