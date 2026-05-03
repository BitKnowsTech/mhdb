package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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

	mux.HandleFunc("GET /", getRoot)
	mux.HandleFunc("GET /monsters", getMonsters)
	mux.HandleFunc("GET /monster/{id}", getMonsterById)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to / recieved")
	fmt.Fprintf(w, "Hello, World!\n")
}

func getMonsters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /monster recieved")
	fmt.Fprintf(w, "MONSTERS HERE\n")
}

func getMonsterById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MONSTER %s", r.PathValue("id"))
}
