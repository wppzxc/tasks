package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

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
	Balance       int64  `gorm:"column:balance" json:"balance"`
}

func (Users) TableName() string {
	return "users"
}

type Tasks struct {
	gorm.Model
	Title            string    `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Description      string    `gorm:"column:Description;type:varchar(255);not null" json:"description"`
	Status           string    `gorm:"column:status;type:varchar(255);not null" json:"status"`
	Expire           time.Time `gorm:"column:expire"json:"expire"`
	GuideImages      string    `gorm:"column:guide_images;type:text;"json:"guideImages"`
	CommentBase      string    `gorm:"column:comment_base;type:varchar(255);not null" json:"commentBase"`
	CommentImages    string    `gorm:"column:comment_images;type:text;"json:"commentImages"`
	Bonus            int64     `gorm:"column:bonus;type:int;" json:"bonus"`
	ParentBonus      int64     `gorm:"column:parent_bonus;type:int;" json:"parentBonus"`
	GrandParentBonus int64     `gorm:"column:grand_parent_bonus;type:int;" json:"grandParentBonus"`
	Steps            string    `gorm:"column:steps;type:text;not null" json:"steps"`
}

func (Tasks) TableName() string {
	return "tasks"
}

type Commits struct {
	gorm.Model
	TaskID      uint   `gorm:"column:task_id;not null" json:"taskID"`
	TaskTitle   string `gorm:"column:task_title;type:varchar(255);" json:"taskTitle"`
	Username    string `gorm:"column:username;type:varchar(255);not null" json:"username"`
	CommentKey  string `gorm:"column:comment_key;type:varchar(255);not null" json:"commentKey"`
	CommitImage string `gorm:"column:commit_image;type:varchar(255)" json:"commitImage"`
	Status      string `gorm:"column:status;type:varchar(255);not null" json:"status"`
}

func (Commits) TableName() string {
	return "commits"
}

type ServiceInfo struct {
	gorm.Model
	ServiceQrCode string `gorm:"column:service_qr_code;type:varchar(255)" json:"serviceQrCode"`
	ServicePhone string `gorm:"column:service_phone;type:varchar(255)" json:"servicePhone"`
	ServiceWechat string `gorm:"column:service_wechat;type:varchar(255)" json:"serviceWechat"`
}

func (ServiceInfo) TableName() string {
	return "service_info"
}
