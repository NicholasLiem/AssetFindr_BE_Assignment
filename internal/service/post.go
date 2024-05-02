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
	GetPostWithTag(postID uint) (*datastruct.Post, error)
	GetAllPost() (*[]datastruct.Post, *utils.HttpError)
	GetTagsForPost(post *datastruct.Post) ([]*datastruct.Tag, error)
	GetAllPostWithTags() ([]*datastruct.Post, error)
}

type postService struct {
	dao repository.DAO
}

func NewPostService(dao repository.DAO) PostService {
	return &postService{dao: dao}
}

func (ps *postService) CreatePost(post datastruct.Post) (bool, *utils.HttpError) {
	tx := ps.dao.NewPostQuery().BeginTransaction()
	if tx.Error != nil {
		return false, &utils.HttpError{Message: "Error starting database transaction", StatusCode: http.StatusInternalServerError}
	}

	tagQuery := ps.dao.NewTagQuery()
	for i := range post.Tags {
		tag, err := tagQuery.FindOrCreateTag(post.Tags[i].Label, tx)
		if err != nil {
			tx.Rollback()
			return false, &utils.HttpError{Message: "Error handling tag: " + err.Error(), StatusCode: http.StatusInternalServerError}
		}
		post.Tags[i] = tag
	}

	success, err := ps.dao.NewPostQuery().CreatePostTx(post, tx)
	if err != nil {
		tx.Rollback()
		return false, &utils.HttpError{Message: "Error creating post: " + err.Error(), StatusCode: http.StatusInternalServerError}
	}

	commitResult := tx.Commit()
	if commitResult.Error != nil {
		tx.Rollback()
		return false, &utils.HttpError{Message: "Error committing transaction: " + commitResult.Error.Error(), StatusCode: http.StatusInternalServerError}
	}

	return success, nil
}

func (ps *postService) UpdatePost(postID uint, updatedPost datastruct.Post) (bool, *utils.HttpError) {
	tx := ps.dao.NewPostQuery().BeginTransaction()
	if tx.Error != nil {
		return false, &utils.HttpError{Message: "Error starting database transaction", StatusCode: http.StatusInternalServerError}
	}

	_, err := ps.dao.NewPostQuery().UpdatePostTx(postID, updatedPost, tx)
	if err != nil {
		tx.Rollback()
		return false, &utils.HttpError{Message: "Error updating post: " + err.Error(), StatusCode: http.StatusInternalServerError}
	}

	commitResult := tx.Commit()
	if commitResult.Error != nil {
		tx.Rollback()
		return false, &utils.HttpError{Message: "Error committing transaction: " + commitResult.Error.Error(), StatusCode: http.StatusInternalServerError}
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

func (ps *postService) GetTagsForPost(post *datastruct.Post) ([]*datastruct.Tag, error) {
	tags, err := ps.dao.NewPostQuery().GetTagsForPost(post.ID)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (ps *postService) GetAllPostWithTags() ([]*datastruct.Post, error) {
	postsPtr, err := ps.dao.NewPostQuery().GetAllPost()
	if err != nil {
		return nil, err
	}

	posts := make([]*datastruct.Post, len(*postsPtr))
	for i, post := range *postsPtr {
		posts[i] = &post
		tags, err := ps.GetTagsForPost(&post)
		if err != nil {
			return nil, err
		}
		posts[i].Tags = tags
	}

	return posts, nil
}

func (ps *postService) GetPostWithTag(postID uint) (*datastruct.Post, error) {
	post, err := ps.dao.NewPostQuery().GetPost(postID)
	if err != nil {
		return nil, err
	}

	tags, err := ps.GetTagsForPost(post)
	if err != nil {
		return nil, err
	}
	post.Tags = tags

	return post, nil
}
