package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/controllers"
)

func setupUserRouter(router *gin.RouterGroup, db *sqlx.DB, redisConn *redis.Client) {
	controller := controllers.NewUserController(db, redisConn)

	router.GET("/profile", controller.GetUserAndPosts)

	router.POST("/login", controller.LoginUser)

	router.POST("/new", controller.RegisterUser)

}
