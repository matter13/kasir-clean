<<<<<<< HEAD
package models

import (
	"context"
)

type Tabler interface {
	TableName() string
}

func (Karyawan) TableName() string {
	return "karyawan"
}

type Karyawan struct {
	Id_karyawan string `json:"id_karyawan"`
	Nama string `json:"nama"`
	Alamat string `json"alamat"`
	No_telp string `json:"no_telp"`

}

type KaryawanEntity interface {
	Get(ctx context.Context)([]Karyawan, error)
	Tambah(karyawan Karyawan)(Karyawan, error)
	View(id_karyawan string)(Karyawan, error)
	Delete(id_karyawan string)(Karyawan,error)
	Edit(karyawan Karyawan , id_karyawan string)(Karyawan , error)
}

type KaryawanRepository interface {
	Get(ctx context.Context)(res []Karyawan, err error)
	Tambah(karyawan Karyawan)(Karyawan , error)
	View(id_karyawan string)(Karyawan , error)
	Delete(id_karyawan string)(Karyawan, error)
	Edit(karyawan Karyawan , id_karyawan string)(Karyawan , error)
}
=======
package models

import (
	"context"
)

type Tabler interface {
	TableName() string
}

func (Karyawan) TableName() string {
	return "karyawan"
}

type Karyawan struct {
	Id_karyawan string `json:"id_karyawan"`
	Nama        string `json:"nama"`
	Alamat      string `json:"alamat"`
	No_telp     string `json:"no_telp"`
}

type KaryawanEntity interface {
	Get(ctx context.Context) ([]Karyawan, error)
	Tambah(karyawan Karyawan) (Karyawan, error)
	View(id_karyawan string) (Karyawan, error)
	Delete(id_karyawan string) (Karyawan, error)
	Edit(karyawan Karyawan, id_karyawan string) (Karyawan, error)
}

type KaryawanRepository interface {
	Get(ctx context.Context) (res []Karyawan, err error)
	Tambah(karyawan Karyawan) (Karyawan, error)
	View(id_karyawan string) (Karyawan, error)
	Delete(id_karyawan string) (Karyawan, error)
	Edit(karyawan Karyawan, id_karyawan string) (Karyawan, error)
}
>>>>>>> master
