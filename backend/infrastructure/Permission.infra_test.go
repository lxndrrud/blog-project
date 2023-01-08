package infrastructure_test

import (
	"testing"

	"github.com/lxndrrud/blog-project/infrastructure"
	mocks "github.com/lxndrrud/blog-project/mocks/repositories"
	"github.com/lxndrrud/blog-project/models"
	"github.com/magiconair/properties/assert"
)

func TestGetPermissionsListSuccess(t *testing.T) {
	permissionsInfrastructure := infrastructure.NewPermissionInfra(
		&mocks.SuccessPermissionRepoMock{},
		&mocks.SuccessUserSessionRepoMock{},
	)

	value, err := permissionsInfrastructure.GetPermissionsList("asdasd")

	assert.Equal(t, err, nil, "Error while processing!")
	assert.Equal(t, value, []models.Permission{
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
	}, "Wrong permissions found!")
}

func TestGetUserIdByToken(t *testing.T) {
	permissionsInfrastructure := infrastructure.NewPermissionInfra(
		&mocks.SuccessPermissionRepoMock{},
		&mocks.SuccessUserSessionRepoMock{},
	)

	id, err := permissionsInfrastructure.GetUserIdByToken("asdasd")
	assert.Equal(t, err, nil, "Error while processing!")
	assert.Equal(t, id, int64(123), "Wrong user identifier!")
}
