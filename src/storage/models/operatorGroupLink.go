package models

import (
	"github.com/google/uuid"
	"time"
)

type OperatorGroupLink struct {
	ID            uuid.UUID     `json:"-" gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4()"`
	Deleted       bool          `json:"-" gorm:"column:deleted;type:boolean;default:false"`
	OperatorId    uuid.UUID     `json:"-" gorm:"column:operator_id;type:uuid;not null"`
	Operator      Operator      `json:"-" gorm:"foreignKey:operator_id;references:id"`
	GroupId       uuid.UUID     `json:"-" gorm:"column:group_id;type:uuid;not null"`
	OperatorGroup OperatorGroup `json:"-" gorm:"foreignKey:group_id;references:id"`

	CreatedAt time.Time `json:"-" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at;type:timestamp;not null"`
}
