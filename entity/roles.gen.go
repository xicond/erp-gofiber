// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRole = "roles"

// Role mapped from table <roles>
type Role struct {
	ID          int32          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	Slug        string         `gorm:"column:slug;not null" json:"slug"`
	Description string         `gorm:"column:description" json:"description"`
	Level       int32          `gorm:"column:level;not null;default:1" json:"level"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	IsDutyRole  bool           `gorm:"column:is_duty_role;not null" json:"is_duty_role"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}
