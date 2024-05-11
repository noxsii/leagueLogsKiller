package main

import (
	"fmt"
	"log"
	"noxsi/league-logs-killer/utils"
	"os"
)

func main() {
	green := "\033[32m"
	reset := "\033[0m"

	// Create a log file to log errors
	logFile, err := os.OpenFile("league_logs_killer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could not create log file:", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	drives, err := utils.GetAllDrives()
	fmt.Printf("%sThe Script search League of Legends... %s\n", green, reset)
	if err != nil {
		fmt.Println("Failed to get drives:", err)
		return
	}

	// Call the function to find the LeagueClient.exe file.
	leagueClientPath, err := utils.FindLeagueClient(drives)

	if err != nil {
		fmt.Println("Error:", err)
	} else {

		fmt.Printf("League of Legends executable found at: %s\n", leagueClientPath)
	}

	// Find Log Folder
	utils.CheckLogFolder(leagueClientPath)
	if err != nil {
		log.Println("Error checking log folder:", err)
		fmt.Println("Error checking log folder. Please see log file for more details.")
		return
	}
}
