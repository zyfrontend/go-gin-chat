package test

import (
	"ginchat/models"
	"ginchat/utils"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGorm() {

	// 迁移 schema
	utils.DB.AutoMigrate(&models.UserBasic{})

	// Create
	//user := &models.UserBasic{}
	//user.Name = "zy"
	//db.Create(user)

	// Read
	//fmt.Println(db.First(user, 1)) // 根据整型主键查找
	// Update - 将 product 的 price 更新为 200
	//db.Model(user).Update("PassWord", "1234")
}
