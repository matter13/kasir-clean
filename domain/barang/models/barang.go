package models

import (
	"context"
)

type Tabler interface {
	TableName() string
}

func (Barang) TableName() string {
	return "barang"
}

type Barang struct {
	Id_barang   string `json:"id_barang"`
	Nama_barang string `json:"nama_barang"`
	Deskripsi   string `json:"deskripsi"`
	Stok        string `json:"stok"`
	Harga       string `json:"harga"`
	Tgl_masuk   string `json:"tgl_masuk"`
}
type BarangEntity interface {
	Get(ctx context.Context) ([]Barang, error)
	Create(barang Barang) (Barang, error)
	Show(id_barang string) (Barang, error)
	Delete(id_barang string) (Barang, error)
	Update(barang Barang, id_barang string) (Barang, error)
}

// BarangRepository ...
type BarangRepository interface {
	Get(ctx context.Context) (res []Barang, err error)
	Create(barang Barang) (Barang, error)
	Show(id_barang string) (Barang, error)
	Delete(id_barang string) (Barang, error)
	Update(barang Barang, id_barang string) (Barang, error)
}
