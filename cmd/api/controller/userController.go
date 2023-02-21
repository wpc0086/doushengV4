package controller

import (
	"doushengV4/cmd/api/middleware"
	"doushengV4/cmd/api/rpc"
	"doushengV4/kitex_gen/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	u := new(user.RegisterUserRequest)
	u.Username = c.Query("username")
	u.Password = c.Query("password")
	response, err := rpc.RegisterUser(c.Copy(), u)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context) {
	u := new(user.LoginUserRequest)
	u.Username = c.Query("username")
	u.Password = c.Query("password")
	response, err := rpc.LoginUser(c.Copy(), u)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func UserInfo(c *gin.Context) {
	u := new(user.InfoUserRequest)
	uid, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
		return
	}
	u.UserId = uid
	response, err := rpc.InfoUser(c.Copy(), u)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
