package handler

import (
	"net/http"

	"gin_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {

	if err != nil {
		code, message := errno.DecodeErr(err)
		if code == http.StatusInternalServerError {
			c.JSON(code, Response{
				Code:    code,
				Message: message,
			})
		}
		c.JSON(http.StatusBadRequest, Response{
			Code:    code,
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
