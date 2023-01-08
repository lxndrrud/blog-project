package mocks

import "github.com/lxndrrud/blog-project/models"

type SuccessUserRepoMock struct {
}

func (c *SuccessUserRepoMock) GetUserById(idUser int64) (models.User, error) {
	return models.User{
		Id:       1,
		Login:    "TestUser",
		Password: "TestPassword",
	}, nil
}

func (c *SuccessUserRepoMock) GetUserByLogin(login string) (models.User, error) {
	return models.User{
		Id:       1,
		Login:    login,
		Password: "TestPassword",
	}, nil
}

func (c *SuccessUserRepoMock) InsertUser(login, password string) (int64, error) {
	return 1002, nil
}
