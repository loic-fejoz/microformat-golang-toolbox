package microformat2

import (
	"testing"
	"golang.org/x/net/html"
	"strings"
)

func TestEmptyParagraph(t *testing.T) {
	s := `<p></p>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Parse(doc)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("No data returned.")
		return
	}
	if r.items == nil {
		t.Fatal("items must be not nil.")
	}
	if len(r.items) != 0 {
		t.Error("Expected no items, Actual ", len(r.items))
	}
	if r.rels == nil {
		t.Fatal("rels must be not nil.")
	}
	if len(r.rels) != 0 {
		t.Error("Expected no rels, Actual ", len(r.rels))
	}
	if r.alternates == nil {
		t.Fatal("alternates must be not nil.")
	}
	if len(r.alternates) != 0 {
		t.Error("Expected no alternates, Actual ", len(r.alternates))
	}
}

func TestSimpleHCard(t *testing.T) {
	s := `<a class='h-card' href='http://www.example.com/Someone'>Someone</>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		t.Fatal(err)
	}
	r, err := Parse(doc)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("No data returned.")
	}
	if len(r.items) != 1 {
		t.Error("Expected 1 items, Actual ", len(r.items))
	}
}
