<<<<<<< HEAD
package controller

import (
	"kasir-clean/domain/auth/models"
	"net/http"

	"github.com/labstack/echo"
)

type AuthController struct {
	AuthController models.AuthEntity
}

func AuthControllerFunc(c *echo.Group, us models.AuthEntity) {
	handler := &AuthController{
		AuthController: us,
	}

	c.POST("/login", handler.Login)
}
func (a *AuthController) Login(c echo.Context) error {
	data := models.Auth{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	auth, err := a.AuthController.Cek(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Error Login"})
	}
	return c.JSON(http.StatusCreated, auth)
}
=======
package controller

import (
	"kasir-clean/domain/auth/models"
	"net/http"

	"github.com/labstack/echo"
)

type AuthController struct {
	AuthController models.AuthEntity
}

func AuthControllerFunc(c *echo.Group, us models.AuthEntity) {
	handler := &AuthController{
		AuthController: us,
	}

	c.POST("/login", handler.Login)
}
func (a *AuthController) Login(c echo.Context) error {
	data := models.Auth{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	auth, err := a.AuthController.Cek(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Error Login"})
	}
	return c.JSON(http.StatusCreated, auth)
}
>>>>>>> master
