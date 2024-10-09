// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePurchaseOrder = "purchase_orders"

// PurchaseOrder mapped from table <purchase_orders>
type PurchaseOrder struct {
	ID                   int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID                 string         `gorm:"column:uuid;not null" json:"uuid"`
	DisplayID            string         `gorm:"column:display_id" json:"display_id"`
	Status               string         `gorm:"column:status" json:"status"`
	SupplierID           int64          `gorm:"column:supplier_id" json:"supplier_id"`
	TargetRestockDate    time.Time      `gorm:"column:target_restock_date" json:"target_restock_date"`
	Remark               string         `gorm:"column:remark" json:"remark"`
	CreatedBy            int64          `gorm:"column:created_by" json:"created_by"`
	ConfirmedBy          int64          `gorm:"column:confirmed_by" json:"confirmed_by"`
	ConfirmedAt          time.Time      `gorm:"column:confirmed_at" json:"confirmed_at"`
	DeletedAt            gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	CreatedAt            time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt            time.Time      `gorm:"column:updated_at" json:"updated_at"`
	PoDate               time.Time      `gorm:"column:po_date" json:"po_date"`
	ContactName          string         `gorm:"column:contact_name" json:"contact_name"`
	ContactTel           string         `gorm:"column:contact_tel" json:"contact_tel"`
	ContactEmail         string         `gorm:"column:contact_email" json:"contact_email"`
	DeliveryAddress      string         `gorm:"column:delivery_address" json:"delivery_address"`
	SupplierContactName  string         `gorm:"column:supplier_contact_name" json:"supplier_contact_name"`
	SupplierContactTel   string         `gorm:"column:supplier_contact_tel" json:"supplier_contact_tel"`
	SupplierContactEmail string         `gorm:"column:supplier_contact_email" json:"supplier_contact_email"`
	ReceiveRemark        string         `gorm:"column:receive_remark" json:"receive_remark"`
	TenantID             int64          `gorm:"column:tenant_id" json:"tenant_id"`
}

// TableName PurchaseOrder's table name
func (*PurchaseOrder) TableName() string {
	return TableNamePurchaseOrder
}
