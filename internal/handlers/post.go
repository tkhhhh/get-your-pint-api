package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Excerpt string `json:"excerpt"`
}

type PostHandler struct {
	DB *pgxpool.Pool
}

func (h *PostHandler) List(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query(r.Context(),
		`SELECT id, title, content, excerpt FROM posts ORDER BY id`)
	if err != nil {
		http.Error(w, "failed to load posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []Post{}
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Excerpt); err != nil {
			http.Error(w, "failed to scan post", http.StatusInternalServerError)
			return
		}
		posts = append(posts, p)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "failed to iterate posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var p Post
	err = h.DB.QueryRow(r.Context(),
		`SELECT id, title, content, excerpt FROM posts WHERE id=$1`, id,
	).Scan(&p.ID, &p.Title, &p.Content, &p.Excerpt)
	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "failed to load post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(p)
}
