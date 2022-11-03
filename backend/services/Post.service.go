package services

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/repositories"
	"github.com/lxndrrud/blog-project/utils"
)

type IPostService interface {
	GetApprovedPosts() ([]models.Post, models.IError)
	CreatePost(title, text, token string, timeEnd *time.Time) models.IError
}

func NewPostService(db *sqlx.DB, redisConn *redis.Client) IPostService {
	return &postService{
		postRepo:          repositories.NewPostsRepo(db),
		userSessionRepo:   repositories.NewUserSessionRepo(redisConn),
		permissionRepo:    repositories.NewPermissionsRepo(db),
		permissionChecker: utils.NewPermissionChecker(),
	}
}

type postService struct {
	postRepo interface {
		GetApprovedPosts() ([]models.Post, error)
		InsertPost(post models.Post) error
	}
	userSessionRepo interface {
		GetUserSession(token string) (models.UserSession, error)
	}
	permissionRepo interface {
		GetPermissionsList(idUser int64) ([]models.Permission, error)
	}
	permissionChecker utils.IPermissionChecker
}

func (c postService) GetApprovedPosts() ([]models.Post, models.IError) {
	posts, err := c.postRepo.GetApprovedPosts()
	if err != nil {
		return posts, models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	return posts, nil
}

func (c postService) CreatePost(title, text, token string, timeEnd *time.Time) models.IError {
	// Получить сессию по токену
	userSession, err := c.userSessionRepo.GetUserSession(token)
	if err == redis.Nil {
		return models.NewError(http.StatusNotFound, "Сеанс не найден. Пожалуйста, войдите по новой.")
	}
	if err != nil {
		return models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	// Получить разрешения пользователя по информации с токена
	permissions, err := c.permissionRepo.GetPermissionsList(userSession.IdUser)
	if err != nil {
		return models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	// Проверить разрешение на создание постов
	if !c.permissionChecker.CanCreatePosts(permissions) {
		return models.NewError(http.StatusForbidden, "Вам запрещено создавать посты!")
	}

	dbTimeEnd := sql.NullTime{}
	if timeEnd != nil {
		dbTimeEnd = sql.NullTime{
			Valid: true,
			Time:  *timeEnd,
		}
	}
	post := models.Post{
		Title:    title,
		Text:     text,
		TimeEnd:  dbTimeEnd,
		IdAuthor: userSession.IdUser,
	}
	err = c.postRepo.InsertPost(post)
	if err != nil {
		return models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	return nil
}
