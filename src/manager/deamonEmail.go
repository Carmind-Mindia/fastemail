package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/Carmind-Mindia/fastemail/src/model"

	"bytes"
	"net/http"
)

var (
	EmailChannel chan model.EmailTemplate
)

var listMapEmail = make(map[string]time.Time)

func DeamonEmail() {
	// Obtener el bearer token del archivo .env
	bearerToken := os.Getenv("ENVIALOSIMPLE_TOKEN")

	// Creamos el channel
	EmailChannel = make(chan model.EmailTemplate)

	for {
		// Esperamos un dato del canal
		data := <-EmailChannel

		err := processEmail(data.EmailTo)
		if err != nil {
			// TODO: logeamos el error
			fmt.Print(err)
			continue
		}

		var body model.EnvialoSimpleApiBody
		body.Html = data.Html
		body.To = data.EmailTo
		body.From = "ayuda@mindia.com.ar"
		body.Subject = data.Title

		// Convertir los datos de la solicitud a JSON
		jsonData, err := json.Marshal(body)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Crear una nueva solicitud POST
		req, err := http.NewRequest("POST", "https://api.envialosimple.email/api/v1/mail/send", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Establecer el encabezado Content-Type a application/json
		req.Header.Set("Content-Type", "application/json")

		// Agregar el bearer token al encabezado Authorization
		req.Header.Set("Authorization", "Bearer "+bearerToken)

		// Realizar la llamada a la API
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Comprobar el código de estado de la respuesta
		if resp.StatusCode != http.StatusOK {
			fmt.Println(fmt.Errorf("API call failed with status code: %d", resp.StatusCode))
			// Leer el cuerpo de la respuesta
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("No se pudo leer el cuerpo de la respuesta")
				continue
			}

			// Imprimir el cuerpo de la respuesta
			fmt.Println(string(body))

			resp.Body.Close()
		}
	}
}

func processEmail(email string) error {
	//Obtenemos el tiempo actual
	now := time.Now()

	//Le agregamos un minuto
	OneMinuteAgo := now.Add(-(time.Second * time.Duration(59)))

	for k, t := range listMapEmail {
		//Si el tiempo guardado en los logs, es de hace mas de un minuto, lo borramos
		if OneMinuteAgo.After(t) {
			delete(listMapEmail, k)
		}
	}

	//Verificamos si esta el correo a mandar el email
	if _, ok := listMapEmail[email]; ok {
		//Si esta, deberiamos tirar error
		return errors.New("intentelo mas tarde")
	} else {
		//Si no esta, lo dejamos proseguir y guardamos el log
		listMapEmail[email] = time.Now()
	}

	return nil
}
