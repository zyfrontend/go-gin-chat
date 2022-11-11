package models

import (
	"ginchat/utils"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name"`
	PassWord      string    `json:"password"`
	Phone         string    `json:"phone"`
	Email         string    `json:"email"`
	Identity      string    `json:"identity"`
	ClientIp      string    `json:"client-ip"`
	ClientPort    string    `json:"client-port"`
	LoginTime     time.Time `json:"login-time"`
	HeartBeatTime time.Time `json:"heart-beat-time"`
	LoginOutTime  time.Time `json:"login-out-time"`
	IsLogout      bool      `json:"is-logout"`
	DeviceInfo    string    `json:"device-info"`
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	return data
}
func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
