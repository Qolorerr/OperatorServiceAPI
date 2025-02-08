package models

import (
	"github.com/google/uuid"
	"time"
)

type Operator struct {
	ID      uuid.UUID `json:"id" gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4()"`
	Deleted bool      `json:"deleted" gorm:"column:deleted;type:boolean;default:false"`

	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`

	OperatorGroupLinks []OperatorGroupLink `json:"-" gorm:"foreignKey:operator_id"`
}
