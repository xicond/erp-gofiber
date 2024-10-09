// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNamePosCart = "pos_carts"

// PosCart mapped from table <pos_carts>
type PosCart struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	StoreID   int64     `gorm:"column:store_id" json:"store_id"`
	UUID      string    `gorm:"column:uuid;not null" json:"uuid"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	Email     string    `gorm:"column:email" json:"email"`
	Tel       string    `gorm:"column:tel" json:"tel"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName PosCart's table name
func (*PosCart) TableName() string {
	return TableNamePosCart
}
