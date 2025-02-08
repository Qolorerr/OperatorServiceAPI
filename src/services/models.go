package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	DB  *gorm.DB
	CTX context.Context
	RDB *redis.Client
}

func NewDBService(ctx context.Context, db *gorm.DB, rdb *redis.Client) *Service {
	return &Service{CTX: ctx, DB: db, RDB: rdb}
}

type Tag struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	Name      string    `json:"name" gorm:"column:name"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type Appeal struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	UserId    uuid.UUID `json:"userId" gorm:"column:user_id"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	Tags      []Tag     `json:"tags" gorm:"-"`
}

type Operator struct {
	ID        uuid.UUID `json:"id" gorm:"column:id"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
}

type OperatorGroup struct {
	ID        uuid.UUID  `json:"id" gorm:"column:id"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"column:updated_at"`
	Tags      []Tag      `json:"tags" gorm:"-"`
	Operators []Operator `json:"operators" gorm:"-"`
}
