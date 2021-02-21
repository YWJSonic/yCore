package web

import "golang.org/x/net/html"

type HtmlTree struct {
	Root *html.Node
	url  string
}

func (self *HtmlTree) Url() string {
	return self.url
}
