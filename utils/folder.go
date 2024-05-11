package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CheckLogFolder(leagueClientPath string) error {
	logFolderPath := filepath.Join(filepath.Dir(leagueClientPath), "Logs")

	if _, err := os.Stat(logFolderPath); os.IsNotExist(err) {
		return fmt.Errorf("log folder not found at: %s", logFolderPath)
	} else {
		fmt.Printf("Log folder found at: %s\n and delete all Files!", logFolderPath)

		files, err := os.ReadDir(logFolderPath)
		if err != nil {
			return fmt.Errorf("failed to read the log folder: %v", err)
		}

		// Iterate through the contents of the folder
		for _, file := range files {
			if file.IsDir() {
				folderPath := filepath.Join(logFolderPath, file.Name())

				// Remove the directory and its contents
				if err := os.RemoveAll(folderPath); err != nil {
					fmt.Printf("Failed to delete folder: %s, error: %v\n", folderPath, err)
				} else {
					fmt.Printf("Successfully deleted folder: %s\n", folderPath)
				}

			}
		}
	}
	return nil
}
