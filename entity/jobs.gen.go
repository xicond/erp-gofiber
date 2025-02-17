// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameJob = "jobs"

// Job mapped from table <jobs>
type Job struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Queue       string `gorm:"column:queue;not null" json:"queue"`
	Payload     string `gorm:"column:payload;not null" json:"payload"`
	Attempts    int32  `gorm:"column:attempts;not null" json:"attempts"`
	ReservedAt  int32  `gorm:"column:reserved_at" json:"reserved_at"`
	AvailableAt int32  `gorm:"column:available_at;not null" json:"available_at"`
	CreatedAt   int32  `gorm:"column:created_at;not null" json:"created_at"`
}

// TableName Job's table name
func (*Job) TableName() string {
	return TableNameJob
}
