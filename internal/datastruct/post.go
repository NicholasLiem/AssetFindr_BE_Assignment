package datastruct

import (
	"encoding/json"
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
	Tags        []*Tag    `gorm:"many2many:post_tags;" json:"tags,omitempty"`
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

func (post *Post) UnmarshalJSON(data []byte) error {
	type Alias Post
	aux := &struct {
		Tags []string `json:"tags"`
		*Alias
	}{
		Alias: (*Alias)(post),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var tags []*Tag
	for _, tagLabel := range aux.Tags {
		tags = append(tags, &Tag{Label: tagLabel})
	}
	post.Tags = tags

	return nil
}

func (post *Post) MarshalJSON() ([]byte, error) {
	type Alias Post
	aux := &struct {
		Tags []string `json:"tags"`
		*Alias
	}{
		Tags:  make([]string, len(post.Tags)),
		Alias: (*Alias)(post),
	}
	for i, tag := range post.Tags {
		aux.Tags[i] = tag.Label
	}
	return json.Marshal(aux)
}
