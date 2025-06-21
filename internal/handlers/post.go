package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
)

// Mock data for blog posts
var posts = []struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Excerpt string `json:"excerpt"`
}{
    {ID: 1, Title: "First Blog Post", Content: "This is the content of the first blog post.", Excerpt: "This is the first blog post."},
    {ID: 2, Title: "Second Blog Post", Content: "This is the content of the second blog post.", Excerpt: "This is the second blog post."},
    {ID: 3, Title: "Third Blog Post", Content: "This is the content of the third blog post.", Excerpt: "This is the third blog post."},
}

// GetPosts handler - returns all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

// GetPostByID handler - returns a single post by ID
func GetPostByID(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    for _, post := range posts {
        if post.ID == id {
            json.NewEncoder(w).Encode(post)
            return
        }
    }

    http.Error(w, "Post not found", http.StatusNotFound)
}