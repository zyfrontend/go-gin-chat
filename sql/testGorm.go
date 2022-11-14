package sql

import (
	"ginchat/models"
	"ginchat/utils"
)

func CreateTable() {
	// 迁移 schema
	//utils.DB.AutoMigrate(&models.Message{})
	utils.DB.AutoMigrate(&models.GroupBasic{})
	utils.DB.AutoMigrate(&models.Contact{})
}
