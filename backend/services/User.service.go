package services

import (
	"database/sql"
	"net/http"

	"github.com/lxndrrud/blog-project/models"
)

type IUserServiceUserRepo interface {
	GetUserById(idUser int64) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
	InsertUser(login, password string) (int64, error)
}
type IUserServicePostsRepo interface {
	GetUserPosts(idUser int64) ([]models.Post, error)
}
type IUserServiceUserSessionRepo interface {
	CreateSession(idUser int64, password string) (models.UserSession, error)
	GetUserSession(token string) (models.UserSession, error)
}
type IUserServicePermissionsRepo interface {
	GetPermissionsList(idUser int64) ([]models.Permission, error)
}
type IUserServiceGenerator interface {
	CheckPassword(hashedPassword, password string) error
	HashPassword(password string) (string, error)
}

type IUserService interface {
	GetUserAndPosts(idUser int64) (models.User, []models.Post, models.IError)
	LoginUser(login, password string) (string, models.User, []models.Permission, models.IError)
	RegisterUser(login, password string) models.IError
}

func NewUserService(
	userRepo IUserServiceUserRepo,
	postsRepo IUserServicePostsRepo,
	userSessionRepo IUserServiceUserSessionRepo,
	permissionsRepo IUserServicePermissionsRepo,
	generator IUserServiceGenerator,
) IUserService {
	return &userService{
		userRepo:        userRepo,
		postsRepo:       postsRepo,
		userSessionRepo: userSessionRepo,
		permissionsRepo: permissionsRepo,
		generator:       generator,
	}
}

type userService struct {
	userRepo        IUserServiceUserRepo
	postsRepo       IUserServicePostsRepo
	userSessionRepo IUserServiceUserSessionRepo
	permissionsRepo IUserServicePermissionsRepo
	generator       IUserServiceGenerator
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
			err1Chan <- nil
		}
	}()

	go func() {
		value, err := c.postsRepo.GetUserPosts(idUser)
		postsChan <- value
		if err != nil {
			err2Chan <- models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
		} else {
			err2Chan <- nil
		}
	}()

	user := <-userChan
	posts := <-postsChan
	err1 := <-err1Chan
	err2 := <-err2Chan

	if err1 != nil {
		return models.User{}, []models.Post{}, err1
	} else if err2 != nil {
		return models.User{}, []models.Post{}, err2
	}
	return user, posts, nil
}

func (c userService) LoginUser(login, password string) (string, models.User, []models.Permission, models.IError) {
	user, err := c.userRepo.GetUserByLogin(login)
	if err == sql.ErrNoRows {
		return "", models.User{}, []models.Permission{}, models.NewError(http.StatusNotFound,
			"Пользователь с таким логином не найден!")
	}
	if err != nil {
		return "", models.User{}, []models.Permission{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}

	err = c.generator.CheckPassword(user.Password, password)
	if err != nil {
		return "", models.User{}, []models.Permission{}, models.NewError(http.StatusNotFound,
			"Пользователь с такими данными для входа не найден!")
	}

	session, err := c.userSessionRepo.CreateSession(user.Id, user.Password)
	if err != nil {
		return "", models.User{}, []models.Permission{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}

	permissions, err := c.permissionsRepo.GetPermissionsList(user.Id)
	if err != nil {
		return "", models.User{}, []models.Permission{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return session.Token, user, permissions, nil
}

func (c userService) RegisterUser(login, password string) models.IError {
	// Проверка существующего пользователя
	_, err := c.userRepo.GetUserByLogin(login)
	if err != sql.ErrNoRows {
		return models.NewError(http.StatusConflict,
			"Пользователь с таким логином уже существует!")
	}
	hashedPassword, err := c.generator.HashPassword(password)
	if err != nil {
		return models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка во время хеширования пароля: "+err.Error())
	}
	// Создание нового пользователя
	_, err = c.userRepo.InsertUser(login, hashedPassword)
	if err != nil {
		return models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return nil
}
