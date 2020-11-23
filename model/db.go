package model

import "time"

// User - tables of users
type User struct {
	ID          int        `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Age         int        `gorm:"column:age;size:3" json:"age,omitempty"`
	Fname       string     `gorm:"column:fname;not null" json:"fname"`
	Lname       string     `gorm:"column:lname;not null" json:"lname"`
	Gender      string     `gorm:"column:gender;not null" json:"gender"`
	Address     string     `gorm:"column:address" json:"address,omitempty"`
	Email       string     `gorm:"column:email" json:"email,omitempty"`
	Phone       string     `gorm:"column:phone" json:"phone,omitempty"`
	Description string     `gorm:"column:description" json:"description,omitempty"`
	Created     *time.Time `gorm:"column:created;autoCreateTime" json:"-"`
}

// TableName - name of table
func (User) TableName() string {
	return "users"
}
