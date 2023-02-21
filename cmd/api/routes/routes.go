package routes

import (
	"doushengV4/cmd/api/controller"
	"doushengV4/cmd/api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// -----basic apis---基础接口
	//feed
	apiRouter.GET("/feed/", controller.Feed)

	//user路由组
	apiRouter.GET("/user/", middleware.JwtMiddleware(), controller.UserInfo) //获取用户信息需要验证token
	apiRouter.POST("/user/register/", controller.Register)                   //需要颁发token
	apiRouter.POST("/user/login/", controller.Login)                         //需要颁发token

	//publish路由组
	apiRouter.POST("/publish/action/", middleware.JwtMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.JwtMiddleware(), controller.PublishList)

	//-----------------

	// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JwtMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JwtMiddleware(), controller.FavoriteList)
	//
	apiRouter.POST("/comment/action/", middleware.JwtMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.JwtMiddleware(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)

}
