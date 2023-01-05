package controllers

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lxndrrud/blog-project/models"
)

type IPostControllerPostService interface {
	GetUserPosts(token string) ([]models.Post, models.IError)
	GetApprovedPosts() ([]models.Post, models.IError)
	GetApprovedPost(idPost int64) (models.Post, models.IError)
	GetPostsNeedToApprove(token string) ([]models.Post, models.IError)
	GetPostNeedToApprove(token string, idPost int64) (models.Post, models.IError)
	CreatePost(title, text, annotation, token string,
		timeStart *time.Time, timeEnd *time.Time, pictureFile multipart.File,
		pictureHeader *multipart.FileHeader) models.IError
	ApprovePost(token string, idPost int64) models.IError
	RejectPost(token string, idPost int64) models.IError
	DeletePost(token string, idPost int64) models.IError
}

type IPostController interface {
	GetUserPosts(ctx *gin.Context)
	GetApprovedPosts(ctx *gin.Context)
	GetApprovedPost(ctx *gin.Context)
	GetPostsNeedToApprove(ctx *gin.Context)
	GetPostNeedToApprove(ctx *gin.Context)
	CreatePost(ctx *gin.Context)
	ApprovePost(ctx *gin.Context)
	RejectPost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
}

func NewPostController(postService IPostControllerPostService) IPostController {
	return &postController{
		postService: postService,
	}
}

type postController struct {
	postService IPostControllerPostService
}

func (c postController) GetUserPosts(ctx *gin.Context) {
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Неавторизованным пользователям доступ запрещен!",
		})
		return
	}
	posts, errService := c.postService.GetUserPosts(ctx.GetHeader("auth-token"))
	if errService != nil {
		ctx.JSON(errService.GetCode(), errService.GetMessage())
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"posts": posts,
	})
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

	if ctx.PostForm("title") == "" {
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
	if ctx.PostForm("text") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Не указан текст поста",
		})
		return
	}

	var timeStart *time.Time = nil
	if ctx.PostForm("timeStart") != "null" {
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
	if ctx.PostForm("timeEnd") != "null" {
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
	file, header, err := ctx.Request.FormFile("picture")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": "Картинка для поста не прикреплена!",
		})
		return
	}
	errService := c.postService.CreatePost(ctx.PostForm("title"), ctx.PostForm("text"),
		ctx.PostForm("annotation"),
		ctx.GetHeader("auth-token"), timeStart, timeEnd, file, header)
	if errService != nil {
		fmt.Println(errService.GetMessage())
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

func (c postController) GetPostNeedToApprove(ctx *gin.Context) {
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Неавторизованным пользователям доступ запрещен!",
		})
		return
	}
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
	post, errService := c.postService.GetPostNeedToApprove(ctx.GetHeader("auth-token"), idPost)
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

func (c postController) ApprovePost(ctx *gin.Context) {
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Неавторизованным пользователям доступ запрещен!",
		})
		return
	}
	var form struct {
		IdPost int64 `json:"idPost"`
	}
	err := json.NewDecoder(ctx.Request.Body).Decode(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
		return
	}
	ctx.Request.Body.Close()
	errService := c.postService.ApprovePost(ctx.GetHeader("auth-token"), form.IdPost)
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
	var form struct {
		IdPost int64 `json:"idPost"`
	}
	err := json.NewDecoder(ctx.Request.Body).Decode(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
		return
	}
	ctx.Request.Body.Close()
	errService := c.postService.RejectPost(ctx.GetHeader("auth-token"), form.IdPost)
	if errService != nil {
		ctx.JSON(errService.GetCode(), map[string]any{
			"message": errService.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func (c postController) DeletePost(ctx *gin.Context) {
	if ctx.GetHeader("auth-token") == "" {
		ctx.JSON(http.StatusForbidden, map[string]any{
			"message": "Неавторизованным пользователям доступ запрещен!",
		})
		return
	}
	if ctx.Query("idPost") == "" {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не указан",
		})
	}
	idPost, err := strconv.ParseInt(ctx.Query("idPost"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": "Идентификатор поста не распознан",
		})
		return
	}
	errService := c.postService.DeletePost(ctx.GetHeader("auth-token"), idPost)
	if errService != nil {
		ctx.JSON(errService.GetCode(), map[string]any{
			"message": errService.GetMessage(),
		})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
