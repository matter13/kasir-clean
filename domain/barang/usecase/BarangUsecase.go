<<<<<<< HEAD
package usecase

import (
	"context"
	"kasir-clean/domain/barang/models"
)

type BarangEntity struct {
	barangRepo models.BarangRepository
}

func NewBarangEntity(a models.BarangRepository) models.BarangEntity {
	return &BarangEntity{
		barangRepo: a,
	}
}

func (be *BarangEntity) Get(c context.Context) (res []models.Barang, err error) {
	res, err = be.barangRepo.Get(c)
	if err != nil {
		return nil, err
	}
	return
}

func (be *BarangEntity) Create(c models.Barang) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Create(c)
	if err != nil {
		return models.Barang{}, err
	}
	return
}

func (be *BarangEntity) Show(id_barang string) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Show(id_barang)
	if err != nil {
		return models.Barang{}, err
	}
	return
}

func (be *BarangEntity) Delete(id_barang string) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Delete(id_barang)
	if err != nil {
		return models.Barang{}, err
	}
	return
}

func (be *BarangEntity) Update(c models.Barang, id_barang string) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Update(c, id_barang)
	if err != nil {
		return models.Barang{}, err
	}
	return
}
=======
package usecase

import (
	"context"
	"kasir-clean/domain/barang/models"
)

type BarangEntity struct {
	barangRepo models.BarangRepository
}

func NewBarangEntity(a models.BarangRepository) models.BarangEntity {
	return &BarangEntity{
		barangRepo: a,
	}
}

func (be *BarangEntity) Get(c context.Context) (res []models.Barang, err error) {
	res, err = be.barangRepo.Get(c)
	if err != nil {
		return nil, err
	}
	return
}

func (be *BarangEntity) Create(c models.Barang) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Create(c)
	if err != nil {
		return models.Barang{}, err
	}
	return
}

func (be *BarangEntity) Show(id_barang string) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Show(id_barang)
	if err != nil {
		return models.Barang{}, err
	}
	return
}

func (be *BarangEntity) Delete(id_barang string) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Delete(id_barang)
	if err != nil {
		return models.Barang{}, err
	}
	return
}

func (be *BarangEntity) Update(c models.Barang, id_barang string) (barang models.Barang, err error) {
	barang, err = be.barangRepo.Update(c, id_barang)
	if err != nil {
		return models.Barang{}, err
	}
	return
}
>>>>>>> master
