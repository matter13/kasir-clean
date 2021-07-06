<<<<<<< HEAD
package main

import (
	"kasir-clean/config"
	"kasir-clean/domain/karyawan/controller"
	"kasir-clean/domain/karyawan/repository"
	"kasir-clean/domain/karyawan/usecase"
	
	controllerB "kasir-clean/domain/karyawan/controller"
	repositoryB "kasir-clean/domain/karyawan/repository"
	usecaseB "kasir-clean/domain/karyawan/usecase"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	db := config.Connect()
	
	//karyawan
	repo := repository.NewKaryawanRepository(db)
	entity := usecase.NewKaryawanEntity(repo)
	
	//barang
	repoK := repositoryB.NewKaryawanRepository(db)
	entityK := usecaseB.NewKaryawanEntity(repoK)

	api := e.Group("/api")
	controller.KaryawanControllerFunc(api, entity)
	controllerB.KaryawanControllerFunc(api, entityK)

	//repoK := repo

	e.Start(":7000")
}
=======
package main

import (
	//koneksi
	"io"
	"kasir-clean/config"
	"net/http"
	"text/template"

	//karyawan
	"kasir-clean/domain/karyawan/controller"
	"kasir-clean/domain/karyawan/repository"
	"kasir-clean/domain/karyawan/usecase"

	//barang
	controllerB "kasir-clean/domain/barang/controller"
	repositoryB "kasir-clean/domain/barang/repository"
	usecaseB "kasir-clean/domain/barang/usecase"

	"github.com/labstack/echo"
)

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()
	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	db := config.Connect()

	//karyawan
	repo := repository.NewKaryawanRepository(db)
	entity := usecase.NewKaryawanEntity(repo)
	//barang
	repoK := repositoryB.NewBarangRepository(db)
	entityK := usecaseB.NewBarangEntity(repoK)

	e.Renderer = NewRenderer("view/index.html", true)

	api := e.Group("/api")
	e.GET("/", LoginHandler)
	controller.KaryawanControllerFunc(api, entity)
	controllerB.BarangControllerFunc(api, entityK)

	//repoK := repo

	e.Start(":7000")

}
func LoginHandler(c echo.Context) error {
	data := map[string]string{"msg": "Login Page"}
	return c.Render(http.StatusOK, "index.html", data)
}
>>>>>>> master
