package main

import (
	"log"
	"net/http"
)

// handler for logic
func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my blog!"))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create post!"))
}

func changePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Change post!"))
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete post"))
}

func main() {
	// servemux for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/create", createPost)
	mux.HandleFunc("/change", changePost)
	mux.HandleFunc("/delete", deletePost)

	log.Print("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
