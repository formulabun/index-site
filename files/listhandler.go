package files

import (
	"html/template"
	"log"
	"net/http"

	"go.formulabun.club/metadatadb"

	"go.formulabun.club/srb2kart/addons"
)

type FileListHandler struct {
	c *metadatadb.Client
	t *template.Template
}

func filesHandler(dbClient *metadatadb.Client) http.Handler {
	t := template.Must(template.ParseFiles("templates/filelist.tmpl.html"))

	return &FileListHandler{
		dbClient,
		t,
	}
}

func (f *FileListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addonType := addons.KartFlag
	if r.FormValue("race") == "on" {
		addonType |= addons.RaceFlag
	}
	if r.FormValue("battle") == "on" {
		addonType |= addons.BattleFlag
	}
	if r.FormValue("char") == "on" {
		addonType |= addons.CharFlag
	}
	if r.FormValue("lua") == "on" {
		addonType |= addons.LuaFlag
	}

	files, err := f.c.FindFiles(addonType, r.FormValue("andor"), r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	f.t.Execute(w, filterFiles(r.FormValue("file"), files))
}
