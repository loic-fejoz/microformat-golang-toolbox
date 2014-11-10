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
	hCard := r.Items[0]
	if hCard.Types == nil {
		t.Fatal("types property must not be null.")
	}
	if len(hCard.Types) != 1 {
		t.Fatal("Expected one type h-card, Actual ", len(hCard.Types), ": ", hCard.Types)
	}
	if hCard.Types[0] != "h-card" {
		t.Fatal("Expected type 'h-card', Actual '", hCard.Types[0], "'")
	}
	if hCard.Properties == nil {
		t.Fatal("properties property must not be null.")
	}
	if len(hCard.Properties) != 2 {
		t.Fatal("Expected two properties, Actual ", len(hCard.Properties), ": ", hCard.Properties)
	}
	urls := hCard.Properties["url"]
	if urls == nil {
		t.Fatal("url must not be null")
	}
	if urls[0] != "http://www.example.com/Someone" {
		t.Fatal("url must be 'http://www.example.com/Someone'")
	}
	names := hCard.Properties["name"]
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
	hCard := r.Items[0]
	if hCard.Types == nil {
		t.Fatal("types property must not be null.")
	}
	if len(hCard.Types) != 1 {
		t.Fatal("Expected one type h-card, Actual ", len(hCard.Types), ": ", hCard.Types)
	}
	if hCard.Types[0] != "h-card" {
		t.Fatal("Expected type 'h-card', Actual '", hCard.Types[0], "'")
	}
	if hCard.Properties == nil {
		t.Fatal("properties property must not be null.")
	}
	if len(hCard.Properties) != 2 {
		t.Fatal("Expected two properties, Actual ", len(hCard.Properties), ": ", hCard.Properties)
	}
	urls := hCard.Properties["url"]
	if urls == nil {
		t.Fatal("url must not be null: ", hCard.Properties)
	}
	if urls[0] != "http://www.example.com/Someone" {
		t.Fatal("url must be 'http://www.example.com/Someone'", hCard.Properties)
	}
	names := hCard.Properties["name"]
	if names == nil {
		t.Fatal("name must not be null")
	}
	if names[0] != "Someone" {
		t.Fatal("Expected name 'Someone', Actual '", names[0], "'")
	}
}

func TestDivAHCardWithAnchorOrg(t *testing.T) {
	s := `<div class='h-card'><a class='u-url p-name' href='http://www.example.com/Someone'>Someone</a><a class='p-org u-url' href="http://www.example.com/MyAcme">MyAcme</a></div>`
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
	hCard := r.Items[0]
	if hCard.Types == nil {
		t.Fatal("types property must not be null.")
	}
	if len(hCard.Types) != 1 {
		t.Fatal("Expected one type h-card, Actual ", len(hCard.Types), ": ", hCard.Types)
	}
	if hCard.Types[0] != "h-card" {
		t.Fatal("Expected type 'h-card', Actual '", hCard.Types[0], "'")
	}
	if hCard.Properties == nil {
		t.Fatal("properties property must not be null.")
	}
	if len(hCard.Properties) != 3 {
		t.Fatal("Expected three properties, Actual ", len(hCard.Properties), ": ", hCard.Properties)
	}
	urls := hCard.Properties["url"]
	if urls == nil {
		t.Fatal("url must not be null: ", hCard.Properties)
	}
	if len(urls) != 1 {
		t.Fatal("url must be unique: ", urls)
	}
	if urls[0] != "http://www.example.com/Someone" {
		t.Fatal("url must be 'http://www.example.com/Someone'", hCard.Properties)
	}
	names := hCard.Properties["name"]
	if names == nil {
		t.Fatal("name must not be null")
	}
	if names[0] != "Someone" {
		t.Fatal("Expected name 'Someone', Actual '", names[0], "'")
	}
	orgs := hCard.Properties["org"]
	if orgs == nil {
		t.Fatal("org must not be null: ", orgs)
	}
	if len(orgs) != 1 {
		t.Fatal("org must be unique: ", orgs)
	}
	theOrg := orgs[0].(*Element)
	urls = theOrg.Properties["url"]
	if len(urls) != 1 {
		t.Fatal("Expected one url, Actual %s\n", len(urls))
	}
	url := urls[0]
	if url != "http://www.example.com/MyAcme" {
		t.Fatal("Expected http://www.example.com/MyAcme, Actual %s\n", url)
	}
	names = theOrg.Properties["name"]
	if len(names) != 1 {
		t.Fatal("Expected one name, Actual %s\n", len(names))
	}
	name := names[0]
	if name != "MyAcme" {
		t.Fatal("Expected MyAcme, Actual %s\n", name)
	}
}

func TestDivAHCardWithOrg(t *testing.T) {
	s := `<div class='h-card'><a class='u-url p-name' href='http://www.example.com/Someone'>Someone</a><span class='p-org'>MyAcme</span></div>`
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
	hCard := r.Items[0]
	if hCard.Types == nil {
		t.Fatal("types property must not be null.")
	}
	if len(hCard.Types) != 1 {
		t.Fatal("Expected one type h-card, Actual ", len(hCard.Types), ": ", hCard.Types)
	}
	if hCard.Types[0] != "h-card" {
		t.Fatal("Expected type 'h-card', Actual '", hCard.Types[0], "'")
	}
	if hCard.Properties == nil {
		t.Fatal("properties property must not be null.")
	}
	if len(hCard.Properties) != 3 {
		t.Fatal("Expected three properties, Actual ", len(hCard.Properties), ": ", hCard.Properties)
	}
	urls := hCard.Properties["url"]
	if urls == nil {
		t.Fatal("url must not be null: ", hCard.Properties)
	}
	if len(urls) != 1 {
		t.Fatal("url must be unique: ", urls)
	}
	if urls[0] != "http://www.example.com/Someone" {
		t.Fatal("url must be 'http://www.example.com/Someone'", hCard.Properties)
	}
	names := hCard.Properties["name"]
	if names == nil {
		t.Fatal("name must not be null")
	}
	if names[0] != "Someone" {
		t.Fatal("Expected name 'Someone', Actual '", names[0], "'")
	}
	orgs := hCard.Properties["org"]
	if orgs == nil {
		t.Fatal("org must not be null: ", orgs)
	}
	if len(orgs) != 1 {
		t.Fatal("org must be unique: ", orgs)
	}
	theOrgName := orgs[0]
	if theOrgName != "MyAcme" {
		t.Fatal("Expected MyAcme, Actual %s\n", theOrgName)
	}
}
