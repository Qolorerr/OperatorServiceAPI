package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"operator_text_channel/src/storage/models"
	"time"
)

func (s *Service) CreateOperator() (*Operator, error) {
	var operator models.Operator
	operator.CreatedAt = time.Now()
	operator.UpdatedAt = time.Now()
	tx := s.DB.Create(&operator)
	if tx.Error != nil {
		return nil, tx.Error
	}
	exportOperator := Operator{ID: operator.ID, CreatedAt: operator.CreatedAt, UpdatedAt: operator.UpdatedAt}
	return &exportOperator, nil
}

func (s *Service) DeleteOperator(idStr string) error {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid operator id")
	}
	tx := s.DB.Model(&models.Operator{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		})
	return tx.Error
}

func (s *Service) GetOperatorTags(idStr string) ([]Tag, error) {
	var tags []Tag

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, errors.New("invalid operator id")
	}
	tx := s.DB.Model(&models.Operator{}).
		Select("tags.id, tags.name, tags.created_at, tags.updated_at").
		Joins("LEFT JOIN operator_group_links AS ogl ON ogl.operator_id = operators.id AND ogl.deleted = false").
		Joins("LEFT JOIN operator_groups AS og ON og.id = ogl.group_id AND og.deleted = false").
		Joins("LEFT JOIN group_tag_links AS gtl ON gtl.group_id = og.id AND gtl.deleted = false").
		Joins("LEFT JOIN tags ON tags.id = gtl.tag_id AND tags.deleted = false").
		Where("operators.deleted = false AND operators.id = ?", id).
		Scan(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (s *Service) getOperatorAppealsQuery(id uuid.UUID) *gorm.DB {
	operatorTagIdsQuery := s.DB.Model(&models.Operator{}).
		Select("gtl.tag_id").
		Joins("LEFT JOIN operator_group_links AS ogl ON ogl.operator_id = operators.id AND ogl.deleted = false").
		Joins("LEFT JOIN operator_groups AS og ON og.id = ogl.group_id AND og.deleted = false").
		Joins("LEFT JOIN group_tag_links AS gtl ON gtl.group_id = og.id AND gtl.deleted = false").
		Where("operators.deleted = false AND operators.id = ?", id)
	unsuitableAppealTagIdsQuery := s.DB.Model(&models.AppealTagLink{}).
		Select("1").
		Where("appeal_tag_links.deleted = false AND appeal_tag_links.appeal_id = appeals.id").
		Where("appeal_tag_links.tag_id NOT IN (?)", operatorTagIdsQuery)
	tx := s.DB.Model(&models.Appeal{}).
		Select("appeals.*").
		Where("appeals.deleted = false").
		Where("NOT EXISTS (?)", unsuitableAppealTagIdsQuery).
		Order("appeals.weight DESC, appeals.created_at")
	return tx
}

func (s *Service) GetOperatorAppeals(idStr string, limit int) ([]Appeal, error) {
	var appeals []models.Appeal

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, errors.New("invalid operator id")
	}
	tx := s.getOperatorAppealsQuery(id).
		Limit(limit).
		Scan(&appeals)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return convertAppealsToExport(appeals), nil
}
