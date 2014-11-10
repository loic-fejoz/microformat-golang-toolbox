package microformat2

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

func accText(a_node *html.Node, buff *bytes.Buffer) {
	if a_node.Type == html.TextNode {
		buff.Write([]byte(a_node.Data))
	}
	// Visit children
	for c := a_node.FirstChild; c != nil; c = c.NextSibling {
		accText(c, buff)
	}
}

func getText(a_node *html.Node) string {
	var b bytes.Buffer
	accText(a_node, &b)
	return b.String()
}

type Element struct {
	Types      []string
	Properties map[string][]interface{}
	isHCard    bool
}

func NewElement(types []string) *Element {
	result := &Element{Types: types, Properties: make(map[string][]interface{})}
	for _, a_type := range types {
		if a_type == "h-card" {
			result.isHCard = true
		}
	}
	return result
}

func AppendProperty(elt *Element, propName string, propVal interface{}) {
	props := elt.Properties[propName]
	if props == nil {
		props = make([]interface{}, 0, 3)
	}
	elt.Properties[propName] = append(props, propVal)
}

type Result struct {
	Items      []*Element
	Rels       map[string]string
	Alternates []interface{}
}

func New() (res *Result) {
	res = &Result{}
	res.Items = make([]*Element, 0, 5)
	res.Rels = map[string]string{}
	res.Alternates = make([]interface{}, 0, 5)
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

func filterRootClass(classes []string) []string {
	out := make([]string, 0, len(classes))
	for _, a_class := range classes {
		if strings.HasPrefix(a_class, "h-") {
			out = append(out, a_class)
		}
	}
	return out
}

func filterPropertyClass(classes []string) []string {
	out := make([]string, 0, len(classes))
	for _, a_class := range classes {
		if strings.HasPrefix(a_class, "p-") || strings.HasPrefix(a_class, "u-") || strings.HasPrefix(a_class, "dt-") || strings.HasPrefix(a_class, "e-") {
			out = append(out, a_class)
		}
	}
	return out
}

func isRoot(classes []string) bool {
	for _, a_class := range classes {
		if strings.HasPrefix(a_class, "h-") {
			return true
		}
	}
	return false
}

func AccParse(a_node *html.Node, result *Result, root *Element) {
	current_root := root
	if a_node.Type == html.ElementNode {
		all_classes := getClasses(a_node)
		if isRoot(all_classes) {
			types := filterRootClass(all_classes)
			current_root = NewElement(types)
			if root == nil {
				result.Items = append(result.Items, current_root)
			}

			// <a class="h-card..." href="...">...</a>
			if (root == current_root || root == nil) && a_node.Data == "a" && current_root.isHCard {
				for _, a := range a_node.Attr {
					if a.Key == "href" {
						AppendProperty(current_root, "url", a.Val)
					}
				}
				AppendProperty(current_root, "name", getText(a_node))
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
