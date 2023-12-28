package manager

import (
	"github.com/Carmind-Mindia/fastemail/src/model"
)

type NotificationManager struct {
}

func NewNotificationManager() NotificationManager {
	return NotificationManager{}
}

func (ma *NotificationManager) SendNotificationToCarmind(data model.SimpleNotification) {
	embudo := NotificationChannel

	//Creamos el personalization con el to y la data dinamica
	simpleNotification := map[string]interface{}{
		"Title":   data.Title,
		"Message": data.Message,
	}

	notification := model.CarmindNotification{
		To:      data.To,
		Prioity: "HIGH",
		Data:    simpleNotification,
	}

	//Enviamos el email al deamon para que se despache
	embudo <- notification
}
