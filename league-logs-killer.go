package main

import (
	"fmt"
	"noxsi/league-logs-killer/utils"
)

func main() {
	drives, err := utils.GetAllDrives()
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
}
