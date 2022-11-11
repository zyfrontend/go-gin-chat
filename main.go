package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Routers()
	r.Run(":9797")
}
