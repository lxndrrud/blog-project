package mocks

import (
	"github.com/lxndrrud/blog-project/models"
)

type SuccessUserSessionRepoMock struct {
}

func (c SuccessUserSessionRepoMock) CreateSession(idUser int64, password string) (models.UserSession, error) {
	return models.UserSession{
		IdUser: idUser,
		Token:  "asdasdTRoken",
	}, nil
}

func (c SuccessUserSessionRepoMock) GetUserSession(token string) (models.UserSession, error) {
	return models.UserSession{
		IdUser: 123,
		Token:  "asdasdTRoken",
	}, nil
}
