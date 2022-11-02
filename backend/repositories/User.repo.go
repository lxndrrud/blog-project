package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
)

type IUserRepo interface {
	GetUserById(idUser int64) (models.User, error)
}

func NewUserRepo(db *sqlx.DB) IUserRepo {
	newRepo := &userRepo{
		db: db,
	}
	return newRepo
}

type userRepo struct {
	db *sqlx.DB
}

func (c *userRepo) GetUserById(idUser int64) (models.User, error) {
	var user models.User

	err := c.db.Get(
		&user,
		`
			SELECT 
				u.id as user_id, u.login as user_login, u.password as user_password
			FROM "public"."users" u
			WHERE u.id = $1
		`,
		idUser)
	return user, err
}
