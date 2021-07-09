package usecase

import (
	"kasir-clean/domain/auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
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
		return models.Response{Pesan: "Unauthorized"}, err
	} else if username != auth.Username {
		return models.Response{Pesan: "Username Tidak Ditemukan"}, err
	} else if username == auth.Username {
		if errr := HashCek(password, auth.Password); !errr {
			return models.Response{Pesan: "Password Salah"}, err
		} else {

			token := jwt.New(jwt.SigningMethodHS256)
			claims := token.Claims.(jwt.MapClaims)
			claims["username"] = auth.Username
			claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

			t, err := token.SignedString([]byte("rahasia"))
			if err != nil {
				panic(err)
			} else {
				return models.Response{Pesan: t}, err
			}
		}
	} else {
		return models.Response{Pesan: "Unauthorized"}, err
	}
	//return auth, err
}
func HashCek(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
