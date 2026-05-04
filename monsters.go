package main

import (
	"fmt"
	"net/http"
)

func init() {
	Register(func(mux *http.ServeMux) {
		mux.HandleFunc("GET /monsters", getMonsters)
		mux.HandleFunc("GET /monster/{id}", getMonsterById)
	})
}

func getMonsters(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	querylog := ""
	for key, values := range query {
		querylog += key + ":"
		for i, value := range values {
			if i > 0 {
				querylog += ";"
			}
			querylog += value
		}
	}
	fmt.Println("Request to /monsters recieved")
	if querylog != "" {
		fmt.Println("With query params: " + querylog)
	}

	fmt.Fprintf(w, "MONSTERS HERE\n")
}

func getMonsterById(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "MONSTER %s", r.PathValue("id"))
}
