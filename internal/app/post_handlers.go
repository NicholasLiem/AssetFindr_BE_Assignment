package app

import (
	"encoding/json"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/datastruct"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func (m *MicroserviceServer) CreatePost(rw http.ResponseWriter, r *http.Request) {
	var post datastruct.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result, httpErr := m.postService.CreatePost(post)
	if httpErr != nil {
		http.Error(rw, httpErr.Message, httpErr.StatusCode)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(rw).Encode(result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *MicroserviceServer) GetPost(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]

	parsedPostId, err := utils.ParseStrToUint(postId)
	if err != nil {
		http.Error(rw, "Invalid post id", http.StatusBadRequest)
		return
	}

	result, httpErr := m.postService.GetPost(*parsedPostId)
	if httpErr != nil {
		http.Error(rw, httpErr.Message, httpErr.StatusCode)
		return
	}

	if err := json.NewEncoder(rw).Encode(result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *MicroserviceServer) UpdatePost(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]

	parsedPostId, err := utils.ParseStrToUint(postId)
	if err != nil {
		http.Error(rw, "Invalid post id", http.StatusBadRequest)
		return
	}

	var updatedPost datastruct.Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(rw, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result, httpErr := m.postService.UpdatePost(*parsedPostId, updatedPost)
	if httpErr != nil {
		http.Error(rw, httpErr.Message, httpErr.StatusCode)
		return
	}

	if err := json.NewEncoder(rw).Encode(result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *MicroserviceServer) DeletePost(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postId := params["id"]

	parsedPostId, err := utils.ParseStrToUint(postId)
	if err != nil {
		http.Error(rw, "Invalid post id", http.StatusBadRequest)
		return
	}

	result, httpErr := m.postService.DeletePost(*parsedPostId)
	if httpErr != nil {
		http.Error(rw, httpErr.Message, httpErr.StatusCode)
		return
	}

	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *MicroserviceServer) GetPagedPost(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	limit := params["limit"]
	offset := params["offset"]

	parsedLimit, err := utils.ParseStrToInt(limit)
	if err != nil {
		http.Error(rw, "Invalid limit number", http.StatusBadRequest)
		return
	}

	parsedOffset, err := utils.ParseStrToInt(offset)
	if err != nil {
		http.Error(rw, "Invalid post id", http.StatusBadRequest)
		return
	}

	pagedPosts, httpErr := m.postService.GetPagedPost(*parsedLimit, *parsedOffset)
	if httpErr != nil {
		http.Error(rw, httpErr.Message, httpErr.StatusCode)
		return
	}

	totalPages := (pagedPosts.TotalCount + *parsedLimit - 1) / *parsedLimit
	currentPage := *parsedOffset / *parsedLimit + 1

	response := struct {
		Posts       []datastruct.Post `json:"posts"`
		TotalPages  int               `json:"totalPages"`
		CurrentPage int               `json:"currentPage"`
	}{
		Posts:       pagedPosts.Posts,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(response); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (m *MicroserviceServer) GetAllPost(rw http.ResponseWriter, r *http.Request) {
	result, httpErr := m.postService.GetAllPost()
	if httpErr != nil {
		http.Error(rw, httpErr.Message, httpErr.StatusCode)
		return
	}

	if err := json.NewEncoder(rw).Encode(result); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}
