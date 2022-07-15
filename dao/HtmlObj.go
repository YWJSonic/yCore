package dao

import "golang.org/x/net/html"

type FilterObj struct {
	FiltAttrs []html.Attribute
	Res       []html.Token
}
