package datastruct

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Label     string    `gorm:"size:100;not null;unique" json:"label"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Posts     []Post    `gorm:"many2many:post_tags;"`
}

func (tag *Tag) TableName() string {
	return "tags"
}

func (tag *Tag) BeforeSave(db *gorm.DB) (err error) {
	if len(tag.Label) < 3 {
		return errors.New("the label must be at least 3 characters long")
	}

	return nil
}
