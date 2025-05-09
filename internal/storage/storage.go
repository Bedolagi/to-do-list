package storage

import (
	"os"
	"path/filepath"
)

func InitConfig() string {
	var configPath string
	configPath = filepath.Join(os.Getenv("APPDATA"), "To-Do-List")
	os.Mkdir(configPath, 0755)

	return configPath
}
