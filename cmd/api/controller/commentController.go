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

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64) //将字符串转换为int64
	if err != nil {
		c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
		return
	}
	actionType, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	comment := &interact.CommentActionRequest{Token: token, VideoId: videoId, ActionType: int32(actionType)}
	if err != nil {
		c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
		return
	}
	if actionType == 1 {
		commentText := c.Query("comment_text")
		comment.CommentText = commentText
	} else if actionType == 2 {
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
			return
		}
		comment.CommentId = commentId
	} else {
		c.JSON(http.StatusOK, middleware.Response{StatusCode: int32(user.ErrCode_ParamErrCode)})
		return
	}
	response, err := rpc.CommentAction(c.Copy(), comment)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func CommentList(c *gin.Context) {
	vid := c.Query("video_id")
	video_id, _ := strconv.ParseInt(vid, 10, 64)
	response, err := rpc.CommentList(c.Copy(), &interact.CommentListRequest{VideoId: video_id})
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
