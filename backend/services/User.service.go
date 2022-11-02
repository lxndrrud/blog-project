package services

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/repositories"
)

type IUserService interface {
	GetUserById(idUser int64) (models.User, models.IError)
}

func NewUserService(db *sqlx.DB) IUserService {
	return &userService{
		userRepo: repositories.NewUserRepo(db),
	}
}

type userService struct {
	userRepo interface {
		GetUserById(idUser int64) (models.User, error)
	}
}

func (c userService) GetUserById(idUser int64) (models.User, models.IError) {
	user, err := c.userRepo.GetUserById(idUser)
	if err == sql.ErrNoRows {
		return user, models.NewError(http.StatusNotFound, "Пользователь не найден!")
	} else if err != nil {
		return user, models.NewError(http.StatusInternalServerError, "Непредвиденная ошибка: "+err.Error())
	}
	return user, nil
}
