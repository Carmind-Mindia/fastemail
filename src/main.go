package main

import (
	"github.com/Carmind-Mindia/fastemail/src/manager"
	"github.com/Carmind-Mindia/fastemail/src/service"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	//Corremos el deamon con el channel
	go manager.DeamonEmail()
	// go manager.DeamonNotification()

	//Creamos la api
	emailApi := service.NewApiEmail()
	notificationApi := service.NewApiNotification()

	//Routeamos
	e.POST("/sendRecoverPassword", emailApi.SendRecoverPassword)

	e.POST("/sendNotificationToCarmind", notificationApi.SendNotificationToCarmind)

	//Start!
	e.Logger.Fatal(e.Start(":5896"))
}
