package test

import (
	"testing"

	"github.com/rikidas/GoTv/epg"
)

// TestReadEPGFile verifica que la función puede descargar y leer la EPG correctamente
func TestReadEPGFile(t *testing.T) {
	url := "https://epgshare01.online/epgshare01/epg_ripper_TR1.xml.gz"

	epgData, err := epg.ReadEPGFile(url)
	if err != nil {
		t.Fatalf("Error al leer la EPG: %v", err)
	}

	if len(epgData) == 0 {
		t.Fatalf("EPG vacía, se esperaba contenido XML")
	}

	t.Logf("EPG cargada correctamente. Primeros caracteres: %s", epgData[:200])
}
