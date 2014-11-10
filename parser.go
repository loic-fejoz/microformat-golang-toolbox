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

func mapProp2Name(prop string) string {
	i := strings.Index(prop, "-") + 1
	return prop[i:len(prop)]
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
		for _, prop := range filterPropertyClass(all_classes) {
			if strings.HasPrefix(prop, "p-") || strings.HasPrefix(prop, "dt-") {
				AppendProperty(current_root, mapProp2Name(prop), getText(a_node))
			} else if strings.HasPrefix(prop, "u-") {
				for _, a := range a_node.Attr {
					if a.Key == "href" {
						AppendProperty(current_root, mapProp2Name(prop), a.Val)
					}
				}
			} else if strings.HasPrefix(prop, "e-") {
				//	AppendProperty(current_root, prop, getHtml(a_node))
			}
		}
	}
	// Visit children
	for c := a_node.FirstChild; c != nil; c = c.NextSibling {
		AccParse(c, result, current_root)
	}
}

func Parse(doc *html.Node) (result *Result, err error) {
	result = NewResult()
	err = nil
	AccParse(doc, result, nil)
	return
}
