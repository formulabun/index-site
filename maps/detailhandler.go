package maps

import (
	"html/template"
	"log"
	"net/http"

	"go.formulabun.club/metadatadb"
)

type MapHandler struct {
	t *template.Template
	c *metadatadb.Client
}

func mapHandler(dbc *metadatadb.Client) http.Handler {
	t := template.Must(template.New("map.tmpl.html").Funcs(templateFuncs).ParseFiles("./templates/map.tmpl.html", "./templates/maptitle.tmpl"))
	return &MapHandler{
		t,
		dbc,
	}
}

func (f *MapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mapID := r.URL.Query().Get("mapID")
	levelName := r.URL.Query().Get("levelName")
	subTitle := r.URL.Query().Get("subTitle")

	key := metadatadb.MapKey{mapID, levelName, subTitle}
	maps, err := f.c.FindMaps("", &key, r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(maps) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	f.t.Execute(w, maps)
}
