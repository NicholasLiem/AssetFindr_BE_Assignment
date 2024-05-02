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
	Category    string    `gorm:"size:100;not null" json:"category"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"type:varchar(100);not null;check:status in ('Publish','Draft','Thrash')" json:"status"`
}

func (post *Post) TableName() string {
	return "posts"
}

type PagedPosts struct {
	Posts      []Post
	TotalCount int
}

func (post *Post) BeforeSave(db *gorm.DB) (err error) {
	if len(post.Title) < 20 {
		return errors.New("the title must be at least 20 characters long")
	}
	if len(post.Content) < 200 {
		return errors.New("the content must be at least 200 characters long")
	}

	if len(post.Category) < 3 {
		return errors.New("the category must be at least 3 characters long")
	}

	validStatus := map[string]bool{"publish": true, "draft": true, "thrash": true}
	if _, ok := validStatus[post.Status]; !ok {
		return errors.New("invalid status, must be 'publish', 'draft', or 'thrash'")
	}
	return nil
}
