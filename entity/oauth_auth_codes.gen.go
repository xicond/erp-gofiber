// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameOauthAuthCode = "oauth_auth_codes"

// OauthAuthCode mapped from table <oauth_auth_codes>
type OauthAuthCode struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	UserID    int64     `gorm:"column:user_id;not null" json:"user_id"`
	ClientID  int64     `gorm:"column:client_id;not null" json:"client_id"`
	Scopes    string    `gorm:"column:scopes" json:"scopes"`
	Revoked   bool      `gorm:"column:revoked;not null" json:"revoked"`
	ExpiresAt time.Time `gorm:"column:expires_at" json:"expires_at"`
}

// TableName OauthAuthCode's table name
func (*OauthAuthCode) TableName() string {
	return TableNameOauthAuthCode
}
