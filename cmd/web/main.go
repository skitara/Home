package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Welcome to my blog!"))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		w.Header().Set("Allow", http.MethodPost)
		//w.WriteHeader(405)
		//w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Now Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create post!"))
}

func changePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		w.Header().Set("Allow", http.MethodPut)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Change post!"))
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		w.Header().Set("Allow", http.MethodDelete)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Delete post"))
}

func main() {
	// servemux for routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/create", createPost)
	mux.HandleFunc("/change", changePost)
	mux.HandleFunc("/delete", deletePost)

	log.Print("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
