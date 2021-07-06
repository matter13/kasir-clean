package controller

import (
	"kasir-clean/domain/karyawan/models"
	"net/http"

	"github.com/labstack/echo"
)

type KaryawanController struct {
	KaryawanController models.KaryawanEntity
}

func KaryawanControllerFunc(c *echo.Group, us models.KaryawanEntity) {
	handler := &KaryawanController{
		KaryawanController: us,
	}

	c.GET("/karyawan", handler.GetKaryawan)
	c.POST("/karyawan", handler.TambahKaryawan)
	c.GET("/karyawan/:id_karyawan", handler.ViewKaryawan)
	c.DELETE("/karyawan/:id_karyawan", handler.DeleteKaryawan)
	c.PUT("/karyawan/:id_karyawan", handler.EditKaryawan)
}

func (a *KaryawanController) GetKaryawan(c echo.Context) error {
	karyawan, _ := a.KaryawanController.Get(c.Request().Context())
	return c.JSON(http.StatusOK, karyawan)
}

func (a *KaryawanController) TambahKaryawan(c echo.Context) error {
	data := models.Karyawan{
		Id_karyawan: c.FormValue("id_karyawan"),
		Nama:        c.FormValue("nama"),
		Alamat:      c.FormValue("alamat"),
		No_telp:     c.FormValue("no_telp"),
	}
	karyawan, err := a.KaryawanController.Tambah(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Error input data"})
	}
	return c.JSON(http.StatusCreated, karyawan)
}

func (a *KaryawanController) ViewKaryawan(c echo.Context) error {
	karyawan, _ := a.KaryawanController.View(c.Param("id_karyawan"))
	return c.JSON(http.StatusOK, karyawan)
}

func (a *KaryawanController) DeleteKaryawan(c echo.Context) error {
	_, err := a.KaryawanController.Delete(c.Param("id_karyawan"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Gagal Hapus Data"})
	}
	return c.JSON(http.StatusOK, map[string]string{"msg": "Data Terhapus!"})
}

func (a *KaryawanController) EditKaryawan(c echo.Context) error {
	data := models.Karyawan{
		Id_karyawan: c.FormValue("id_karyawan"),
		Nama:        c.FormValue("nama"),
		Alamat:      c.FormValue("alamat"),
		No_telp:     c.FormValue("no_telp"),
	}
	karyawan, err := a.KaryawanController.Edit(data, c.Param("id_karyawan"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Ggal Edit Data"})
	}
	return c.JSON(http.StatusOK, karyawan)
}
