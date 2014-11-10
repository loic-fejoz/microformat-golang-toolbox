package microformat2

import (
	"golang.org/x/net/html"
	"strings"
)

type Element struct {
	types []string
	properties map[string]interface{}
}

type Result struct {
	items []*Element
	rels map[string]string
	alternates []interface{}
}

func New() (res *Result) {
	res = &Result{}
	res.items = make([]*Element, 0, 5)
	res.rels = map[string]string{}
	res.alternates = make([]interface{}, 0, 5)
	return
}

func getClasses(a_node *html.Node) []string {
	for _, a := range a_node.Attr {
		if a.Key == "class" {
			return strings.Split(a.Val, " ")
		}
	}
	return make([]string, 0, 0)
}

func filterRootEntry(classes []string) []string {
	out := make([]string, 0 , len(classes))
	for _, a_class := range classes {
		if strings.HasPrefix(a_class, "h-") {
			out = append(out, a_class)
		}
	}
	return out
}

func isRoot(classes []string) bool {
	for _, a_class := range classes {
		if strings.HasPrefix(a_class, "h-") {
			return true;
		}
	}
	return false
}

func AccParse(a_node *html.Node, result *Result, root *Element) {
	current_root := root
	if a_node.Type == html.ElementNode {
		all_classes := getClasses(a_node)
		if isRoot(all_classes) {
			types := filterRootEntry(all_classes)
			current_root = &Element{types: types}
			if root == nil {
				result.items = append(result.items, current_root)
			}
		}
	}
	// Visit children
	for c := a_node.FirstChild; c != nil; c = c.NextSibling {
		AccParse(c, result, current_root)
	}
}

func Parse(doc *html.Node) (result *Result, err error) {
	result = New()
	err = nil
	AccParse(doc, result, nil)
	return
}
