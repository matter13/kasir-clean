package repository

import (
	"context"
	"kasir-clean/domain/karyawan/models"
	"gorm.io/gorm"
)

type KaryawanRepository struct {
	Conn *gorm.DB	
}

func NewKaryawanRepository(Conn *gorm.DB) models.KaryawanRepository {
	return &KaryawanRepository{Conn}
}

//Get
func (m *KaryawanRepository)Get(ctx context.Context)(res []models.Karyawan , err error) {
	var karyawan []models.Karyawan	
	m.Conn.Find(&karyawan)
	return karyawan,nil
}

//Tambah
func (m *KaryawanRepository) Tambah(k models.Karyawan)(karyawan models.Karyawan,err error) {
	m.Conn.Create(&k)
	return k,nil
}

//View
func(m *KaryawanRepository)View(id_karyawan string)(res models.Karyawan , err error) {
	var karyawan models.Karyawan
	m.Conn.First(&karyawan , id_karyawan)
	return karyawan , nil
}

//Delete
func(m *KaryawanRepository)Delete(id_karyawan string)(res models.Karyawan, err error) {
	var karyawan models.Karyawan
	m.Conn.Delete(&karyawan , id_karyawan)
	return res , nil
}

//Edit
func (m *KaryawanRepository) Edit(k models.Karyawan , id_karyawan string)(res models.Karyawan,err error) {
	var karyawan models.Karyawan
	m.Conn.First(&karyawan , id_karyawan)
	karyawan = k
	m.Conn.Save(&karyawan)
	return karyawan , nil
}