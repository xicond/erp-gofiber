// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameOtpAttempt = "otp_attempts"

// OtpAttempt mapped from table <otp_attempts>
type OtpAttempt struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Module    string    `gorm:"column:module;not null" json:"module"`
	EntityID  string    `gorm:"column:entity_id;not null" json:"entity_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName OtpAttempt's table name
func (*OtpAttempt) TableName() string {
	return TableNameOtpAttempt
}
