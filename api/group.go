package api

import (
	"fmt"
	"regexp"
	"strings"
)

// GroupChannelsByTLD agrupa los canales por país usando el TLD de la URL
func GroupChannelsByTLD(channels []Channel) map[string][]Channel {
	groups := make(map[string][]Channel)

	// Expresión regular para extraer el TLD
	tldRegex := regexp.MustCompile(`\.(\w{2,3})(?:/|$)`)

	// ...existing code...

	for _, channel := range channels {
		if channel.Name == "" {
			fmt.Println("⚠️ Canal sin nombre, no se agrupará:", channel.URL)
			continue
		}

		matches := tldRegex.FindStringSubmatch(channel.Name)
		if len(matches) > 1 {
			tld := strings.ToUpper(matches[1]) // Convierte a mayúsculas para estandarizar
			groups[tld] = append(groups[tld], channel)
		} else {
			fmt.Println("⚠️ No se pudo extraer el TLD de:", channel.URL)
		}
	}

	// ...existing code...

	return groups
}
