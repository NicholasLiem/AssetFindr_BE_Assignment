package service

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/datastruct"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/repository"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/utils"
	"net/http"
)

type PostService interface {
	CreatePost(post datastruct.Post) (bool, *utils.HttpError)
	UpdatePost(postID uint, updatedPost datastruct.Post) (bool, *utils.HttpError)
	DeletePost(postID uint) (bool, *utils.HttpError)
	GetPost(postID uint) (*datastruct.Post, *utils.HttpError)
	GetAllPost() (*[]datastruct.Post, *utils.HttpError)
}

type postService struct {
	dao repository.DAO
}

func NewPostService(dao repository.DAO) PostService {
	return &postService{dao: dao}
}

func (ps *postService) CreatePost(post datastruct.Post) (bool, *utils.HttpError) {
	success, err := ps.dao.NewPostQuery().CreatePost(post)
	if err != nil || success == false {
		return false, &utils.HttpError{Message: "Error creating post: " +
			err.Error(), StatusCode: http.StatusInternalServerError}
	}
	return true, nil
}

func (ps *postService) UpdatePost(postID uint, updatedPost datastruct.Post) (bool, *utils.HttpError) {
	success, err := ps.dao.NewPostQuery().UpdatePost(postID, updatedPost)
	if err != nil || success == false {
		return false, &utils.HttpError{Message: "Error updating post: " +
			err.Error(), StatusCode: http.StatusInternalServerError}
	}
	return true, nil
}

func (ps *postService) DeletePost(postID uint) (bool, *utils.HttpError) {
	success, err := ps.dao.NewPostQuery().DeletePost(postID)
	if err != nil || success == false {
		return false, &utils.HttpError{Message: "Error deleting post: " +
			err.Error(), StatusCode: http.StatusInternalServerError}
	}
	return true, nil
}

func (ps *postService) GetPost(postID uint) (*datastruct.Post, *utils.HttpError) {
	post, err := ps.dao.NewPostQuery().GetPost(postID)
	if err != nil {
		return nil, &utils.HttpError{Message: "Error getting post: " +
			err.Error(), StatusCode: http.StatusInternalServerError}
	}
	return post, nil
}

func (ps *postService) GetAllPost() (*[]datastruct.Post, *utils.HttpError) {
	posts, err := ps.dao.NewPostQuery().GetAllPost()
	if err != nil {
		return nil, &utils.HttpError{Message: "Error getting post: " +
			err.Error(), StatusCode: http.StatusInternalServerError}
	}

	return posts, nil
}
