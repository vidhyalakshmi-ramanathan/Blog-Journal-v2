package handlers

import (
	"backend/internal/entities"
	"backend/internal/services"
	"backend/middleware"
	"encoding/json"
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
)

type PostHandler struct {
	service services.PostService
}

func NewPostHandler(service services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var postData entities.PostData
	claims, ok := middleware.GetClaims(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	post := entities.Post{
		Title:    postData.Title,
		Content:  postData.Content,
		AuthorID: claims.UserID,
	}
	if err := h.service.Create(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post Created Successfully!"})
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var postData entities.PostData

	claims, ok := middleware.GetClaims(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&postData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post := entities.Post{
		Title:    postData.Title,
		Content:  postData.Content,
		AuthorID: claims.UserID,
	}
	if err := h.service.Update(uint(id), claims.UserID, &post); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post Updated Successfully!"})
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "invalid post id", http.StatusBadRequest)
		return
	}
	if err := h.service.Delete(uint(id), claims.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post Deleted Successfully!"})
}

func (h *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	post, err := h.service.GetPost(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(entities.ApiResponse{
		Message: "Retrieved user post",
		Data:    post,
	})
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	claims, ok := middleware.GetClaims(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	posts, err := h.service.GetPosts(claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(entities.ApiResponse{
		Message: "Retrieved the user posts",
		Data:    posts,
	})
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func (h *PostHandler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	// userData, err := h.service.GetUserDetails()
	claims, ok := middleware.GetClaims(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(entities.ApiResponse{
		Message: "Retrieved the user details",
		Data:    map[string]string{"username": claims.Username},
	})
}
