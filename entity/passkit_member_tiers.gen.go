// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePasskitMemberTier = "passkit_member_tiers"

// PasskitMemberTier mapped from table <passkit_member_tiers>
type PasskitMemberTier struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Passkit       string         `gorm:"column:passkit;not null" json:"passkit"`
	PasskitTierID int32          `gorm:"column:passkit_tier_id;not null" json:"passkit_tier_id"`
	Point         int32          `gorm:"column:point;not null" json:"point"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName PasskitMemberTier's table name
func (*PasskitMemberTier) TableName() string {
	return TableNamePasskitMemberTier
}
