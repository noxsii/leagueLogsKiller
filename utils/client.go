package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindLeagueClient searches for the League of Legends executable across all drives.
func FindLeagueClient(drives []string) (string, error) {
	var leagueClient string
	var found bool

	for _, baseDir := range drives {
		if found {
			break
		}

		err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return nil
			}
			if info.IsDir() && (info.Name() == "Windows" || info.Name() == "Program Files" || info.Name() == "Program Files (x86)") {
				return filepath.SkipDir
			}
			if filepath.Base(path) == "LeagueClient.exe" {
				leagueClient = path
				found = true
				return filepath.SkipDir
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking through %s: %s\n", baseDir, err)
		}
	}

	if !found {
		return "", fmt.Errorf("LeagueClient.exe not found on any drive")
	}

	return leagueClient, nil
}
