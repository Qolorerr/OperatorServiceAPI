package models

import (
	"time"

	"github.com/google/uuid"
)

type Appeal struct {
	ID      uuid.UUID `json:"id" gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4()"`
	Deleted bool      `json:"-" gorm:"column:deleted;type:boolean;default:false"`
	UserId  uuid.UUID `json:"userId" gorm:"column:user_id;type:uuid;not null"`
	Weight  int32     `json:"weight" gorm:"column:weight;type:int;not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;type:timestamp;not null"`

	AppealTagLinks []AppealTagLink `json:"-" gorm:"foreignKey:appeal_id"`
}
