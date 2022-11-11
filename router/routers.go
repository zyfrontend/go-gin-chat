package router

import (
	"ginchat/docs"
	"ginchat/server"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	docs.SwaggerInfo.BasePath = "/api"
	v1 := r.Group("/api")
	{
		v1.GET("/index", server.GetIndex)
		v2 := v1.Group("/user")
		{
			v2.GET("/list", server.GetUserList)
			v2.POST("/create", server.CreateUser)
			v2.POST("/delete", server.DeleteUser)
		}
	}
	return r
}
