package repository

import (
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/datastruct"
	"gorm.io/gorm"
)

type PostQuery interface {
	BeginTransaction() *gorm.DB
	CreatePost(post datastruct.Post) (bool, error)
	CreatePostTx(post datastruct.Post, tx *gorm.DB) (bool, error)
	UpdatePost(postId uint, updatedPost datastruct.Post) (bool, error)
	UpdatePostTx(postId uint, updatedPost datastruct.Post, tx *gorm.DB) (bool, error)
	DeletePost(postId uint) (bool, error)
	GetPost(postId uint) (*datastruct.Post, error)
	GetAllPost() ([]*datastruct.Post, error)
	GetTagsForPost(postID uint) ([]*datastruct.Tag, error)
	DisassociateTagsFromPost(postID uint) error
}

type postQuery struct {
	pgdb *gorm.DB
}

func NewPostQuery(mysql *gorm.DB) PostQuery {
	return &postQuery{
		pgdb: mysql,
	}
}

func (pq *postQuery) BeginTransaction() *gorm.DB {
	return pq.pgdb.Begin()
}

func (pq *postQuery) CreatePost(post datastruct.Post) (bool, error) {
	result := pq.pgdb.Create(&post)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (pq *postQuery) CreatePostTx(post datastruct.Post, tx *gorm.DB) (bool, error) {
	result := tx.Create(&post)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (pq *postQuery) UpdatePost(postId uint, updatedPost datastruct.Post) (bool, error) {
	var post datastruct.Post
	if err := pq.pgdb.Preload("Tags").First(&post, postId).Error; err != nil {
		return false, err
	}

	post.Title = updatedPost.Title
	post.Content = updatedPost.Content

	post.Tags = updatedPost.Tags

	if err := pq.pgdb.Save(&post).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pq *postQuery) UpdatePostTx(postId uint, updatedPost datastruct.Post, tx *gorm.DB) (bool, error) {
	var post datastruct.Post
	if err := tx.Preload("Tags").First(&post, postId).Error; err != nil {
		return false, err
	}

	post.Title = updatedPost.Title
	post.Content = updatedPost.Content

	post.Tags = updatedPost.Tags

	if err := tx.Save(&post).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (pq *postQuery) DeletePost(postId uint) (bool, error) {
	if err := pq.DisassociateTagsFromPost(postId); err != nil {
		return false, err
	}

	var post datastruct.Post
	if err := pq.pgdb.First(&post, postId).Error; err != nil {
		return false, err
	}

	result := pq.pgdb.Delete(&post)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (pq *postQuery) GetPost(postId uint) (*datastruct.Post, error) {
	var post datastruct.Post
	if err := pq.pgdb.First(&post, postId).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (pq *postQuery) GetAllPost() ([]*datastruct.Post, error) {
	var posts []*datastruct.Post

	result := pq.pgdb.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	return posts, nil
}

func (pq *postQuery) GetTagsForPost(postID uint) ([]*datastruct.Tag, error) {
	var tags []*datastruct.Tag
	if err := pq.pgdb.Model(&datastruct.Post{ID: postID}).Association("Tags").Find(&tags); err != nil {
		return nil, err
	}
	return tags, nil
}

func (pq *postQuery) DisassociateTagsFromPost(postID uint) error {
	if err := pq.pgdb.Exec("DELETE FROM post_tags WHERE post_id = ?", postID).Error; err != nil {
		return err
	}
	return nil
}
