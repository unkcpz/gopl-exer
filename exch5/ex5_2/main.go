// populate a mapping from element names to number of elements in HTML tree
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	eleCounts := make(map[string]int)
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5_2: %v\n", err)
		os.Exit(1)
	}
	eleCounts = mapElement(eleCounts, doc)

	for ele, c := range eleCounts {
		fmt.Printf("%4d\t%s\n", c, ele)
	}
}

// mapping the HTML tree element to population counts
func mapElement(ec map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		ec[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ec = mapElement(ec, c)
	}
	return ec
}
