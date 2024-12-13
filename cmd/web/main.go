package main

import (
	"log"
	"net/http"
)

func methodCheck(handler func(http.ResponseWriter, *http.Request), allowedMethod string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			w.Header().Set("Allow", allowedMethod)
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", methodCheck(home, http.MethodGet))
	mux.HandleFunc("/create", methodCheck(createPost, http.MethodPost))
	mux.HandleFunc("/change", methodCheck(changePost, http.MethodPut))
	mux.HandleFunc("/delete", methodCheck(deletePost, http.MethodDelete))

	log.Print("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
