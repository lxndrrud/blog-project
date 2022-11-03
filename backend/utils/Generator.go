package utils

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type IGenerator interface {
	GenerateToken(idUser int64, password string) (string, error)
	HashPassword(password string) (string, error)
	CheckPassword(hashedPassword, password string) error
}

func NewGenerator() IGenerator {
	return &generator{}
}

type generator struct {
}

func (c generator) GenerateToken(idUser int64, password string) (string, error) {
	bytes_ := []byte(strconv.FormatInt(idUser, 10) + password)
	hash, err := bcrypt.GenerateFromPassword(bytes_, bcrypt.DefaultCost)
	return string(hash), err
}

func (c generator) HashPassword(password string) (string, error) {
	bytes_ := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytes_, bcrypt.DefaultCost)
	return string(hash), err
}

func (c generator) CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
