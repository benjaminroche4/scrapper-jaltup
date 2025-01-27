package html

import (
	"strings"

	"golang.org/x/net/html"
)

func NextNodeElement(node *html.Node) *html.Node {
	next := node
	for next != nil {
		next = next.NextSibling
		if next != nil && next.Type == html.ElementNode {
			return next
		}
	}

	return nil
}

func NextChildNodeElement(node *html.Node) *html.Node {
	if node == nil {
		return nil
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {
			return child
		}
	}

	return nil
}

func FindNodeByTagName(node *html.Node, name string) *html.Node {
	if node == nil {
		return nil
	}
	if node.Type == html.ElementNode && node.Data == name {
		return node
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		found := FindNodeByTagName(child, name)
		if found != nil {
			return found
		}
	}

	return nil
}

func FindAllNodesByTagName(node *html.Node, name string) []*html.Node {
	nodes := []*html.Node{}

	if node == nil {
		return nodes
	}

	if node.Type == html.ElementNode && node.Data == name {
		return []*html.Node{node}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		found := FindAllNodesByTagName(child, name)
		if len(found) > 0 {
			nodes = append(nodes, found...)
		}
	}

	return nodes
}

func FindNodeByClassName(node *html.Node, name string) *html.Node {
	if node != nil {
		if node.Type == html.ElementNode {
			for _, a := range node.Attr {
				if a.Key == "class" {
					classes := strings.Split(a.Val, " ")
					for _, class := range classes {
						if class == name {
							return node
						}
					}
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			found := FindNodeByClassName(child, name)
			if found != nil {
				return found
			}
		}
	}

	return nil
}

func GetNodeAttr(node *html.Node, key string) (string, bool) {
	if node != nil {
		for _, a := range node.Attr {
			if a.Key == key {
				return a.Val, true
			}
		}
	}
	return "", false
}

func GetTextContent(node *html.Node) (string, bool) {
	if node != nil {
		data := strings.TrimSpace(node.Data)
		if node.Type == html.TextNode && data != "" {
			return data, true
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			data, found := GetTextContent(child)
			if found {
				return data, true
			}
		}
	}
	return "", false
}

func GetNodeLink(node *html.Node) (string, bool) {
	if node != nil {
		if node.Type == html.ElementNode && node.Data == "a" {
			return GetNodeAttr(node, "href")
		}
	}

	return "", false
}
