package maps

import (
	"github.com/gorilla/mux"
	"go.formulabun.club/metadatadb"
)

func SetupRouter(r *mux.Router, dbc *metadatadb.Client) {
	r.Handle("/maps", mapsHandler(dbc))
}
