package api

import (
	"strings"
)

// GetCountryCode extrae el código de país de una URL.
func GetCountryCode(url string) string {
	// Separar la URL por "."
	parts := strings.Split(url, ".")

	// Si tiene al menos dos partes, tomamos la última como el país
	if len(parts) > 1 {
		return parts[len(parts)-1]
	}

	// Si no se puede determinar, devolver "Desconocido"
	return "Desconocido"
}
