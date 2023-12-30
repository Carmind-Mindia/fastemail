package service

import (
	"net/http"

	"github.com/Carmind-Mindia/fastemail/src/manager"
	"github.com/Carmind-Mindia/fastemail/src/model"
	"github.com/labstack/echo"
)

type ApiEmail struct {
	manager manager.EmailManager
}

func NewApiEmail() ApiEmail {
	m := manager.NewEmailManager()
	return ApiEmail{manager: m}
}

func (api *ApiEmail) SendRecoverPassword(c echo.Context) error {
	data := model.RecuperarContraseña{}
	c.Bind(&data)

	err := api.manager.SendRecoverPassword(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.NoContent(http.StatusOK)
}
