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

func (be *AuthEntity) Cek(username, password string) (models.Response, error) {

	auth, err := be.authRepo.Cek(username, password)
	if err != nil {
		return models.Response{Pesan: "Maaf Username Password salah"}, err
	} else if username != auth.Username {
		return models.Response{Pesan: "Username Tidak Ditemukan"}, err
	} else if username == auth.Username && password == auth.Password {
		return models.Response{Pesan: "Username & Pw benar"}, err
	}else{
		return models.Response{Pesan: "Unauthorized"},err
	}
	//return auth, err
}
