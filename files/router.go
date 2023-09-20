package files

import (
	"github.com/gorilla/mux"
	"go.formulabun.club/metadatadb"
)

func SetupRouter(r *mux.Router, dbc *metadatadb.Client) {
	r.Handle("/files", filesHandler(dbc))
	r.Handle("/file", fileHandler(dbc))
}
