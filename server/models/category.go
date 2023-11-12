package models

import "time"

type Category struct {
	ID        int64  	`gorm:"primaryKey" json:"id"`
	Name      string 	`gorm:"index;type:varchar(20)" json:"name"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}