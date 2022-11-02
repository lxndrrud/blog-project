package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/controllers"
)

func setupUserRouter(router *gin.RouterGroup, db *sqlx.DB) {
	controller := controllers.NewUserController(db)

	router.GET("/user", controller.GetUserById)

}
