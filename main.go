package main

import (
	"fmt"
	"log"
	"moomonitor.je/checkers"
	"moomonitor.je/common"
	"time"

	//"github.com/shirou/gopsutil/v3/host"
	//"github.com/shirou/gopsutil/v3/mem"
	"github.com/joho/godotenv"
)

func main() {

	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Assign env variables

	// 24 hour timeout duration
	timeoutDuration := 24 * time.Hour
	// Check temperature every 5 min
	checkInterval := 5 * time.Second

	// Start infinite loop
	for {
		// If the script is in timeout, sleep for the timeout duration
		if common.IsTimeout {
			time.Sleep(timeoutDuration)
			common.IsTimeout = false
		}
		// Main Checking function
		timeRn := time.Now()
		fmt.Println("At " + timeRn.Format(time.RFC3339))
		checkers.CheckTemp()
		time.Sleep(checkInterval)
	}

}
