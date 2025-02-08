package models

import (
	"github.com/google/uuid"
	"time"
)

type Tag struct {
	ID      uuid.UUID `json:"id" gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4()"`
	Deleted bool      `json:"-" gorm:"column:deleted;type:boolean;default:false"`
	Name    string    `json:"name" gorm:"column:name;not null;unique"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:timestamp;not null"`

	GroupTagLinks  []GroupTagLink  `json:"-" gorm:"foreignKey:tag_id"`
	AppealTagLinks []AppealTagLink `json:"-" gorm:"foreignKey:tag_id"`
}
