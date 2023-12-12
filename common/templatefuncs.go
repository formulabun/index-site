package common

import (
	"net/url"
	"text/template"

	"go.formulabun.club/storage"
)

var FuncMap = template.FuncMap{
	"ToUrl": toUrl,
}

func toUrl(item storage.Storable) string {
	p, _ := url.JoinPath(FilePath, item.ToKey().String())
	return p
}
