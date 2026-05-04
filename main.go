package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db, _ = sql.Open("sqlite3", "./database/mhdb.db")

var RouteRegistry []func(*http.ServeMux)

func Register(fn func(*http.ServeMux)) {
	RouteRegistry = append(RouteRegistry, fn)
}

func main() {
	mux := http.NewServeMux()

	denyMethod := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(os.Stderr, "Invalid method %s on route %s\n", r.Method, r.URL)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	mux.HandleFunc("PUT /{path...}", denyMethod)
	mux.HandleFunc("PATCH /{path...}", denyMethod)
	mux.HandleFunc("POST /{path...}", denyMethod)
	mux.HandleFunc("DELETE /{path...}", denyMethod)

	for _, registerRoute := range RouteRegistry {
		registerRoute(mux)
	}

	mux.HandleFunc("GET /", getRoot)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to / recieved")
	fmt.Fprintf(w, "Hello, World!\n")
}
