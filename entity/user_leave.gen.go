// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUserLeave = "user_leave"

// UserLeave mapped from table <user_leave>
type UserLeave struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int64     `gorm:"column:user_id;not null" json:"user_id"`
	LeaveID   int64     `gorm:"column:leave_id;not null" json:"leave_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName UserLeave's table name
func (*UserLeave) TableName() string {
	return TableNameUserLeave
}
