package services

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/repositories"
)

type IPostService interface {
	GetApprovedPosts() ([]models.Post, models.IError)
}

func NewPostService(db *sqlx.DB) IPostService {
	return &postService{
		postRepo: repositories.NewPostsRepo(db),
	}
}

type postService struct {
	postRepo interface {
		GetApprovedPosts() ([]models.Post, error)
	}
}

func (c *postService) GetApprovedPosts() ([]models.Post, models.IError) {
	posts, err := c.postRepo.GetApprovedPosts()
	if err != nil {
		return posts, models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	return posts, nil
}
