package microformat2

import (
//	"fmt"
)

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

func Append(elt1 *Element, elt2 *Element) *Element {
	elt1.Types = append(elt1.Types, elt2.Types...)
	elt1.isHCard = elt1.isHCard || elt2.isHCard
	for key, vals := range elt2.Properties {
		for _, v := range vals {
			AppendProperty(elt1, key, v)
		}
	}
//	fmt.Printf("shall merge ", elt1, " and ", elt2)
	return elt1
}

type Result struct {
	Items      []*Element
	Rels       map[string]string
	Alternates []interface{}
}

func NewResult() (res *Result) {
	res = &Result{}
	res.Items = make([]*Element, 0, 5)
	res.Rels = map[string]string{}
	res.Alternates = make([]interface{}, 0, 5)
	return
}
