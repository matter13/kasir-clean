package controller

import (
	"kasir-clean/domain/barang/models"
	"net/http"

	"github.com/labstack/echo"
)

type BarangController struct {
	BarangController models.BarangEntity
}

func BarangControllerFunc(c *echo.Group, us models.BarangEntity) {
	handler := &BarangController{
		BarangController: us,
	}

	c.GET("/barang", handler.GetBarang)
	c.POST("/barang", handler.CreateBarang)
	c.GET("/barang/:id_barang", handler.Showbarang)
	c.DELETE("/barang/:id_barang", handler.DeleteBarang)
	c.PUT("/barang/:id_barang", handler.UpdateBarang)

}

//Getkelas(semua barang)
func (a *BarangController) GetBarang(c echo.Context) error {
	barang, _ := a.BarangController.Get(c.Request().Context())
	return c.JSON(http.StatusOK, barang)
}

//CreatedBarang (tambah barang)
func (a *BarangController) CreateBarang(c echo.Context) error {
	data := models.Barang{
		Id_barang:   c.FormValue("id_barang"),
		Nama_barang: c.FormValue("nama_barang"),
		Stok:        c.FormValue("stok"),
		Deskripsi:   c.FormValue("deskripsi"),
		Harga:       c.FormValue("harga"),
		Tgl_masuk:   c.FormValue("tgl_masuk"),
	}
	barang, err := a.BarangController.Create(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Error Meng-input data"})
	}
	return c.JSON(http.StatusCreated, barang)
}

//Tampil barang(parameter)
func (a *BarangController) Showbarang(c echo.Context) error {
	barang, _ := a.BarangController.Show(c.Param("id_barang"))
	return c.JSON(http.StatusOK, barang)
}

//hapus barang
func (a *BarangController) DeleteBarang(c echo.Context) error {
	_, err := a.BarangController.Delete(c.Param("id_barang"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Gagal hapus data"})
	}
	return c.JSON(http.StatusOK, map[string]string{"msg": "Data Berhasil Dihapus"})
}

//Edit/update barang
func (a *BarangController) UpdateBarang(c echo.Context) error {
	data := models.Barang{
		Id_barang:   c.FormValue("id_barang"),
		Nama_barang: c.FormValue("nama_barang"),
		Stok:        c.FormValue("stok"),
		Deskripsi:   c.FormValue("deskripsi"),
		Harga:       c.FormValue("harga"),
		Tgl_masuk:   c.FormValue("tgl_masuk"),
	}
	barang, err := a.BarangController.Update(data, c.Param("id_barang"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"Msg": "Gagal Meng-Update Data"})
	}
	return c.JSON(http.StatusOK, barang)
}
