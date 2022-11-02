package services

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/repositories"
)

type IUserService interface {
	GetUserAndPosts(idUser int64) (models.User, []models.Post, models.IError)
}

func NewUserService(db *sqlx.DB) IUserService {
	return &userService{
		userRepo:  repositories.NewUserRepo(db),
		postsRepo: repositories.NewPostsRepo(db),
	}
}

type userService struct {
	userRepo interface {
		GetUserById(idUser int64) (models.User, error)
	}
	postsRepo interface {
		GetUserPosts(idUser int64) ([]models.Post, error)
	}
}

func (c userService) GetUserAndPosts(idUser int64) (models.User, []models.Post, models.IError) {

	userChan := make(chan models.User)
	postsChan := make(chan []models.Post)
	err1Chan := make(chan models.IError)
	err2Chan := make(chan models.IError)

	go func() {
		value, err := c.userRepo.GetUserById(idUser)
		fmt.Println(value, err)
		userChan <- value
		if err == sql.ErrNoRows {
			err1Chan <- models.NewError(http.StatusNotFound, "Пользователь не найден!")
		} else if err != nil {
			err1Chan <- models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
		} else {
			err1Chan <- models.NewError(0, "")
		}
	}()

	go func() {
		value, err := c.postsRepo.GetUserPosts(idUser)
		fmt.Println(value, err)
		postsChan <- value
		if err != nil {
			err2Chan <- models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
		} else {
			err2Chan <- models.NewError(0, "")
		}
	}()

	user := <-userChan
	posts := <-postsChan
	err1 := <-err1Chan
	err2 := <-err2Chan

	if err1.GetCode() != 0 {
		return models.User{}, []models.Post{}, err1
	} else if err2.GetCode() != 0 {
		return models.User{}, []models.Post{}, err2
	}
	return user, posts, nil
}
