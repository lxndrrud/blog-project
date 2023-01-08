package mocks

import "github.com/lxndrrud/blog-project/models"

type SuccessPermissionRepoMock struct{}

func (c *SuccessPermissionRepoMock) GetPermissionsList(idUser int64) ([]models.Permission, error) {
	return []models.Permission{
		{
			Id:    1,
			Title: "test permission",
			Code:  "СОЗД_ПОСТЫ",
		},
		{
			Id:    2,
			Title: "test permission",
			Code:  "МОДЕР_ПОСТЫ",
		},
	}, nil
}
