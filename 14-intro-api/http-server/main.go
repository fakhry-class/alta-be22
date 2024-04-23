package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type article struct {
	ID      int
	Title   string
	Content string
}

// data dummy
var data = []article{
	{1, "lorem", "lorem"},
	{2, "ipsum", "ipsum"},
	{3, "alterra", "academy"},
	{4, "BE", "22"},
}

func articles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		// object -> json
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	} else if r.Method == "POST" {
		// object -> json
		var response = map[string]any{
			"response": true,
			"status":   "success",
			"message":  "success add data",
		}
		var result, err = json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	// define endpoint
	http.HandleFunc("/articles", articles)
	http.HandleFunc("/books", articles)
	fmt.Println("starting web server at http://localhost:8080/")
	// start server and port
	http.ListenAndServe(":8080", nil)
}
