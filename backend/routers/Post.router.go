package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/controllers"
	"github.com/lxndrrud/blog-project/infrastructure"
	"github.com/lxndrrud/blog-project/repositories"
	"github.com/lxndrrud/blog-project/services"
	"github.com/lxndrrud/blog-project/utils"
)

func setupPostRouter(router *gin.RouterGroup, db *sqlx.DB, redisConn *redis.Client) {
	postRepo := repositories.NewPostsRepo(db)
	userRepo := repositories.NewUserRepo(db)
	permissionInfra := infrastructure.NewPermissionInfra(db, redisConn)
	fileProcessor := utils.NewFileProcessor()
	permissionChecker := utils.NewPermissionChecker()

	postService := services.NewPostService(
		postRepo,
		userRepo,
		permissionInfra,
		fileProcessor,
		permissionChecker,
	)

	controller := controllers.NewPostController(postService)

	router.GET("/userPosts", controller.GetUserPosts)
	router.GET("/actual", controller.GetApprovedPosts)
	router.GET("/needToApprove", controller.GetPostsNeedToApprove)
	router.GET("/get", controller.GetApprovedPost)
	router.GET("/needToApprove/get", controller.GetPostNeedToApprove)
	router.POST("/new", controller.CreatePost)
	router.POST("/approve", controller.ApprovePost)
	router.POST("/reject", controller.RejectPost)
	router.DELETE("/delete", controller.DeletePost)
}
