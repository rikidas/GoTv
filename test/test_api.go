package test

import (
	"fmt"
	"testing"

	"github.com/rikidas/GoTv/api"
)

// Test para verificar la obtención de canales
func TestFetchChannels(t *testing.T) {
	apiURL := "https://iptv-org.github.io/api/streams.json"

	channels, err := api.FetchChannels(apiURL)
	if err != nil {
		t.Errorf("Error al obtener canales: %v", err)
	}

	fmt.Println("Lista de canales obtenida correctamente:")
	for _, ch := range channels {
		fmt.Printf(" Nombre: %s | URL: %s\n", ch.channel, ch.URL)
	}
}
