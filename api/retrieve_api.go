package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// Definir la estructura que coincide con la API
type Channel struct {
	Name         string  `json:"channel"`
	URL          string  `json:"url"`
	TimeShift    *string `json:"timeshift"`
	HTTPReferrer *string `json:"http_referrer"`
	UserAgent    *string `json:"user_agent"`
}

// Función para obtener los canales desde la API
func FetchChannels(apiURL string) ([]Channel, error) {
	// Hacer la petición HTTP
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Crear una variable para almacenar los canales
	var channels []Channel

	// Decodificar JSON en la estructura
	err = json.Unmarshal(body, &channels)
	if err != nil {
		return nil, err
	}

	return channels, nil
}
