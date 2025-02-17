// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameAudit = "audits"

// Audit mapped from table <audits>
type Audit struct {
	ID            int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserType      string    `gorm:"column:user_type" json:"user_type"`
	UserID        int64     `gorm:"column:user_id" json:"user_id"`
	Event         string    `gorm:"column:event;not null" json:"event"`
	AuditableType string    `gorm:"column:auditable_type;not null" json:"auditable_type"`
	AuditableID   int64     `gorm:"column:auditable_id;not null" json:"auditable_id"`
	OldValues     string    `gorm:"column:old_values" json:"old_values"`
	NewValues     string    `gorm:"column:new_values" json:"new_values"`
	URL           string    `gorm:"column:url" json:"url"`
	IPAddress     string    `gorm:"column:ip_address" json:"ip_address"`
	UserAgent     string    `gorm:"column:user_agent" json:"user_agent"`
	Tags          string    `gorm:"column:tags" json:"tags"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName Audit's table name
func (*Audit) TableName() string {
	return TableNameAudit
}
