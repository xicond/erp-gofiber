// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePosOrderItem = "pos_order_items"

// PosOrderItem mapped from table <pos_order_items>
type PosOrderItem struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null" json:"uuid"`
	PosOrderID   int64          `gorm:"column:pos_order_id;not null" json:"pos_order_id"`
	PosProductID int64          `gorm:"column:pos_product_id;not null" json:"pos_product_id"`
	Qty          int32          `gorm:"column:qty;not null;default:1" json:"qty"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	Price        float64        `gorm:"column:price;not null;default:0.00" json:"price"`
	Remark       string         `gorm:"column:remark" json:"remark"`
	ExtraData    string         `gorm:"column:extra_data" json:"extra_data"`
}

// TableName PosOrderItem's table name
func (*PosOrderItem) TableName() string {
	return TableNamePosOrderItem
}
