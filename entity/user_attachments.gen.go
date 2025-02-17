// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserAttachment = "user_attachments"

// UserAttachment mapped from table <user_attachments>
type UserAttachment struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID    int64          `gorm:"column:user_id" json:"user_id"`
	URL       string         `gorm:"column:url;not null" json:"url"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName UserAttachment's table name
func (*UserAttachment) TableName() string {
	return TableNameUserAttachment
}
