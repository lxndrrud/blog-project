package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/services"
)

type IPostController interface {
	GetApprovedPosts(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
}

func NewPostController(db *sqlx.DB, redisConn *redis.Client) IPostController {
	return &postController{
		postService: services.NewPostService(db, redisConn),
	}
}

type postController struct {
	postService interface {
		GetApprovedPosts() ([]models.Post, models.IError)
		CreatePost(title, text, annotation, token string,
			timeStart *time.Time, timeEnd *time.Time) models.IError
	}
}

func (c postController) GetApprovedPosts(ctx *gin.Context) {
	posts, err := c.postService.GetApprovedPosts()
	if err != nil {
		ctx.JSON(err.GetCode(), map[string]any{
			"message": err.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"posts": posts,
	})
}

func (c postController) CreatePost(ctx *gin.Context) {
	if ctx.PostForm("title") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан заголовок поста",
		})
		return
	}
	if ctx.PostForm("text") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан заголовок поста",
		})
		return
	}
	if ctx.PostForm("annotation") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан аннотация поста",
		})
		return
	}
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Создание постов разрешено только авторизованным пользователям!",
		})
		return
	}

	var timeStart *time.Time = nil
	if ctx.PostForm("timeStart") != "" {
		parsed, err := time.Parse("2006-01-02 15:04:05", ctx.PostForm("timeStart"))
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, map[string]any{
				"message": "Неправильный формат времени начала срока доступности поста!",
			})
			return
		}
		if err == nil {
			timeStart = &parsed
		}
	}
	var timeEnd *time.Time = nil
	if ctx.PostForm("timeEnd") != "" {
		parsed, err := time.Parse("2006-01-02 15:04:05", ctx.PostForm("timeEnd"))
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusBadRequest, map[string]any{
				"message": "Неправильный формат времени окончания срока доступности поста!",
			})
			return
		}
		if err == nil {
			timeEnd = &parsed
		}
	}

	errService := c.postService.CreatePost(ctx.PostForm("title"), ctx.PostForm("text"),
		ctx.PostForm("annotation"),
		ctx.GetHeader("auth-token"), timeStart, timeEnd)
	if errService != nil {
		ctx.JSON(errService.GetCode(), map[string]any{
			"message": errService.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
