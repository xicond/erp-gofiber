// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameUserExpotoken = "user_expotoken"

// UserExpotoken mapped from table <user_expotoken>
type UserExpotoken struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID     int64     `gorm:"column:user_id;not null" json:"user_id"`
	ExpoToken  string    `gorm:"column:expo_token;not null" json:"expo_token"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeviceType string    `gorm:"column:device_type;not null;default:android" json:"device_type"`
	InvalidAt  time.Time `gorm:"column:invalid_at" json:"invalid_at"`
}

// TableName UserExpotoken's table name
func (*UserExpotoken) TableName() string {
	return TableNameUserExpotoken
}
