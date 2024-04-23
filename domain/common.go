package domain

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint         `gorm:"primarykey" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"deleted_at"`
}

type ApplicationType interface {
	GetDB() *gorm.DB
}
