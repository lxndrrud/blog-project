package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/controllers"
)

func setupPostRouter(router *gin.RouterGroup, db *sqlx.DB) {
	controller := controllers.NewPostController(db)
	router.GET("/actual", controller.GetApprovedPosts)
}
