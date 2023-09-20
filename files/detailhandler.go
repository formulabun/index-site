package files

import (
	"html/template"
	"net/http"

	"go.formulabun.club/metadatadb"
	"go.formulabun.club/srb2kart/addons"
)

type FileHandler struct {
	t *template.Template
	c *metadatadb.Client
}

func fileHandler(dbc *metadatadb.Client) http.Handler {
	funcs := template.FuncMap{
		"isRace": func(file string) bool {
			return addons.GetAddonType(file)&addons.RaceFlag != 0
		},
	}
	t := template.Must(template.New("file.tmpl.html").Funcs(funcs).ParseFiles("./templates/file.tmpl.html"))
	return &FileHandler{
		t,
		dbc,
	}
}

func (f *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("id")
	f.t.Execute(w, filename)
}
