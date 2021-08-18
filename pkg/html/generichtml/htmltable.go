package generichtml

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/openshift/sippy/pkg/util/sets"
)

type HTMLItem interface {
	ToHTML() string
}

type htmlTextElement string

func (h htmlTextElement) ToHTML() string {
	return string(h)
}

func NewHTMLTextElement(text string) HTMLItem {
	return htmlTextElement(text)
}

func SpaceHTMLItems(htmlItems []HTMLItem) []HTMLItem {
	out := []HTMLItem{}

	for _, item := range htmlItems {
		out = append(out, item, NewHTMLTextElement(" "))
	}

	return out
}

type HTMLElement struct {
	Params    map[string]string
	Text      string
	HTMLItems []HTMLItem
	Element   string
}

func (t HTMLElement) ToHTML() string {
	sb := &strings.Builder{}

	fmt.Fprint(sb, "<"+t.Element)

	// Order param keys
	for _, paramKey := range sets.StringKeySet(t.Params).List() {
		fmt.Fprint(sb, " "+paramKey+"="+`"`)
		fmt.Fprint(sb, t.Params[paramKey]+`"`)
	}

	fmt.Fprint(sb, ">")

	if len(t.HTMLItems) != 0 {
		for _, item := range t.HTMLItems {
			fmt.Fprint(sb, item.ToHTML())
		}
	} else {
		fmt.Fprint(sb, t.Text)
	}

	fmt.Fprint(sb, "</"+t.Element+">")

	return sb.String()
}

func NewHTMLLinkWithParams(text string, linkURL *url.URL, params map[string]string) HTMLElement {
	t := HTMLElement{
		Element: "a",
		Text:    text,
		Params:  map[string]string{},
	}

	for k, v := range params {
		t.Params[k] = v
	}

	_, ok := t.Params["href"]
	if ok {
		return t
	}

	t.Params["href"] = linkURL.String()
	return t
}

func NewHTMLLink(text string, linkURL *url.URL) HTMLElement {
	return NewHTMLLinkWithParams(
		text,
		linkURL,
		map[string]string{
			"href": linkURL.String(),
		})
}

type HTMLTableHeaderRowItem struct {
	Text      string
	HTMLItems []HTMLItem
	Params    map[string]string
}

func (r HTMLTableHeaderRowItem) ToHTML() string {
	t := HTMLElement{
		Element:   "th",
		Params:    r.Params,
		Text:      r.Text,
		HTMLItems: r.HTMLItems,
	}

	return t.ToHTML()
}

type HTMLTableRowItem struct {
	Text      string
	HTMLItems []HTMLItem
	Params    map[string]string
}

func (r HTMLTableRowItem) ToHTML() string {
	t := HTMLElement{
		Element:   "td",
		Params:    r.Params,
		HTMLItems: r.HTMLItems,
		Text:      r.Text,
	}

	return t.ToHTML()
}

type HTMLTableRow struct {
	items  []HTMLItem
	params map[string]string
}

func NewHTMLTableRowWithItems(p map[string]string, items []HTMLItem) HTMLTableRow {
	return HTMLTableRow{
		items:  items,
		params: p,
	}
}

func NewHTMLTableRow(p map[string]string) HTMLTableRow {
	return HTMLTableRow{
		params: p,
	}
}

func (r *HTMLTableRow) AddItems(items []HTMLItem) {
	r.items = append(r.items, items...)
}

func (r HTMLTableRow) ToHTML() string {
	sb := &strings.Builder{}

	fmt.Fprint(sb, "\n  ")

	for _, item := range r.items {
		fmt.Fprint(sb, "  "+item.ToHTML()+"\n  ")
	}

	t := HTMLElement{
		Element: "tr",
		Params:  r.params,
		Text:    sb.String(),
	}

	return "  " + t.ToHTML() + "\n"
}

type HTMLTable struct {
	headerRows []HTMLTableRow
	rows       []HTMLTableRow
	params     map[string]string
}

func NewHTMLTable(p map[string]string) HTMLTable {
	return HTMLTable{
		params: p,
	}
}

func (h *HTMLTable) AddHeaderRow(headerRow HTMLTableRow) {
	h.headerRows = append(h.headerRows, headerRow)
}

func (h *HTMLTable) AddRow(row HTMLTableRow) {
	h.rows = append(h.rows, row)
}

func (h HTMLTable) ToHTML() string {
	sb := &strings.Builder{}

	fmt.Fprintln(sb, "")

	for _, row := range h.headerRows {
		fmt.Fprint(sb, row.ToHTML())
	}

	for _, row := range h.rows {
		fmt.Fprint(sb, row.ToHTML())
	}

	t := HTMLElement{
		Element: "table",
		Params:  h.params,
		Text:    sb.String(),
	}

	return t.ToHTML()
}