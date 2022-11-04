package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/services"
)

type IUserController interface {
	GetUserAndPosts(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	RegisterUser(ctx *gin.Context)
}

func NewUserController(db *sqlx.DB, redisConn *redis.Client) IUserController {
	return &userController{
		userService: services.NewUserService(db, redisConn),
	}
}

type userController struct {
	userService interface {
		GetUserAndPosts(idUser int64) (models.User, []models.Post, models.IError)
		LoginUser(login, password string) (string, models.IError)
		RegisterUser(login, password string) models.IError
	}
}

func (c userController) GetUserAndPosts(ctx *gin.Context) {
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
	user, posts, errService := c.userService.GetUserAndPosts(idUser)
	if errService != nil {
		ctx.JSON(errService.GetCode(), errService)
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"user":  user,
		"posts": posts,
	})
}

func (c userController) LoginUser(ctx *gin.Context) {
	if ctx.PostForm("login") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан логин!",
		})
		return
	}
	if ctx.PostForm("password") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан пароль!",
		})
		return
	}
	token, err := c.userService.LoginUser(ctx.PostForm("login"), ctx.PostForm("password"))
	if err.GetCode() != 0 {
		ctx.JSON(err.GetCode(), map[string]any{
			"message": err.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (c userController) RegisterUser(ctx *gin.Context) {
	if ctx.PostForm("login") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан логин!",
		})
		return
	}
	if ctx.PostForm("password") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан пароль!",
		})
		return
	}
	err := c.userService.RegisterUser(ctx.PostForm("login"), ctx.PostForm("password"))
	if err.GetCode() != 0 {
		ctx.JSON(err.GetCode(), map[string]any{
			"message": err.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
