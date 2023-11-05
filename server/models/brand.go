package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time		 `json:"created_at"`
	UpdatedAt time.Time		 `json:"updated_at"`
  }

type Brand struct {
	ID        int    `gorm:"primaryKey;type:smallint" json:"id"`
	Name      string `gorm:"type:varchar(100)" json:"name"`
	gorm.Model
}