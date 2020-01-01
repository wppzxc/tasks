package database

import "github.com/jinzhu/gorm"

const (
	UserNormal = "normal"
	UserError  = "error"
)

type Users struct {
	gorm.Model
	Name          string `gorm:"column:name;type:varchar(255);primary_key;not null;unique" json:"name"`
	Password      string `gorm:"column:password;type:varchar(255)" json:"password"`
	AlipayAccount string `gorm:"column:alipay_account;type:varchar(255)" json:"alipayAccount"`
	Status        string `gorm:"column:status;type:varchar(255)" json:"status"`
	Registered    bool   `gorm:"column:registered" json:"registered"`
	Parent        string `gorm:"column:parent" json:"parent"`
	SubUsers      int64  `gorm:"column:sub_users" json:"subUsers"`
	SubTwoUsers   int64  `gorm:"column:sub_two_users" json:"subTwoUsers"`
	// 账户余额
	Balance float32 `gorm:"column:balance" json:"balance"`
}

func (Users) TableName() string {
	return "users"
}
