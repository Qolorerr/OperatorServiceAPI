package services

import "operator_text_channel/src/storage/models"

func (s *Service) loadAppealTags(appeal *Appeal) error {
	var tags []Tag
	tx := s.DB.Model(&models.AppealTagLink{}).
		Select("tags.*").
		Joins("LEFT JOIN tags ON tags.id = appeal_tag_links.tag_id AND tags.deleted = false").
		Where("appeal_tag_links.deleted = false AND appeal_tag_links.appeal_id = ?", appeal.ID).
		Scan(&tags)
	if tx.Error != nil {
		return tx.Error
	}
	appeal.Tags = tags
	return nil
}

func (s *Service) loadGroupTags(group *OperatorGroup) error {
	var tags []Tag
	tx := s.DB.Model(&models.GroupTagLink{}).
		Select("tags.*").
		Joins("LEFT JOIN tags ON tags.id = group_tag_links.tag_id AND tags.deleted = false").
		Where("group_tag_links.deleted = false AND group_tag_links.group_id = ?", group.ID).
		Scan(&tags)
	if tx.Error != nil {
		return tx.Error
	}
	group.Tags = tags
	return nil
}

func convertAppealsToExport(appeals []models.Appeal) []Appeal {
	var exportAppeals []Appeal
	for _, appeal := range appeals {
		exportAppeal := Appeal{ID: appeal.ID, UserId: appeal.UserId, CreatedAt: appeal.CreatedAt, UpdatedAt: appeal.UpdatedAt}
		for _, link := range appeal.AppealTagLinks {
			exportAppeal.Tags = append(exportAppeal.Tags, Tag{link.Tag.ID, link.Tag.Name, link.CreatedAt, link.UpdatedAt})
		}
		exportAppeals = append(exportAppeals, exportAppeal)
	}
	return exportAppeals
}

func convertGroupsToExport(groups []models.OperatorGroup) []OperatorGroup {
	var exportGroups []OperatorGroup
	for _, group := range groups {
		exportGroup := OperatorGroup{ID: group.ID, CreatedAt: group.CreatedAt, UpdatedAt: group.UpdatedAt}
		for _, link := range group.GroupTagLinks {
			exportGroup.Tags = append(exportGroup.Tags, Tag{link.Tag.ID, link.Tag.Name, link.CreatedAt, link.UpdatedAt})
		}
		for _, link := range group.OperatorGroupLinks {
			exportGroup.Operators = append(exportGroup.Operators, Operator{link.Operator.ID, link.CreatedAt, link.UpdatedAt})
		}
		exportGroups = append(exportGroups, exportGroup)
	}
	return exportGroups
}
