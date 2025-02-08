package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"operator_text_channel/src/storage/models"
	"time"
)

func (s *Service) GetAppealsByUserId(userIdStr string) ([]Appeal, error) {
	var appeals []models.Appeal

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	tx := s.DB.Preload("AppealTagLinks", "deleted = false").
		Preload("AppealTagLinks.Tag", "deleted = false").
		Find(&appeals, "appeals.deleted = false AND appeals.user_id = ?", userId)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return convertAppealsToExport(appeals), nil
}

func (s *Service) CreateAppeal(userIdStr string, tagIdsStr []string) (*Appeal, error) {
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return nil, errors.New("invalid user id")
	}
	tagIds := make([]uuid.UUID, len(tagIdsStr))
	for i, tagIdStr := range tagIdsStr {
		tagId, err := uuid.Parse(tagIdStr)
		if err != nil {
			return nil, errors.New("invalid tag id")
		}
		tagIds[i] = tagId
	}

	var appeal models.Appeal
	appeal.UserId = userId
	appeal.CreatedAt = time.Now()
	appeal.UpdatedAt = time.Now()

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&appeal)
		if tx.Error != nil {
			return tx.Error
		}
		for _, tagId := range tagIds {
			var appealTag models.AppealTagLink
			appealTag.AppealId = appeal.ID
			appealTag.TagId = tagId
			appealTag.CreatedAt = time.Now()
			appealTag.UpdatedAt = time.Now()
			tx.Create(&appealTag)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	var exportAppeal = Appeal{ID: appeal.ID, UserId: appeal.UserId, CreatedAt: appeal.CreatedAt, UpdatedAt: appeal.UpdatedAt}
	err = s.loadAppealTags(&exportAppeal)
	if err != nil {
		return nil, err
	}
	return &exportAppeal, nil
}

func (s *Service) DeleteAppeal(idStr string) error {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid appeal id")
	}
	tx := s.DB.Model(&models.Appeal{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		})
	if tx.Error != nil {
		return errors.New("invalid appeal id")
	}
	return nil
}

func (s *Service) AddTagsToAppeal(appealIdStr string, tagIdsStr []string) error {
	appealId, err := uuid.Parse(appealIdStr)
	if err != nil {
		return errors.New("invalid appeal id")
	}
	err = s.DB.Transaction(func(tx *gorm.DB) error {
		for _, tagIdStr := range tagIdsStr {
			tagId, err := uuid.Parse(tagIdStr)
			if err != nil {
				return errors.New("invalid tag id")
			}

			var appealTagLink models.AppealTagLink
			appealTagLink.AppealId = appealId
			appealTagLink.TagId = tagId
			appealTagLink.CreatedAt = time.Now()
			appealTagLink.UpdatedAt = time.Now()
			tx.Create(&appealTagLink)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	})
	return err
}

func (s *Service) RemoveTagsFromAppeal(appealIdStr string, tagIdsStr []string) error {
	appealId, err := uuid.Parse(appealIdStr)
	if err != nil {
		return errors.New("invalid appeal id")
	}
	var tagIds []uuid.UUID
	for _, tagIdStr := range tagIdsStr {
		tagId, err := uuid.Parse(tagIdStr)
		if err != nil {
			return errors.New("invalid tag id")
		}
		tagIds = append(tagIds, tagId)
	}
	tx := s.DB.Model(&models.AppealTagLink{}).
		Where("appeal_id = ?", appealId).
		Where("tag_id IN ?", tagIds).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		})
	return tx.Error
}
