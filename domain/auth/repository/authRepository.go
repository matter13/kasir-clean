<<<<<<< HEAD
package repository

import (
	"kasir-clean/domain/auth/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Conn *gorm.DB
}

// Cek ...
func (m *AuthRepository) Cek(username string) (res models.Auth, err error) {
	var auth models.Auth
	m.Conn.First(&auth, username)
	return auth , nil
}
=======
package repository

import (
	"kasir-clean/domain/auth/models"

	"gorm.io/gorm"
)

type AuthRepository struct {
	Conn *gorm.DB
}

// Cek ...
func (m *AuthRepository) Cek(username string) (res models.Auth, err error) {
	var auth models.Auth
	m.Conn.First(&auth, username)
	return auth , nil
}
>>>>>>> master
