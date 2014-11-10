package microformat2

import (
	"golang.org/x/net/html"
	"strings"
	"testing"
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
	if r.Items == nil {
		t.Fatal("items must be not nil.")
	}
	if len(r.Items) != 0 {
		t.Error("Expected no items, Actual ", len(r.Items))
	}
	if r.Rels == nil {
		t.Fatal("rels must be not nil.")
	}
	if len(r.Rels) != 0 {
		t.Error("Expected no rels, Actual ", len(r.Rels))
	}
	if r.Alternates == nil {
		t.Fatal("alternates must be not nil.")
	}
	if len(r.Alternates) != 0 {
		t.Error("Expected no alternates, Actual ", len(r.Alternates))
	}
}

func TestSimpleHCard(t *testing.T) {
	s := `<a class='h-card' href='http://www.example.com/Someone'>Someone</a>`
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
	if len(r.Items) != 1 {
		t.Error("Expected 1 items, Actual ", len(r.Items))
	}
	h_card := r.Items[0]
	if h_card.Types == nil {
		t.Fatal("types property must not be null.")
	}
	if len(h_card.Types) != 1 {
		t.Fatal("Expected one type h-card, Actual ", len(h_card.Types), ": ", h_card.Types)
	}
	if h_card.Types[0] != "h-card" {
		t.Fatal("Expected type 'h-card', Actual '", h_card.Types[0], "'")
	}
	if h_card.Properties == nil {
		t.Fatal("properties property must not be null.")
	}
	if len(h_card.Properties) != 2 {
		t.Fatal("Expected two properties, Actual ", len(h_card.Properties), ": ", h_card.Properties)
	}
	urls := h_card.Properties["url"]
	if urls == nil {
		t.Fatal("url must not be null")
	}
	if urls[0] != "http://www.example.com/Someone" {
		t.Fatal("url must be 'http://www.example.com/Someone'")
	}
	names := h_card.Properties["name"]
	if names == nil {
		t.Fatal("name must not be null")
	}
	if names[0] != "Someone" {
		t.Fatal("Expected name 'Someone', Actual '", names[0], "'")
	}
}

func TestDivAHCard(t *testing.T) {
	s := `<div class='h-card'><a class='u-url p-name' href='http://www.example.com/Someone'>Someone</a></div>`
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
	if len(r.Items) != 1 {
		t.Error("Expected 1 items, Actual ", len(r.Items))
	}
	h_card := r.Items[0]
	if h_card.Types == nil {
		t.Fatal("types property must not be null.")
	}
	if len(h_card.Types) != 1 {
		t.Fatal("Expected one type h-card, Actual ", len(h_card.Types), ": ", h_card.Types)
	}
	if h_card.Types[0] != "h-card" {
		t.Fatal("Expected type 'h-card', Actual '", h_card.Types[0], "'")
	}
	if h_card.Properties == nil {
		t.Fatal("properties property must not be null.")
	}
	if len(h_card.Properties) != 2 {
		t.Fatal("Expected two properties, Actual ", len(h_card.Properties), ": ", h_card.Properties)
	}
	urls := h_card.Properties["url"]
	if urls == nil {
		t.Fatal("url must not be null: ", h_card.Properties)
	}
	if urls[0] != "http://www.example.com/Someone" {
		t.Fatal("url must be 'http://www.example.com/Someone'", h_card.Properties)
	}
	names := h_card.Properties["name"]
	if names == nil {
		t.Fatal("name must not be null")
	}
	if names[0] != "Someone" {
		t.Fatal("Expected name 'Someone', Actual '", names[0], "'")
	}
}
