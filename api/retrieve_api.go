package api

import (
	"encoding/json"
	"io"
	"net/http"
)

// Estructura para los datos del canal
type Channel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Estructura para la respuesta de la API
type ApiResponse struct {
	Channels []Channel `json:"channels"`
}

// Función para obtener los canales de la API externa
func FetchChannels(apiURL string) ([]Channel, error) {
	// Hacer la solicitud HTTP GET
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

	// Decodificar JSON en la estructura ApiResponse
	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response.Channels, nil
}
