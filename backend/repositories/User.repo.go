package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/lxndrrud/blog-project/models"
)

type IUserRepo interface {
	GetUserById(idUser int64) (models.User, error)
	GetUserByLogin(login string) (models.User, error)
	InsertUser(login, password string) (int64, error)
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

func (c *userRepo) InsertUser(login, password string) (int64, error) {
	var idUser int64
	err := c.db.Get(&idUser,
		`
			INSERT INTO public.users (login, password, id_role) VALUES
			($1, $2, 2)
			RETURNING id;
		`,
		login, password)
	return idUser, err
}
