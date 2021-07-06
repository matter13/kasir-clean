package repository

import (
	"kasir-clean/domain/auth/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Conn *gorm.DB
}

func NewAuthRepository(Conn *gorm.DB) models.AuthRepository {
	return &AuthRepository{Conn}
}

// Cek ...
func (m *AuthRepository) Cek(username, password string) (res models.Auth, err error) {
	var auth models.Auth
	m.Conn.First(&auth, "username", username)
	return auth, nil
}
