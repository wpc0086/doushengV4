package controller

import (
	"doushengV4/cmd/api/middleware"
	"doushengV4/cmd/api/rpc"
	"doushengV4/kitex_gen/interact"
	"doushengV4/kitex_gen/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	vid := c.Query("video_id")
	videoId, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
		return
	}
	act := c.Query("action_type")
	actionType, err := strconv.ParseInt(act, 10, 32)
	if err != nil {
		c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
		return
	}
	f := new(interact.FavoriteActionRequest)
	f.Token = token
	f.VideoId = videoId
	f.ActionType = int32(actionType)

	response, err := rpc.FavoriteAction(c.Copy(), f)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	userId, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
		return
	}
	f := new(interact.FavoriteListRequest)
	f.Token = token
	f.UserId = userId
	response, err := rpc.FavoriteList(c.Copy(), f)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
