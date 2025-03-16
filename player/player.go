package player

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenVLC abre una URL en VLC Media Player
func OpenVLC(url string) {
	fmt.Println("📺 Reproduciendo en VLC:", url)
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "vlc", url) // VLC en Windows
	case "darwin":
		cmd = exec.Command("open", "-a", "VLC", url) // VLC en macOS
	case "linux":
		cmd = exec.Command("vlc", url) // VLC en Linux
	}

	if cmd != nil {
		err := cmd.Start()
		if err != nil {
			fmt.Println("❌ Error al abrir VLC:", err)
		}
	}
}
