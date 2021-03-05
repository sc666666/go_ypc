package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 统一返回格式
func response(httpStatusCode int, msg string, data interface{}, c *gin.Context) {
	c.JSON(httpStatusCode, gin.H{
		"code": httpStatusCode,
		"msg":  msg,
		"data": data,
	})
}

// 返回 200
func Ok(data interface{}, c *gin.Context) {
	msg := "请求成功"
	response(http.StatusOK, msg, data, c)
}

// 返回 400
func BadRequest(data interface{}, c *gin.Context) {
	msg := "请求失败"
	response(http.StatusBadRequest, msg, data, c)
}

// 返回 422
func UnprocessableEntity(data interface{}, c *gin.Context) {
	msg := "请求参数验证失败"
	response(http.StatusUnprocessableEntity, msg, data, c)
}
