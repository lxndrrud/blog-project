package services

import (
	"database/sql"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/repositories"
	"github.com/lxndrrud/blog-project/utils"
)

type IUserService interface {
	GetUserAndPosts(idUser int64) (models.User, []models.Post, models.IError)
	LoginUser(login, password string) (string, models.IError)
}

func NewUserService(db *sqlx.DB, redisConn *redis.Client) IUserService {
	return &userService{
		userRepo:        repositories.NewUserRepo(db),
		postsRepo:       repositories.NewPostsRepo(db),
		userSessionRepo: repositories.NewUserSessionRepo(redisConn),
		permissionsRepo: repositories.NewPermissionsRepo(db),
		generator:       utils.NewGenerator(),
	}
}

type userService struct {
	userRepo interface {
		GetUserById(idUser int64) (models.User, error)
		GetUserByLogin(login string) (models.User, error)
	}
	postsRepo interface {
		GetUserPosts(idUser int64) ([]models.Post, error)
	}
	userSessionRepo interface {
		CreateSession(idUser int64, password string) (models.UserSession, error)
		GetUserSession(token string) (models.UserSession, error)
	}
	permissionsRepo interface {
		GetPermissionsList(idUser int64) ([]models.Permission, error)
	}
	generator interface {
		CheckPassword(hashedPassword, password string) error
	}
}

func (c userService) GetUserAndPosts(idUser int64) (models.User, []models.Post, models.IError) {

	userChan := make(chan models.User)
	postsChan := make(chan []models.Post)
	err1Chan := make(chan models.IError)
	err2Chan := make(chan models.IError)

	go func() {
		value, err := c.userRepo.GetUserById(idUser)
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

func (c userService) LoginUser(login, password string) (string, models.IError) {
	user, err := c.userRepo.GetUserByLogin(login)
	if err == sql.ErrNoRows {
		return "", models.NewError(http.StatusNotFound, "Пользователь с такими данными для входа не найден!")
	}
	if err != nil {
		return "", models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}

	err = c.generator.CheckPassword(user.Password, password)
	if err != nil {
		return "", models.NewError(http.StatusNotFound, "Пользователь с такими данными для входа не найден!")
	}

	session, err := c.userSessionRepo.CreateSession(user.Id, user.Password)
	if err != nil {
		return "", models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	return session.Token, models.NewError(0, "")
}
