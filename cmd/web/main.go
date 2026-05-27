package main

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/raffy-io/testdeploy"
)

func main() {

	// routes
	mux := http.NewServeMux()

	// static assets
	staticFS,err := fs.Sub(testdeploy.EmbeddedAssets,"static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	// server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on clean port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))

}