package files

import (
	"html/template"
	"log"
	"net/http"

	"go.formulabun.club/index-site/common"
	"go.formulabun.club/metadatadb"
	"go.formulabun.club/srb2kart/addons"
)

type FileHandler struct {
	t *template.Template
	c *metadatadb.Client
}

func fileHandler(dbc *metadatadb.Client) http.Handler {
	funcs := common.FuncMap
	funcs["isRace"] = func(file string) bool {
		return addons.GetAddonType(file)&(addons.RaceFlag|addons.BattleFlag) != 0
	}

	t := template.Must(template.New("file.tmpl.html").Funcs(funcs).ParseFiles("./templates/file.tmpl.html"))
	return &FileHandler{
		t,
		dbc,
	}
}

func (f *FileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	files, err := f.c.FindFilesByFilename(filename, r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	f.t.Execute(w, files)
}
