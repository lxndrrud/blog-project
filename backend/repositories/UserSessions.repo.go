package repositories

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lxndrrud/blog-project/models"
	"github.com/lxndrrud/blog-project/utils"
)

type IUserSessionRepo interface {
	GetUserSession(token string) (models.UserSession, error)
	CreateSession(idUser int64, password string) (models.UserSession, error)
}

func NewUserSessionRepo(db *redis.Client) IUserSessionRepo {
	return &userSessionRepo{
		db:        db,
		generator: utils.NewGenerator(),
	}
}

type userSessionRepo struct {
	db        *redis.Client
	generator utils.IGenerator
}

func (c userSessionRepo) CreateSession(idUser int64, password string) (models.UserSession, error) {
	// Генерация токена
	token, err := c.generator.GenerateToken(idUser, password)
	if err != nil {
		return models.UserSession{}, err
	}
	session := models.UserSession{
		IdUser: idUser,
		Token:  token,
	}
	// Конвертация в JSON
	hashedSession, err := json.Marshal(&session)
	if err != nil {
		return models.UserSession{}, err
	}
	// Удаление предыдущего сеанса
	_, err = c.db.Del(c.db.Context(), "user_session:"+strconv.FormatInt(idUser, 10)).Result()
	if err != nil {
		return models.UserSession{}, err
	}
	// Создание нового сеанса
	_, err = c.db.SetEX(c.db.Context(), "user_session:"+strconv.FormatInt(idUser, 10), hashedSession, time.Hour*2).Result()
	if err != nil {
		return models.UserSession{}, err
	}
	// Создание обратного ключа по токену
	_, err = c.db.SetEX(c.db.Context(), "user_session_rev:"+token, hashedSession, time.Hour*2).Result()
	if err != nil {
		return models.UserSession{}, err
	}
	return session, nil
}

func (c userSessionRepo) GetUserSession(token string) (models.UserSession, error) {
	var session models.UserSession

	result, err := c.db.Get(c.db.Context(), "user_session_rev:"+token).Bytes()
	if err != nil {
		return models.UserSession{}, err
	}

	err = json.Unmarshal(result, &session)
	return session, err
}
