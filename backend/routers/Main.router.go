package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func BootstrapRouter(db *sqlx.DB) *gin.Engine {
	app := gin.Default()
	prefixRouter := app.Group("/backend")
	userRouter := prefixRouter.Group("/users")
	setupUserRouter(userRouter, db)

	return app
}
