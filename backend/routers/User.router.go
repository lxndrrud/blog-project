package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/controllers"
	"github.com/lxndrrud/blog-project/repositories"
	"github.com/lxndrrud/blog-project/services"
	"github.com/lxndrrud/blog-project/utils"
)

func setupUserRouter(router *gin.RouterGroup, db *sqlx.DB, redisConn *redis.Client) {
	/*
		userRepo:        repositories.NewUserRepo(db),
		postsRepo:       repositories.NewPostsRepo(db),
		userSessionRepo: repositories.NewUserSessionRepo(redisConn),
		permissionsRepo: repositories.NewPermissionsRepo(db),
		generator:       utils.NewGenerator(),
	*/

	userRepo := repositories.NewUserRepo(db)
	postsRepo := repositories.NewPostsRepo(db)
	userSessionRepo := repositories.NewUserSessionRepo(redisConn)
	permissionsRepo := repositories.NewPermissionsRepo(db)
	generator := utils.NewGenerator()
	userService := services.NewUserService(
		userRepo,
		postsRepo,
		userSessionRepo,
		permissionsRepo,
		generator,
	)

	controller := controllers.NewUserController(userService)

	router.GET("/profile", controller.GetUserAndPosts)

	router.POST("/login", controller.LoginUser)

	router.POST("/new", controller.RegisterUser)

}
