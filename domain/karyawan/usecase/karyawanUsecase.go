<<<<<<< HEAD
package usecase

import (
	"context"
	"kasir-clean/domain/karyawan/models"
)

type KaryawanEntity struct {
	karyawanRepo models.KaryawanRepository
}

func NewKaryawanEntity(a models.KaryawanRepository) models.KaryawanEntity {
	return &KaryawanEntity{
		karyawanRepo: a,
	}

}

func (ke *KaryawanEntity) Get(c context.Context) (res []models.Karyawan, err error) {
	res, err = ke.karyawanRepo.Get(c)
	if err != nil {
		return nil, err
	}
	return
}

func (ke *KaryawanEntity) Tambah(c models.Karyawan) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.Tambah(c)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}

func (ke *KaryawanEntity) View(id_karyawan string) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.View(id_karyawan)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}

func (ke *KaryawanEntity) Delete(id_karyawan string) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.Delete(id_karyawan)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}

func (ke *KaryawanEntity) Edit(c models.Karyawan, id_karyawan string) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.Edit(c, id_karyawan)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}
=======
package usecase

import (
	"context"
	//barang "kasir-clean/domain/barang/models"
	models "kasir-clean/domain/karyawan/models"
)

/*type Karyawan struct {
	Barang models.BarangUsecase
}*/

type KaryawanEntity struct {
	karyawanRepo models.KaryawanRepository
	//BarangUsecase barang.BarangEntity
}

func NewKaryawanEntity(a models.KaryawanRepository) models.KaryawanEntity {
	return &KaryawanEntity{
		karyawanRepo: a,
	}

}

func (ke *KaryawanEntity) Get(c context.Context) (res []models.Karyawan, err error) {
	res, err = ke.karyawanRepo.Get(c)
	if err != nil {
		return nil, err
	}
	return
}

func (ke *KaryawanEntity) Tambah(c models.Karyawan) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.Tambah(c)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}

func (ke *KaryawanEntity) View(id_karyawan string) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.View(id_karyawan)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}

func (ke *KaryawanEntity) Delete(id_karyawan string) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.Delete(id_karyawan)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}

func (ke *KaryawanEntity) Edit(c models.Karyawan, id_karyawan string) (karyawan models.Karyawan, err error) {
	karyawan, err = ke.karyawanRepo.Edit(c, id_karyawan)
	if err != nil {
		return models.Karyawan{}, err
	}
	return
}
>>>>>>> master
