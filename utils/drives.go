package utils

import (
	"os"
)

// GetAllDrives retrieves a list of all logical drives on the system.
func GetAllDrives() ([]string, error) {
	var drives []string
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		path := string(drive) + ":\\"
		if _, err := os.Stat(path); err == nil {
			drives = append(drives, path)
		}
	}

	return drives, nil
}
