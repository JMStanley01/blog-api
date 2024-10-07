package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var posts []Post

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get the route parameters
	for _, post := range posts {
		if post.ID == params["id"] {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}

func createPost(w http.ResponseWriter, r *http.Request) {
	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)
	posts = append(posts, newPost)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newPost)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, post := range posts {
		if post.ID == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)
			var updatedPost Post
			json.NewDecoder(r.Body).Decode(&updatedPost)
			updatedPost.ID = params["id"]
			posts = append(posts, updatedPost)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedPost)
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, post := range posts {
		if post.ID == params["id"] {
			posts = append(posts[:i], posts[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(posts)
			return
		}
	}
	http.Error(w, "Post not found", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	// Define the routes
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts", createPost).Methods("POST")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	fmt.Println("Server is running on port 8000...")
	http.ListenAndServe(":8000", router)
}
