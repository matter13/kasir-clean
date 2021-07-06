package usecase

import (
	"kasir-clean/domain/auth/models"
)

type AuthEntity struct {
	authRepo models.AuthRepository
}

func NewAuthEntity(a models.AuthRepository) models.AuthEntity {
	return &AuthEntity{
		authRepo: a,
	}
}

func (be *AuthEntity) Cek(username, password string) (models.Auth, error) {
	auth, err := be.authRepo.Cek(username, password)
	if err != nil {
		return models.Auth{}, err
	} else if username != auth.Username {
		return models.Auth{Username: "Maaf Username Salah"}, err
	} else {
		return models.Auth{Username: auth.Username, Password: auth.Password}, err
	}
	//return auth, err
}
