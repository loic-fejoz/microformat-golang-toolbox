package microformat2

import (
	"fmt"
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

func TestDivAHCardWithAnchorOrgWithUID(t *testing.T) {
	s := `<div class='h-card'><a class='u-url p-name' href='http://www.example.com/Someone'>Someone</a><a class='p-org u-uid' href="http://www.example.com/MyAcme">MyAcme</a></div>`
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
	uids := theOrg.Properties["uid"]
	if len(uids) != 1 {
		t.Fatal("Expected one uid, Actual %s\n", len(uids))
	}
	uid := uids[0]
	if uid != "http://www.example.com/MyAcme" {
		t.Fatal("Expected http://www.example.com/MyAcme, Actual %s\n", uid)
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

func TestealUseCaseFromNybiCCWiki(t *testing.T) {
	s := `<div style="background: none repeat scroll 0 0 #f9f9f9;border: 1px solid #aaa;clear: right;float: right;font-size: 0.9em;line-height: 1.4em;margin: 0 0 0.5em 1em;max-width: 325px;padding: 5px;width: 25em;word-wrap: break-word;" class="h-card">
<p style="background-color: #dfedff;border: 2px solid #dfedff;font-size: 1.4em;font-weight: bold;line-height: 1.1em;margin: 0 0 10px !important;padding: 3px 0;text-align: center;">

<a href="http://loic.fejoz.net" class="p-name u-url ">Loïc Fejoz</a></p>
<p style="margin: 0 0 10px !important;padding: 3px 0;text-align: center;">

<img alt="me" src="http://www.gravatar.com/avatar/00000000000000000000000000000000?d=mm&amp;f=y" class="u-photo "></p>
<dl><dt>Pays
</dt><dd> <span class="p-country-name">France</span>
</dd><dt>FabLab
</dt><dd> 
</dd></dl>
<p><a title="Nancy Bidouille Création Construction" href="http://nybi.cc" class="p-org u-url ">Nybi.cc</a>
</p>
<dl><dt>Email
</dt><dd> loic__AT__fejoz_DOT_net
</dd><dt>Web
</dt><dd> <a href="http://www.fejoz.net" class="external free" rel="nofollow">http://www.fejoz.net</a>
</dd></dl>
<p><a href="http://makezine.com/day-of-making/"><img width="208" height="208" src="/images/f/f7/Day-of-making-badge.png" alt="Day-of-making-badge.png"></a>
</p>
</div>`
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
	if urls[0] != "http://loic.fejoz.net" {
		t.Fatal("url must be 'http://loic.fejoz.net'", hCard.Properties)
	}
	names := hCard.Properties["name"]
	if names == nil {
		t.Fatal("name must not be null")
	}
	if names[0] != "Loïc Fejoz" {
		t.Fatal("Expected name 'Loïc Fejoz', Actual '", names[0], "'")
	}
	orgs := hCard.Properties["org"]
	if orgs == nil {
		t.Fatal("org must not be null: ", orgs)
	}
	if len(orgs) != 1 {
		t.Fatal("org must be unique: ", orgs)
	}
	theOrg := orgs[0].(*Element)
	if theOrg == nil {
		t.Fatal("hCard expected for the organisation")
	}
	fmt.Printf("%s\n", theOrg)
	theOrgName := theOrg.Properties["name"][0]
	if theOrgName != "Nybi.cc" {
		t.Fatal("Expected Nybi.cc, Actual %s\n", theOrgName)
	}
	theOrgUrl := theOrg.Properties["url"][0]
	if theOrgUrl != "http://nybi.cc" {
		t.Fatal("Expected http://nybi.cc, Actual %s\n", theOrgUrl)
	}
}
