package entities

import (
	"gorm.io/gorm"
	"time"
)

type Game struct {
	ID        int `json:"id,omitempty" gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
