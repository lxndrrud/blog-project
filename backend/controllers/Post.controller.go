package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/services"
)

type IPostController interface {
	GetApprovedPosts(ctx *gin.Context)
	GetApprovedPost(ctx *gin.Context)
	GetPostsNeedToApprove(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
	ApprovePost(ctx *gin.Context)
	RejectPost(ctx *gin.Context)
}

func NewPostController(db *sqlx.DB, redisConn *redis.Client) IPostController {
	return &postController{
		postService: services.NewPostService(db, redisConn),
	}
}

type postController struct {
	postService interface {
		GetApprovedPosts() ([]models.Post, models.IError)
		GetApprovedPost(idPost int64) (models.Post, models.IError)
		GetPostsNeedToApprove(token string) ([]models.Post, models.IError)
		CreatePost(title, text, annotation, token string,
			timeStart *time.Time, timeEnd *time.Time) models.IError
		ApprovePost(token string, idPost int64) models.IError
		RejectPost(token string, idPost int64) models.IError
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
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Создание постов разрешено только авторизованным пользователям!",
		})
		return
	}
	var form struct {
		Title      string `json:"title"`
		Text       string `json:"text"`
		Annotation string `json:"annotation"`
		TimeStart  string `json:"timeStart"`
		TimeEnd    string `json:"timeEnd"`
	}
	err := json.NewDecoder(ctx.Request.Body).Decode(&form)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]any{
			"message": err.Error(),
		})
		return
	}
	ctx.Request.Body.Close()
	if form.Title == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан заголовок поста",
		})
		return
	}
	if form.Text == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан заголовок поста",
		})
		return
	}
	if form.Annotation == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан аннотация поста",
		})
		return
	}

	var timeStart *time.Time = nil
	if form.TimeStart != "" {
		parsed, err := time.Parse("2006-01-02 15:04:05", form.TimeStart)
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
	if form.TimeEnd != "" {
		parsed, err := time.Parse("2006-01-02 15:04:05", form.TimeEnd)
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

	errService := c.postService.CreatePost(form.Title, form.Text,
		form.Annotation,
		ctx.GetHeader("auth-token"), timeStart, timeEnd)
	if errService != nil {
		ctx.JSON(errService.GetCode(), map[string]any{
			"message": errService.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}

func (c postController) GetApprovedPost(ctx *gin.Context) {
	if ctx.Query("idPost") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не указан!",
		})
		return
	}
	idPost, err := strconv.ParseInt(ctx.Query("idPost"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не распознан!",
		})
		return
	}
	post, errService := c.postService.GetApprovedPost(idPost)
	if errService != nil {
		ctx.JSON(errService.GetCode(), map[string]any{
			"message": errService.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"post": post,
	})
}

func (c postController) GetPostsNeedToApprove(ctx *gin.Context) {
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Неавторизованным пользователям доступ запрещен!",
		})
		return
	}
	posts, err := c.postService.GetPostsNeedToApprove(ctx.GetHeader("auth-token"))
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

func (c postController) ApprovePost(ctx *gin.Context) {
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Неавторизованным пользователям доступ запрещен!",
		})
		return
	}
	if ctx.PostForm("idPost") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не указан!",
		})
		return
	}
	idPost, err := strconv.ParseInt(ctx.PostForm("idPost"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не распознан!",
		})
		return
	}
	errService := c.postService.ApprovePost(ctx.GetHeader("auth-token"), idPost)
	if errService != nil {
		ctx.JSON(errService.GetCode(), map[string]any{
			"message": errService.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (c postController) RejectPost(ctx *gin.Context) {
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Неавторизованным пользователям доступ запрещен!",
		})
		return
	}
	if ctx.PostForm("idPost") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не указан!",
		})
		return
	}
	idPost, err := strconv.ParseInt(ctx.PostForm("idPost"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не распознан!",
		})
		return
	}
	errService := c.postService.RejectPost(ctx.GetHeader("auth-token"), idPost)
	if errService != nil {
		ctx.JSON(errService.GetCode(), map[string]any{
			"message": errService.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
