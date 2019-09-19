package wechat

import (
	. "gin_server/handler"
	"gin_server/pkg/errno"
	"gin_server/service"
	"github.com/gin-gonic/gin"
)

// @Summary Get wechat access token
// @Description Get access token from middle-platform
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} wechat.GetTokenResponse "{"token":"xx"}"
// @Router /v1/wechat/token [get]
func GetToken(c *gin.Context) {
	token, err := service.GetAccessToken()
	if err != nil {
		SendResponse(c, errno.ErrWeChatRequest, nil)
		return
	}
	rsp := GetTokenResponse{
		Token: token,
	}
	SendResponse(c, nil, rsp)
}

func SendMessage(c *gin.Context) {
	var r SendMessageRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}
	//service.SendMessage(&r)
}
