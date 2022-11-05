package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/controllers"
)

func setupPostRouter(router *gin.RouterGroup, db *sqlx.DB, redisConn *redis.Client) {
	controller := controllers.NewPostController(db, redisConn)
	router.GET("/actual", controller.GetApprovedPosts)
	router.GET("/needToApprove", controller.GetPostsNeedToApprove)
	router.GET("/get", controller.GetApprovedPost)
	router.POST("/new", controller.CreatePost)
	router.POST("/approve", controller.ApprovePost)
	router.POST("/reject", controller.RejectPost)
}
