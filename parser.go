package microformat2

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

func accText(aNode *html.Node, buff *bytes.Buffer) {
	if aNode.Type == html.TextNode {
		buff.Write([]byte(aNode.Data))
	}
	// Visit children
	for c := aNode.FirstChild; c != nil; c = c.NextSibling {
		accText(c, buff)
	}
}

func getText(aNode *html.Node) string {
	var b bytes.Buffer
	accText(aNode, &b)
	return b.String()
}

func getClasses(aNode *html.Node) []string {
	for _, a := range aNode.Attr {
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

func isRoot(aNode *html.Node, classes []string) bool {
	for _, a_class := range classes {
		if strings.HasPrefix(a_class, "h-") {
			return true
		}
		if (a_class == "p-org") && (aNode.Type == html.ElementNode) && (aNode.Data == "a") {
			return true
		}
	}
	return false
}

func mapProp2Name(prop string) string {
	i := strings.Index(prop, "-") + 1
	return prop[i:len(prop)]
}

func AccParse(aNode *html.Node, result *Result, root *Element) {
	currentRoot := root
	if aNode.Type == html.ElementNode {
		allClasses := getClasses(aNode)
		if isRoot(aNode, allClasses) {
			types := filterRootClass(allClasses)
			currentRoot = NewElement(types)
			if root == nil {
				result.Items = append(result.Items, currentRoot)
			}

			// <a class="h-card..." href="...">...</a>
			if aNode.Data == "a" && currentRoot.isHCard {
				for _, a := range aNode.Attr {
					if a.Key == "href" {
						AppendProperty(currentRoot, "url", a.Val)
					}
				}
				AppendProperty(currentRoot, "name", getText(aNode))
			}
		}
		for _, prop := range filterPropertyClass(allClasses) {
			if strings.HasPrefix(prop, "p-") || strings.HasPrefix(prop, "dt-") {
				if prop == "p-org" && currentRoot != root {
					AppendProperty(root, mapProp2Name(prop), currentRoot)
					AppendProperty(currentRoot, "name", getText(aNode))
				} else {
					AppendProperty(currentRoot, mapProp2Name(prop), getText(aNode))
				}
			} else if strings.HasPrefix(prop, "u-") {
				for _, a := range aNode.Attr {
					if a.Key == "href" {
						AppendProperty(currentRoot, mapProp2Name(prop), a.Val)
					}
				}
			} else if strings.HasPrefix(prop, "e-") {
				//	AppendProperty(currentRoot, prop, getHtml(aNode))
			}
		}
	}
	// Visit children
	for c := aNode.FirstChild; c != nil; c = c.NextSibling {
		AccParse(c, result, currentRoot)
	}
}

func Parse(doc *html.Node) (result *Result, err error) {
	result = NewResult()
	err = nil
	AccParse(doc, result, nil)
	return
}
