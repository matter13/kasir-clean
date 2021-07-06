<<<<<<< HEAD
package repository

import (
	"context"
	"kasir-clean/domain/barang/models"

	"gorm.io/gorm"
)

type BarangRepository struct {
	Conn *gorm.DB
}

func NewBarangRepository(Conn *gorm.DB) models.BarangRepository {
	return &BarangRepository{Conn}
}

// Get ...
func (m *BarangRepository) Get(ctx context.Context) (res []models.Barang, err error) {
	var barang []models.Barang
	m.Conn.Find(&barang)
	return barang, nil
}

// Create ...
func (m *BarangRepository) Create(k models.Barang) (barang models.Barang, err error) {
	m.Conn.Create(&k)
	return k, nil
}

// Show ...
func (m *BarangRepository) Show(id_barang string) (res models.Barang, err error) {
	var barang models.Barang
	m.Conn.First(&barang, id_barang)
	return barang, nil
}

// Delete ...
func (m *BarangRepository) Delete(id_barang string) (res models.Barang, err error) {
	var barang models.Barang
	m.Conn.Delete(&barang, id_barang)
	return res, nil
}

// Update ...
func (m *BarangRepository) Update(k models.Barang, id_barang string) (res models.Barang, err error) {
	var barang models.Barang
	m.Conn.First(&barang, id_barang)
	barang = k
	m.Conn.Save(&barang)
	return barang, nil
}
=======
package repository

import (
	"context"
	"kasir-clean/domain/barang/models"

	"gorm.io/gorm"
)

type BarangRepository struct {
	Conn *gorm.DB
}

func NewBarangRepository(Conn *gorm.DB) models.BarangRepository {
	return &BarangRepository{Conn}
}

// Get ...
func (m *BarangRepository) Get(ctx context.Context) (res []models.Barang, err error) {
	var barang []models.Barang
	m.Conn.Find(&barang)
	return barang, nil
}

// Create ...
func (m *BarangRepository) Create(k models.Barang) (barang models.Barang, err error) {
	m.Conn.Create(&k)
	return k, nil
}

// Show ...
func (m *BarangRepository) Show(id_barang string) (res models.Barang, err error) {
	var barang models.Barang
	m.Conn.First(&barang, id_barang)
	return barang, nil
}

// Delete ...
func (m *BarangRepository) Delete(id_barang string) (res models.Barang, err error) {
	var barang models.Barang
	m.Conn.Delete(&barang, id_barang)
	return res, nil
}

// Update ...
func (m *BarangRepository) Update(k models.Barang, id_barang string) (res models.Barang, err error) {
	var barang models.Barang
	m.Conn.First(&barang, id_barang)
	barang = k
	m.Conn.Save(&barang)
	return barang, nil
}
>>>>>>> master
