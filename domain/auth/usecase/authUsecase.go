<<<<<<< HEAD
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

func (be *AuthEntity) Cek(username string) (auth models.Auth, err error) {
	auth , err = be.authRepo.Cek(username,password)
	if err != nil {
		if auth.Username == 
	}
return
}
=======
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

func (be *AuthEntity) Cek(username string) (auth models.Auth, err error) {
	auth , err = be.authRepo.Cek(username,password)
	if err != nil {
		if auth.Username == 
	}
return
}
>>>>>>> master
