package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/services"
)

type IPostController interface {
	GetApprovedPosts(ctx *gin.Context)
}

func NewPostController(db *sqlx.DB) IPostController {
	return &postController{
		postService: services.NewPostService(db),
	}
}

type postController struct {
	postService interface {
		GetApprovedPosts() ([]models.Post, models.IError)
	}
}

func (c *postController) GetApprovedPosts(ctx *gin.Context) {
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
