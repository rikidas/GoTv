package main

import (
	"fmt"

	"github.com/rikidas/GoTv/api"
	"github.com/rikidas/GoTv/player"
	"github.com/rikidas/GoTv/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println("📺 Iniciando la aplicación...")

	// Crear la aplicación Fyne
	myApp := app.New()

	// Obtener los canales desde la API
	apiURL := "https://iptv-org.github.io/api/streams.json"
	channels, err := api.FetchChannels(apiURL)
	if err != nil {
		fmt.Println("❌ Error al obtener canales:", err)
		return
	}

	// Agrupar los canales por TLD
	groupedChannels := api.GroupChannelsByTLD(channels)

	// Iniciar la UI
	ui.StartUI(myApp, groupedChannels, player.OpenVLC)
}
