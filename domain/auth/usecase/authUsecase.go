package usecase

import (
	"kasir-clean/domain/auth/models"
)

type Response struct {
	Pesan string
}
type AuthEntity struct {
	authRepo models.AuthRepository
}

func NewAuthEntity(a models.AuthRepository) models.AuthEntity {
	return &AuthEntity{
		authRepo: a,
	}
}

func (be *AuthEntity) Cek(username string, password string) (Response, error) {

	auth, err := be.authRepo.Cek(username, password)
	if err != nil {
		return Response{Pesan: "Maaf"}, err
	} else if username != auth.Username {
		return Response{Pesan: "Maap"}, err
	} else {
		return Response{Pesan: "Username & Pw benar"}, err
	}
	//return auth, err
}
