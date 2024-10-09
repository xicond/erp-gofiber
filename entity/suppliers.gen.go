// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSupplier = "suppliers"

// Supplier mapped from table <suppliers>
type Supplier struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID        string         `gorm:"column:uuid;not null" json:"uuid"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	Tel         string         `gorm:"column:tel" json:"tel"`
	Email       string         `gorm:"column:email" json:"email"`
	Address     string         `gorm:"column:address" json:"address"`
	Remark      string         `gorm:"column:remark" json:"remark"`
	UpdatedBy   int64          `gorm:"column:updated_by" json:"updated_by"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	ContactName string         `gorm:"column:contact_name" json:"contact_name"`
	TenantID    int64          `gorm:"column:tenant_id" json:"tenant_id"`
	Shortname   string         `gorm:"column:shortname" json:"shortname"`
}

// TableName Supplier's table name
func (*Supplier) TableName() string {
	return TableNameSupplier
}
