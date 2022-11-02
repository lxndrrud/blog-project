package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/services"
)

type IUserController interface {
	GetUserById(ctx *gin.Context)
}

func NewUserController(db *sqlx.DB) IUserController {
	return &userController{
		userService: services.NewUserService(db),
	}
}

type userController struct {
	userService interface {
		GetUserById(idUser int64) (models.User, models.IError)
	}
}

func (c userController) GetUserById(ctx *gin.Context) {
	if ctx.Query("idUser") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан идентификатор пользователя!",
		})
		return
	}
	idUser, err := strconv.ParseInt(ctx.Query("idUser"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Неверно указан идентификатор!",
		})
		return
	}
	user, errService := c.userService.GetUserById(idUser)
	if errService != nil {
		ctx.JSON(errService.GetCode(), errService)
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"user": user,
	})

}
