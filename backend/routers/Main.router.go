package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func BootstrapRouter(db *sqlx.DB) *gin.Engine {
	app := gin.Default()

	return app
}
