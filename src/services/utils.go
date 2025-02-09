package services

import "operator_text_channel/src/storage/models"

func (s *Service) loadAppealTags(appeal *Appeal) error {
	var tags []Tag
	tx := s.DB.Model(&models.AppealTagLink{}).
		Distinct("tags.*").
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
		Distinct("tags.*").
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
	exportAppeals := make([]Appeal, len(appeals))
	for i, appeal := range appeals {
		exportAppeal := Appeal{
			ID:        appeal.ID,
			UserId:    appeal.UserId,
			Weight:    appeal.Weight,
			CreatedAt: appeal.CreatedAt,
			UpdatedAt: appeal.UpdatedAt,
		}
		for _, link := range appeal.AppealTagLinks {
			exportAppeal.Tags = append(exportAppeal.Tags, Tag{link.Tag.ID, link.Tag.Name, link.CreatedAt, link.UpdatedAt})
		}
		exportAppeals[i] = exportAppeal
	}
	return exportAppeals
}

func convertGroupsToExport(groups []models.OperatorGroup) []OperatorGroup {
	exportGroups := make([]OperatorGroup, len(groups))
	for i, group := range groups {
		exportGroup := OperatorGroup{ID: group.ID, CreatedAt: group.CreatedAt, UpdatedAt: group.UpdatedAt}
		for _, link := range group.GroupTagLinks {
			exportGroup.Tags = append(exportGroup.Tags, Tag{link.Tag.ID, link.Tag.Name, link.CreatedAt, link.UpdatedAt})
		}
		for _, link := range group.OperatorGroupLinks {
			exportGroup.Operators = append(exportGroup.Operators, Operator{link.Operator.ID, link.CreatedAt, link.UpdatedAt})
		}
		exportGroups[i] = exportGroup
	}
	return exportGroups
}
