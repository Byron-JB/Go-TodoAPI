package models

import "time"

type TblTodo struct {
	ID             uint `gorm:"primaryKey"`
	StrTitle       string
	StrDescription string
	IPriority      *int
	DtCompleted    *time.Time
	DtCreatedAt    time.Time
	DtDueDate      time.Time
}

// TableName overrides default table name
func (TblTodo) TableName() string {
	return "tbltodos"
}
