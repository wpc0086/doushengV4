package controller

import (
	"bytes"
	"doushengV4/cmd/api/rpc"
	"doushengV4/kitex_gen/publish"
	"doushengV4/pkg/consts"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

func Publish(c *gin.Context) {
	r := new(publish.ActionRequest)
	r.Token = c.PostForm("token")
	r.Title = c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		})
		return
	}
	//判断视频类型和大小
	if strings.Split(data.Filename, ".")[1] != "mp4" {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ParamErrCode),
			StatusMsg:  "请上传mp4文件",
		})
		return
	}
	if data.Size > consts.VideoMaxSize {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ParamErrCode),
			StatusMsg:  "视频文件过大",
		})
		return
	}

	err, bs := FileToBytes(c.Copy(), data)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		})
		return
	}
	r.Data = bs
	response, err := rpc.ActionPublic(c.Copy(), r)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func FileToBytes(c *gin.Context, data *multipart.FileHeader) (error, []byte) {
	buff := new(bytes.Buffer)
	open, err := data.Open()
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		})
		return err, nil
	}
	_, err = io.Copy(buff, open)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		})
		return err, nil
	}
	bs := buff.Bytes()
	return nil, bs
}

func PublishList(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		})
		return
	}
	r := new(publish.ListRequest)
	r.UserId = uid
	r.Token = c.Query("token")
	response, err := rpc.ListPublish(c.Copy(), r)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func Feed(c *gin.Context) {
	lTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: int32(publish.ErrCode_ServiceErrCode),
			StatusMsg:  err.Error(),
		})
		return
	}
	r := new(publish.FeedRequest)
	r.LatestTime = &lTime
	token := c.Query("token")
	r.Token = &token
	response, err := rpc.FeedPublish(c.Copy(), r)
	if err != nil {
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
