package entity

import "time"

type TodoListCategory struct {
	Name        string    `gorm:"column:name"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   int64     `gorm:"column:created_by"`
	ID          int64     `gorm:"column:id"`
}

func (TodoListCategory) TableName() string {
	return "todo_list_categories"
}
