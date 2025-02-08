package services

import (
	"errors"
	"github.com/google/uuid"
	"operator_text_channel/src/storage/models"
	"time"
)

func (s *Service) GetTags(limit, offset int) ([]Tag, error) {
	var tags []Tag

	tx := s.DB.Model(&models.Tag{}).
		Where("deleted = false").
		Limit(limit).
		Offset(offset).
		Scan(&tags)
	return tags, tx.Error
}

func (s *Service) GetTagById(idStr string) (*Tag, error) {
	var tag Tag

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, errors.New("invalid tag id")
	}
	tx := s.DB.Model(&models.Tag{}).
		Where("id = ?", id).
		First(&tag)
	if tx.Error != nil {
		return nil, errors.New("tag not found")
	}
	return &tag, nil
}

func (s *Service) CreateTag(name string) (*Tag, error) {
	var tag models.Tag

	tag.Name = name
	tag.CreatedAt = time.Now()
	tag.UpdatedAt = time.Now()

	tx := s.DB.Model(&models.Tag{}).Create(&tag)
	if tx.Error != nil {
		return nil, tx.Error
	}
	exportTag := Tag{ID: tag.ID, Name: tag.Name, CreatedAt: tag.CreatedAt, UpdatedAt: tag.UpdatedAt}
	return &exportTag, nil
}

func (s *Service) DeleteTag(idStr string) error {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid tag id")
	}
	tx := s.DB.Model(&models.Tag{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		})
	if tx.Error != nil {
		return errors.New("invalid tag id")
	}
	return nil
}
