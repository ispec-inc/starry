// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameDoctorDetail = "doctor_details"

// DoctorDetail mapped from table <doctor_details>
type DoctorDetail struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	DoctorID  string    `gorm:"column:doctor_id;not null" json:"doctor_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName DoctorDetail's table name
func (*DoctorDetail) TableName() string {
	return TableNameDoctorDetail
}
