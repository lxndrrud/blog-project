package infrastructure

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/lxndrrud/blog-project/models"
)

type IPermissionInfraPermissionRepo interface {
	GetPermissionsList(idUser int64) ([]models.Permission, error)
}

type IPermissionInfraUserSessionRepo interface {
	GetUserSession(token string) (models.UserSession, error)
}

type IPermissionInfra interface {
	GetPermissionsList(token string) ([]models.Permission, models.IError)
	GetUserIdByToken(token string) (int64, models.IError)
}

func NewPermissionInfra(
	permissionRepo IPermissionInfraPermissionRepo,
	userSessionRepo IPermissionInfraUserSessionRepo,
) IPermissionInfra {
	return &permissionInfra{
		permissionRepo:  permissionRepo,
		userSessionRepo: userSessionRepo,
	}
}

type permissionInfra struct {
	permissionRepo interface {
		GetPermissionsList(idUser int64) ([]models.Permission, error)
	}
	userSessionRepo interface {
		GetUserSession(token string) (models.UserSession, error)
	}
}

func (c *permissionInfra) GetPermissionsList(token string) ([]models.Permission, models.IError) {
	userSession, err := c.userSessionRepo.GetUserSession(token)
	if err == redis.Nil {
		return []models.Permission{}, models.NewError(http.StatusNotFound,
			"Сеанс не найден. Пожалуйста, войдите по новой.")
	}
	if err != nil {
		return []models.Permission{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	// Получить разрешения пользователя по информации с токена
	permissions, err := c.permissionRepo.GetPermissionsList(userSession.IdUser)
	if err != nil {
		return []models.Permission{}, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return permissions, nil
}

func (c *permissionInfra) GetUserIdByToken(token string) (int64, models.IError) {
	userSession, err := c.userSessionRepo.GetUserSession(token)
	if err == redis.Nil {
		return 0, models.NewError(http.StatusNotFound,
			"Сеанс не найден. Пожалуйста, войдите по новой.")
	}
	if err != nil {
		return 0, models.NewError(http.StatusInternalServerError,
			"Непредвиденная ошибка: "+err.Error())
	}
	return userSession.IdUser, nil
}
