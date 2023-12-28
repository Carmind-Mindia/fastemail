package service

import (
	"net/http"

	"github.com/Carmind-Mindia/fastemail/src/manager"
	"github.com/Carmind-Mindia/fastemail/src/model"
	"github.com/labstack/echo"
)

type ApiNotification struct {
	manager manager.NotificationManager
}

func NewApiNotification() ApiNotification {
	m := manager.NewNotificationManager()
	return ApiNotification{manager: m}
}

func (api *ApiNotification) SendNotificationToCarmind(c echo.Context) error {
	data := model.SimpleNotification{}
	c.Bind(&data)

	api.manager.SendNotificationToCarmind(data)

	return c.NoContent(http.StatusOK)
}
