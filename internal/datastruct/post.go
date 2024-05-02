package datastruct

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"size:200;not null" json:"title"`
	Content     string    `gorm:"type:text;not null" json:"content"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
	//Tags        []Tag     `gorm:"many2many:post_tags;" json:"tags"`
}

func (post *Post) TableName() string {
	return "posts"
}

func (post *Post) BeforeSave(db *gorm.DB) (err error) {
	if len(post.Title) < 20 {
		return errors.New("the title must be at least 20 characters long")
	}
	if len(post.Content) < 200 {
		return errors.New("the content must be at least 200 characters long")
	}

	return nil
}
