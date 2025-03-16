package ui

import (
	"sort"

	"github.com/rikidas/GoTv/api"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// StartUI inicializa la interfaz gráfica
func StartUI(myApp fyne.App, groupedChannels map[string][]api.Channel, openFunc func(string)) {
	myWindow := myApp.NewWindow("TV por País")

	// Crear un slice con los nombres de los países
	var countries []string
	for country := range groupedChannels {
		countries = append(countries, country)
	}
	sort.Strings(countries) // Ordenar países alfabéticamente

	// Crear un selector de países
	countrySelect := widget.NewSelect(countries, nil)
	countrySelect.PlaceHolder = "Selecciona un país"

	// Lista de botones de canales con scroll
	channelButtons := container.NewVBox()
	scrollableChannels := container.NewVScroll(channelButtons)
	scrollableChannels.SetMinSize(fyne.NewSize(380, 400)) // Ajustamos el tamaño mínimo del scroll

	// Función para actualizar la lista de canales al seleccionar un país
	countrySelect.OnChanged = func(selectedCountry string) {
		channelButtons.Objects = nil // Limpiar lista antes de actualizar

		if selectedCountry == "" {
			return
		}

		// Obtener los canales del país seleccionado
		chList := groupedChannels[selectedCountry]

		// Crear botones para cada canal
		for _, channel := range chList {
			ch := channel // Copia local para evitar problemas de referencia
			btn := widget.NewButton(ch.Name, func() {
				openFunc(ch.URL) // Llamar a la función de reproducción
			})
			channelButtons.Add(btn)
		}

		// Refrescar la lista de botones
		channelButtons.Refresh()
	}

	// Crear la interfaz con selector y lista de canales con scroll
	content := container.NewVBox(
		widget.NewLabel("Selecciona un país:"),
		countrySelect,
		widget.NewLabel("Canales disponibles:"),
		container.NewMax(scrollableChannels), // Hacemos que el scroll use el espacio disponible
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 600))
	myWindow.SetFixedSize(true) // Evita que la ventana se haga más pequeña de lo necesario
	myWindow.ShowAndRun()
}
