// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNamePasskitTier = "passkit_tiers"

// PasskitTier mapped from table <passkit_tiers>
type PasskitTier struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProgramID string `gorm:"column:program_id;not null" json:"program_id"`
	TierID    string `gorm:"column:tier_id;not null" json:"tier_id"`
	Remark    string `gorm:"column:remark" json:"remark"`
}

// TableName PasskitTier's table name
func (*PasskitTier) TableName() string {
	return TableNamePasskitTier
}
