package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/xueqiya/go_project/utils/errno"
	"net/http"
)

// Resp 统一返回格式
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response 根据数据返回响应
func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Resp{Code: httpCode, Msg: errno.Msg[errCode], Data: data})
}

// BindAndValid 绑定并验证表单
func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	// c.Bind(form) 会根据 Content-Type 选择 binding
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, errno.InvalidParams
	} else {
		return http.StatusOK, errno.Success
	}
}
