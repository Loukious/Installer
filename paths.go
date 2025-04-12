package main

import (
	"os"
	path "path/filepath"
	"runtime"
)

func getVencordAsarPath() string {
	// 1. Check environment variable (user override)
	if customPath := os.Getenv("VENCORD_USER_DATA_DIR"); customPath != "" {
		return path.Join(customPath, "vencord.asar")
	}

	// 2. Platform-specific default paths
	switch runtime.GOOS {
	case "windows":
		return path.Join(os.Getenv("APPDATA"), "Vencord", "vencord.asar")
	case "darwin":
		return path.Join(os.Getenv("HOME"), "Library", "Application Support", "Vencord", "vencord.asar")
	default: // Linux/BSD
		if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
			return path.Join(xdg, "vencord", "vencord.asar")
		}
		return path.Join(os.Getenv("HOME"), ".local", "share", "vencord", "vencord.asar")
	}
}

var VencordAsarPath = getVencordAsarPath()
