package services

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"operator_text_channel/src/storage/models"
	"time"
)

func (s *Service) GetGroups() ([]OperatorGroup, error) {
	var groups []models.OperatorGroup

	tx := s.DB.Preload("GroupTagLinks", "deleted = false").
		Preload("GroupTagLinks.Tag", "deleted = false").
		Preload("OperatorGroupLinks", "deleted = false").
		Preload("OperatorGroupLinks.Operator", "deleted = false").
		Find(&groups, "operator_groups.deleted = false")
	if tx.Error != nil {
		return nil, tx.Error
	}
	return convertGroupsToExport(groups), nil
}

func (s *Service) CreateGroup(tagIdsStr []string) (*OperatorGroup, error) {
	tagIds := make([]uuid.UUID, len(tagIdsStr))
	for i, tagIdStr := range tagIdsStr {
		tagId, err := uuid.Parse(tagIdStr)
		if err != nil {
			return nil, errors.New("invalid tag id")
		}
		tagIds[i] = tagId
	}

	var group models.OperatorGroup
	group.CreatedAt = time.Now()
	group.UpdatedAt = time.Now()

	err := s.DB.Transaction(func(tx *gorm.DB) error {
		tx.Create(&group)
		if tx.Error != nil {
			return tx.Error
		}
		for _, tagId := range tagIds {
			var groupTagLink models.GroupTagLink
			groupTagLink.GroupId = group.ID
			groupTagLink.TagId = tagId
			groupTagLink.CreatedAt = time.Now()
			groupTagLink.UpdatedAt = time.Now()
			tx.Create(&groupTagLink)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	exportGroup := OperatorGroup{ID: group.ID, CreatedAt: group.CreatedAt, UpdatedAt: group.UpdatedAt}
	err = s.loadGroupTags(&exportGroup)
	if err != nil {
		return nil, err
	}
	return &exportGroup, nil
}

func (s *Service) DeleteGroup(idStr string) error {
	id, err := uuid.Parse(idStr)
	if err != nil {
		return errors.New("invalid group id")
	}
	tx := s.DB.Model(&models.OperatorGroup{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		})
	if tx.Error != nil {
		return errors.New("failed to delete group")
	}
	return nil
}

func (s *Service) GetGroupOperators(idStr string) ([]Operator, error) {
	var operators []Operator

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, errors.New("invalid group id")
	}
	tx := s.DB.Model(&models.OperatorGroup{}).
		Select("operators.*").
		Joins("LEFT JOIN operator_group_links AS ogl ON ogl.group_id = operator_groups.id AND ogl.deleted = false").
		Joins("LEFT JOIN operators ON operators.id = ogl.operator_id AND operators.deleted = false").
		Where("operator_groups.deleted = false AND operator_groups.id = ?", id).
		Scan(&operators)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return operators, nil
}

func (s *Service) AddOperatorsToGroup(groupIdStr string, operatorIdsStr []string) error {
	groupId, err := uuid.Parse(groupIdStr)
	if err != nil {
		return errors.New("invalid group id")
	}
	operatorIds := make([]uuid.UUID, len(operatorIdsStr))
	for i, operatorIdStr := range operatorIdsStr {
		operatorId, err := uuid.Parse(operatorIdStr)
		if err != nil {
			return errors.New("invalid operator id")
		}
		operatorIds[i] = operatorId
	}
	err = s.DB.Transaction(func(tx *gorm.DB) error {
		for _, operatorId := range operatorIds {
			var operatorGroupLink models.OperatorGroupLink
			operatorGroupLink.GroupId = groupId
			operatorGroupLink.OperatorId = operatorId
			operatorGroupLink.CreatedAt = time.Now()
			operatorGroupLink.UpdatedAt = time.Now()
			tx.Create(&operatorGroupLink)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	})
	if err != nil {
		return errors.New("failed to add operators to group")
	}
	return nil
}

func (s *Service) RemoveOperatorsFromGroup(groupIdStr string, operatorIdsStr []string) error {
	groupId, err := uuid.Parse(groupIdStr)
	if err != nil {
		return errors.New("invalid group id")
	}
	operatorIds := make([]uuid.UUID, len(operatorIdsStr))
	for i, operatorIdStr := range operatorIdsStr {
		operatorId, err := uuid.Parse(operatorIdStr)
		if err != nil {
			return errors.New("invalid operator id")
		}
		operatorIds[i] = operatorId
	}
	tx := s.DB.Model(&models.OperatorGroupLink{}).
		Where("group_id = ?", groupId).
		Where("operator_id IN ?", operatorIds).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		})
	if tx.Error != nil {
		return errors.New("failed to remove operators from group")
	}
	return nil
}

func (s *Service) GetGroupTags(idStr string) ([]Tag, error) {
	var tags []Tag

	id, err := uuid.Parse(idStr)
	if err != nil {
		return nil, errors.New("invalid group id")
	}
	tx := s.DB.Model(&models.OperatorGroup{}).
		Select("tags.*").
		Joins("LEFT JOIN group_tag_links AS gtl ON gtl.group_id = operator_groups.id AND gtl.deleted = false").
		Joins("LEFT JOIN tags ON tags.id = gtl.tag_id AND tags.deleted = false").
		Where("operator_groups.deleted = false AND operator_groups.id = ?", id).
		Scan(&tags)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tags, nil
}

func (s *Service) AddTagsToGroup(groupIdStr string, tagIdsStr []string) error {
	groupId, err := uuid.Parse(groupIdStr)
	if err != nil {
		return errors.New("invalid group id")
	}
	tagIds := make([]uuid.UUID, len(tagIdsStr))
	for i, tagIdStr := range tagIdsStr {
		tagId, err := uuid.Parse(tagIdStr)
		if err != nil {
			return errors.New("invalid tag id")
		}
		tagIds[i] = tagId
	}
	err = s.DB.Transaction(func(tx *gorm.DB) error {
		for _, tagId := range tagIds {
			var groupTagLink models.GroupTagLink
			groupTagLink.GroupId = groupId
			groupTagLink.TagId = tagId
			groupTagLink.CreatedAt = time.Now()
			groupTagLink.UpdatedAt = time.Now()
			tx.Create(&groupTagLink)
			if tx.Error != nil {
				return tx.Error
			}
		}
		return nil
	})
	if err != nil {
		return errors.New("failed to add tags to group")
	}
	return nil
}

func (s *Service) RemoveTagsFromGroup(groupIdStr string, tagIdsStr []string) error {
	groupId, err := uuid.Parse(groupIdStr)
	if err != nil {
		return errors.New("invalid group id")
	}
	tagIds := make([]uuid.UUID, len(tagIdsStr))
	for i, tagIdStr := range tagIdsStr {
		tagId, err := uuid.Parse(tagIdStr)
		if err != nil {
			return errors.New("invalid tag id")
		}
		tagIds[i] = tagId
	}
	tx := s.DB.Model(&models.GroupTagLink{}).
		Where("group_id = ?", groupId).
		Where("tag_id IN ?", tagIds).
		Updates(map[string]interface{}{
			"deleted":    true,
			"updated_at": time.Now(),
		})
	if tx.Error != nil {
		return errors.New("failed to remove tags from group")
	}
	return nil
}
