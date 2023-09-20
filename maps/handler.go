package maps

import (
	"html/template"
	"log"
	"net/http"

	"go.formulabun.club/metadatadb"
	srb2kartstrings "go.formulabun.club/srb2kart/strings"
)

type MapsHandler struct {
	c *metadatadb.Client
	t *template.Template
}

func mapsHandler(dbClient *metadatadb.Client) http.Handler {
	funcs := template.FuncMap{
		"removeColorCodes": srb2kartstrings.RemoveColorCodes,
	}

	t := template.Must(
		template.New("mapslist.tmpl.html").Funcs(funcs).ParseFiles("templates/mapslist.tmpl.html", "templates/maptitle.tmpl"),
	)

	return &MapsHandler{
		dbClient,
		t,
	}
}

func (f *MapsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")

	files, err := f.c.FindMaps(filename, r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = f.t.Execute(w, files)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
