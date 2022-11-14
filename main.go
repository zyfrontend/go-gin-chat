package main

import (
	"ginchat/router"
	"ginchat/sql"
	"ginchat/utils"
)

// @title GO-Gin-Chat-API
// @version 1.0 版本
// @description 聊天系统 API文档
// @BasePath /api
// @query.collection.format multi
func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	sql.CreateTable()
	r := router.Routers()
	r.Run(":9797")
}
