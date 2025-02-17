// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAttendanceCredit = "attendance_credits"

// AttendanceCredit mapped from table <attendance_credits>
type AttendanceCredit struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID            string         `gorm:"column:uuid;not null" json:"uuid"`
	Time            time.Time      `gorm:"column:time;not null" json:"time"`
	UserID          int64          `gorm:"column:user_id;not null" json:"user_id"`
	Type            string         `gorm:"column:type;not null;default:late_adjustment" json:"type"`
	AdjustmentValue float64        `gorm:"column:adjustment_value;not null;default:0.00" json:"adjustment_value"`
	Remark          string         `gorm:"column:remark" json:"remark"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AttendanceCredit's table name
func (*AttendanceCredit) TableName() string {
	return TableNameAttendanceCredit
}
