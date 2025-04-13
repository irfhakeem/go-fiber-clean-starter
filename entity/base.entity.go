package entity

import (
	"time"

	"gorm.io/gorm"
)

type TimeStamps struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp with time zone"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp with time zone"`
	IsDeleted bool      `json:"is_deleted" gorm:"default:false"`
	DeletedAt gorm.DeletedAt
}
