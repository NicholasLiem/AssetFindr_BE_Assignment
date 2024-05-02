package repository

import (
	"errors"
	"github.com/NicholasLiem/AssetFindr_BackendAssignment/internal/datastruct"
	"gorm.io/gorm"
)

type TagQuery interface {
	BeginTransaction() *gorm.DB
	CreateTag(tag datastruct.Tag) (bool, error)
	CreateTagTx(tag datastruct.Tag, tx *gorm.DB) (bool, error)
	UpdateTagTx(tagId uint, updatedTag datastruct.Tag, tx *gorm.DB) (bool, error)
	DeleteTag(tagId uint) (bool, error)
	GetTag(tagId uint) (*datastruct.Tag, error)
	GetAllTags() ([]datastruct.Tag, error)
	FindOrCreateTag(tagName string, tx *gorm.DB) (*datastruct.Tag, error)
	FindTagByLabel(tagName string) (*datastruct.Tag, error)
}

type tagQuery struct {
	pgdb *gorm.DB
}

func NewTagQuery(db *gorm.DB) TagQuery {
	return &tagQuery{pgdb: db}
}

func (tq *tagQuery) BeginTransaction() *gorm.DB {
	return tq.pgdb.Begin()
}

func (tq *tagQuery) CreateTag(tag datastruct.Tag) (bool, error) {
	result := tq.pgdb.Create(&tag)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (tq *tagQuery) CreateTagTx(tag datastruct.Tag, tx *gorm.DB) (bool, error) {
	result := tx.Create(&tag)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (tq *tagQuery) DeleteTag(tagId uint) (bool, error) {
	result := tq.pgdb.Delete(&datastruct.Tag{}, tagId)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (tq *tagQuery) UpdateTagTx(tagId uint, updatedTag datastruct.Tag, tx *gorm.DB) (bool, error) {
	var tag datastruct.Tag
	if err := tx.First(&tag, tagId).Error; err != nil {
		return false, err
	}

	result := tx.Model(&tag).Updates(updatedTag)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (tq *tagQuery) UpdateTag(tagId uint, updatedTag datastruct.Tag) (bool, error) {
	var tag datastruct.Tag
	result := tq.pgdb.First(&tag, tagId)
	if result.Error != nil {
		return false, result.Error
	}

	result = tq.pgdb.Model(&tag).Updates(updatedTag)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (tq *tagQuery) GetTag(tagId uint) (*datastruct.Tag, error) {
	var tag datastruct.Tag
	result := tq.pgdb.First(&tag, tagId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tag, nil
}

func (tq *tagQuery) GetAllTags() ([]datastruct.Tag, error) {
	var tags []datastruct.Tag
	result := tq.pgdb.Find(&tags)
	if result.Error != nil {
		return nil, result.Error
	}
	return tags, nil
}

func (tq *tagQuery) FindOrCreateTag(tagName string, tx *gorm.DB) (*datastruct.Tag, error) {
	var tag datastruct.Tag
	if err := tx.Where("label = ?", tagName).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tag.Label = tagName
			if err := tx.Create(&tag).Error; err != nil {
				return nil, err
			}
			return &tag, nil
		}
		return nil, err
	}
	return &tag, nil
}

func (tq *tagQuery) FindTagByLabel(tagName string) (*datastruct.Tag, error) {
	var tag datastruct.Tag
	if err := tq.pgdb.Where("label = ?", tagName).First(&tag).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}
