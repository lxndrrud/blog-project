package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func BootstrapRouter(db *sqlx.DB) *gin.Engine {
	app := gin.Default()

	// Префиксный роутер
	prefixRouter := app.Group("/backend")
	prefixRouter.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"message": "Blog Index message...",
		})
	})

	// Пользователи
	userRouter := prefixRouter.Group("/users")
	setupUserRouter(userRouter, db)

	// Посты
	postsRouter := prefixRouter.Group("/posts")
	setupPostRouter(postsRouter, db)

	return app
}
