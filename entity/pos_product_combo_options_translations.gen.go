// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePosProductComboOptionsTranslation = "pos_product_combo_options_translations"

// PosProductComboOptionsTranslation mapped from table <pos_product_combo_options_translations>
type PosProductComboOptionsTranslation struct {
	ID                      int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PosProductComboOptionID int64          `gorm:"column:pos_product_combo_option_id;not null" json:"pos_product_combo_option_id"`
	Locale                  string         `gorm:"column:locale;not null;default:en" json:"locale"`
	Name                    string         `gorm:"column:name;not null" json:"name"`
	DeletedAt               gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt               time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt               time.Time      `gorm:"column:updated_at" json:"updated_at"`
}

// TableName PosProductComboOptionsTranslation's table name
func (*PosProductComboOptionsTranslation) TableName() string {
	return TableNamePosProductComboOptionsTranslation
}
