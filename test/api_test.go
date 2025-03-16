package test

import (
	"fmt"
	"testing"

	"github.com/rikidas/GoTv/api"
)

func TestAPI(t *testing.T) {
	apiURL := "https://iptv-org.github.io/api/streams.json" // Reemplázalo con la URL real

	// Obtener los canales
	channels, err := api.FetchChannels(apiURL)
	if err != nil {
		fmt.Println("❌ Error al obtener canales:", err)
		return
	}

	// Agrupar canales por país
	groupedChannels := api.GroupChannelsByTLD(channels)

	// Imprimir los canales organizados por país
	fmt.Println("✅ Canales organizados por país:")

	for country, chs := range groupedChannels {
		fmt.Printf("\n🌍 País: %s\n", country)
		for _, ch := range chs {
			fmt.Printf("📺 Canal: %s | 🔗 URL: %s\n", ch.Name, ch.URL)
		}
	}
}
