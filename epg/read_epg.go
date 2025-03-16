package epg

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

// ReadEPGFile descarga y descomprime el archivo XML de la EPG
func ReadEPGFile(url string) (string, error) {
	// Descargar el archivo EPG
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error descargando la EPG: %v", err)
	}
	defer resp.Body.Close()

	// Crear un lector gzip
	gzReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error descomprimiendo el archivo: %v", err)
	}
	defer gzReader.Close()

	// Leer el contenido descomprimido
	xmlData, err := io.ReadAll(gzReader)
	if err != nil {
		return "", fmt.Errorf("error leyendo el XML: %v", err)
	}

	return string(xmlData), nil
}
