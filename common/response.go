package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	ERROR   = 405
	SUCCESS = 200
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// Success 固定 提示信息
func Success(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

// SuccessMsg 自定义提示信息
func SuccessMsg(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

// SuccessDataMsg 返回数据，自定义提示信息
func SuccessDataMsg(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}
func FailMsg(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}
