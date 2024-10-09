// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePermissionUser = "permission_user"

// PermissionUser mapped from table <permission_user>
type PermissionUser struct {
	ID           int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PermissionID int32          `gorm:"column:permission_id;not null" json:"permission_id"`
	UserID       int64          `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName PermissionUser's table name
func (*PermissionUser) TableName() string {
	return TableNamePermissionUser
}
