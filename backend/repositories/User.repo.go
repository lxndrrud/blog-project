package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
)

type IUserRepo interface {
	GetUserById(idUser int64) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
}

func NewUserRepo(db *sqlx.DB) IUserRepo {
	return &userRepo{
		db: db,
	}
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

func (c *userRepo) GetUserByLogin(login string) (models.User, error) {
	var user models.User

	err := c.db.Get(
		&user,
		`
			SELECT 
				u.id as user_id, u.login as user_login, u.password as user_password
			FROM public.users u
			WHERE login = $1
		`,
		login,
	)

	return user, err
}
