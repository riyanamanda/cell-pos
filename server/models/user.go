package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"index;type:varchar(25)" json:"name"`
	Email     string `gorm:"index;type:varchar(25)" json:"email"`
	Password  string `gorm:"type:varchar(50)" json:"password"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}