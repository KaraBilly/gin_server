package user

import (
	. "gin_server/handler"
	"gin_server/model"
	"gin_server/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"username":"xx"}"
// @Router /user [post]
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//// Validate the data.
	//if err := u.Validate(); err != nil {
	//	SendResponse(c, errno.ErrValidation, nil)
	//	return
	//}
	//
	//// Encrypt the user password.
	//if err := u.Encrypt(); err != nil {
	//	SendResponse(c, errno.ErrEncrypt, nil)
	//	return
	//}
	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
