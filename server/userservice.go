package server

import (
	"fmt"
	"ginchat/common"
	"ginchat/models"
	"ginchat/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// @BasePath /api

// FindUserByNameAndPassword
// @Summary 用户登录
// @Tags 用户模块
// @Param name query string false "用户名"
// @Param password query string false "密码"
// @Success 200 {string} data
// @Router /login [post]
func FindUserByNameAndPassword(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	// 查询用户
	user := models.FindUserByName(name)
	if user.Name == "" {
		common.FailMsg("该用户不存在", c)
		return
	}
	boolStatus := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !boolStatus {
		common.FailMsg("密码不正确", c)
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPassword(name, pwd)
	common.SuccessDataMsg(data, "登录成功", c)
}

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} data
// @Router /user/list [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	common.SuccessDataMsg(data, "查询成功", c)
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
		common.FailMsg("用户名不能为空", c)
		return
	} else if password == "" {
		common.FailMsg("密码不能为空", c)
	} else if repassword == "" {
		common.FailMsg("验证密码不能为空", c)
		return
	} else if password != repassword {
		common.FailMsg("两次密码不一致", c)
		return
	}

	// 查询用户
	data := models.FindUserByName(name)
	// 不等于空，表示存在用户
	if data.Name != "" {
		common.FailMsg("用户名已注册", c)
		return
	}
	// 密码加密
	salt := fmt.Sprintf("%06d", rand.Int31())

	user.Name = name
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt
	models.CreateUser(user)
	common.SuccessDataMsg(&user, "新增用户成功", c)
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
		common.FailMsg("找不到用户ID", c)
		return
	}
	models.DeleteUser(user)
	common.SuccessMsg("删除用户成功", c)
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
	if strconv.Itoa(id) == "" {
		common.FailMsg("用户id不能为空", c)
		return
	} else if name == "" {
		common.FailMsg("用户名不能为空", c)
		return
	} else if password == "" {
		common.FailMsg("密码不能为空", c)
		return
	}
	// 查找用户
	data := models.FindUserByID(id)
	if data.Name == "" {
		common.FailMsg("找不到用户ID", c)
		return
	}

	user.ID = uint(id)
	user.Name = name
	user.PassWord = password
	models.UpdateUser(user)
	common.SuccessMsg("修改用户成功", c)
}

// UserInfo
// @Summary 用户信息
// @Tags 用户模块
// @Param id formData string false "用户id"
// @Success 200 {string} json{"code", "message"}
// @Router /user/info [get]
func UserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	if strconv.Itoa(id) == "" {
		common.FailMsg("用户id不能为空", c)
		return
	}
	// 查找用户
	data := models.FindUserByID(id)
	if data.Name == "" {
		common.FailMsg("找不到用户ID", c)
		return
	}
	common.SuccessDataMsg(&data, "查询成功", c)
}

// 处理跨域
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)

	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}

}
