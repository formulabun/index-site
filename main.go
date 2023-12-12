package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"go.formulabun.club/index-site/common"
	"go.formulabun.club/index-site/files"
	"go.formulabun.club/index-site/maps"
	"go.formulabun.club/metadatadb"
	"go.formulabun.club/storage"

	"github.com/gorilla/mux"
)

var host = flag.String("host", ":8080", "Hostname on which to bind")

func main() {
	flag.Parse()
	r := mux.NewRouter()

	connectContext, _ := context.WithTimeout(context.Background(), time.Second*10)
	dbc, err := metadatadb.NewClient(connectContext)
	if err != nil {
		panic(err)
	}

	// files
	r.PathPrefix(common.FilePath).Handler(http.StripPrefix(common.FilePath, storage.StoreHandler{}))

	// api
	apiRouter := r.Methods("GET", "POST").PathPrefix("/api").Subrouter()
	files.SetupRouter(apiRouter, dbc)
	maps.SetupRouter(apiRouter, dbc)

	// pages
	staticRoute := r.Methods("GET").Subrouter()
	staticRoute.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("static/public/"))))
	staticRoute.PathPrefix("/").Handler(http.FileServer(http.Dir("static/html/")))

	log.Printf("hosting on http://%s\n", *host)
	log.Fatal(http.ListenAndServe(*host, r))
}
