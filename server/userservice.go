package server

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @BasePath /api

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
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
// @Summary 创建用户
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Param repassword query string false "验证密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/create [post]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	if name == "" {
		c.JSON(405, gin.H{
			"message": "用户名不能为空",
		})
		return
	} else if password == "" {
		c.JSON(405, gin.H{
			"message": "密码不能为空",
		})
		return
	} else if repassword == "" {
		c.JSON(405, gin.H{
			"message": "验证密码不能为空",
		})
		return
	} else if password != repassword {
		c.JSON(405, gin.H{
			"message": "两次密码不一致",
		})
		return
	}

	// 查询用户
	data := models.FindUserByName(name)
	// 不等于空，表示存在用户
	if data.Name != "" {
		c.JSON(405, gin.H{
			"message": "用户名已注册",
		})
		return
	}

	user.Name = name
	user.PassWord = password
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增用户成功",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Param id query string false "用户id"
// @Success 200 {string} data
// @Router /user/delete [post]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	data := models.FindUserByID(id)
	if data.Name == "" {
		c.JSON(405, gin.H{
			"message": "找不到用户ID",
		})
		return
	}
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "删除用户成功",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @Param id formData string false "用户id"
// @Param name formData string false "新用户名"
// @Param password formData string false "新密码"
// @Success 200 {string} json{"code", "message"}
// @Router /user/update [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	name := c.PostForm("name")
	password := c.PostForm("password")
	if string(id) == "" {
		c.JSON(405, gin.H{
			"message": "用户id不能为空",
		})
		return
	} else if name == "" {
		c.JSON(405, gin.H{
			"message": "用户名不能为空",
		})
		return
	} else if password == "" {
		c.JSON(405, gin.H{
			"message": "密码不能为空",
		})
		return
	}
	user.ID = uint(id)
	user.Name = name
	user.PassWord = password
	data := models.FindUserByID(id)
	if data.Name == "" {
		c.JSON(405, gin.H{
			"message": "找不到用户ID",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "修改用户成功",
	})
}
