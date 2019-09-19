package handler

import (
	"net/http"

	"gin_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {

	if err != nil {
		code, message := errno.DecodeErr(err)
		if code == http.StatusInternalServerError {
			c.JSON(code, Response{
				Code:    code,
				Message: message,
				Data:    data,
			})
		}
		c.JSON(http.StatusBadRequest, Response{
			Code:    code,
			Message: message,
			Data:    data,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
