package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
)

type IPermissionsRepo interface {
	GetPermissionsList(idUser int64) ([]models.Permission, error)
}

func NewPermissionsRepo(db *sqlx.DB) IPermissionsRepo {
	return &permissionsRepo{
		db: db,
	}
}

type permissionsRepo struct {
	db *sqlx.DB
}

func (c *permissionsRepo) GetPermissionsList(idUser int64) ([]models.Permission, error) {
	var permissions []models.Permission

	err := c.db.Select(
		&permissions,
		`
		SELECT 
			p.id as permission_id, p.title as permission_title, p.code as permission_code
		FROM public.users u
		INNER JOIN public.roles AS r ON r.id = u.id_role
		INNER JOIN public.roles_permissions AS rp ON rp.id_role = r.id
		INNER JOIN public.permissions AS p ON p.id = rp.id_permission
		WHERE u.id = $1
		`,
		idUser)

	return permissions, err
}
