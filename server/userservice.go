package server

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @BasePath /api

// GetUserList
// @Tags 用户列表
// @Success 200 {string} data
// @Router /user/list [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

// CreateUser
// @Tags 用户列表
// @Success 200 {string} data
// @Router /user/create [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")
	if password != repassword {
		c.JSON(405, gin.H{
			"message": "两次密码不一致",
		})
	}
	user.PassWord = password
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增用户成功",
	})
}

// DeleteUser
// @Tags 用户列表
// @Success 200 {string} data
// @Router /user/delete [post]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除用户成功",
	})
}
