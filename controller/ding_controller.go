package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xueqiya/zheng_server/model"
	"github.com/xueqiya/zheng_server/utils"
	"github.com/xueqiya/zheng_server/utils/errno"
	"net/http"
)

type Msg struct {
	MsgType string `form:"msgType" valid:"Required;MaxSize(100)"`
	Content string `form:"content" valid:"Required;MaxSize(100)"`
}

func SendSms(c *gin.Context) {
	var form Msg
	httpCode, errCode := utils.BindAndValid(c, &form)
	if errCode != errno.Success {
		utils.Response(c, httpCode, errCode, nil)
		return
	}

	str, err := model.SendSms(form.Content)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, errno.Error, err)
		return
	}

	utils.Response(c, http.StatusOK, errno.Success, str)
}
