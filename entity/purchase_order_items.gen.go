// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNamePurchaseOrderItem = "purchase_order_items"

// PurchaseOrderItem mapped from table <purchase_order_items>
type PurchaseOrderItem struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PurchaseOrderID int64     `gorm:"column:purchase_order_id;not null" json:"purchase_order_id"`
	RawMaterialID   int64     `gorm:"column:raw_material_id" json:"raw_material_id"`
	Qty             float64   `gorm:"column:qty;not null" json:"qty"`
	OriginalPrice   float64   `gorm:"column:original_price;not null;default:0.00" json:"original_price"`
	UnitPrice       float64   `gorm:"column:unit_price;not null;default:0.00" json:"unit_price"`
	Remark          string    `gorm:"column:remark" json:"remark"`
	CreatedBy       int64     `gorm:"column:created_by" json:"created_by"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName PurchaseOrderItem's table name
func (*PurchaseOrderItem) TableName() string {
	return TableNamePurchaseOrderItem
}
