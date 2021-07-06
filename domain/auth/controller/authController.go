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

	user := c.FormValue("username")
	pass := c.FormValue("password")

	auth, err := a.AuthController.Cek(user, pass)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"msg": "Error Login"})
	}
	return c.JSON(http.StatusCreated, auth)
}
