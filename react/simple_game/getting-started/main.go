package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


func main() {
	r := mux.NewRouter()

	// Serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// Serve static assets from/static/{file} route.
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe("0.0.0.0:3000", handlers.LoggingHandler(os.Stdout, r))
}
