package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name      string         `json:"name"`
	Content   string         `json:"content"`
}

type ToDos []Todo
