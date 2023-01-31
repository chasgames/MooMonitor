package main

import (
    "fmt"
	"time"
	"log"
	"os"

	"github.com/shirou/gopsutil/v3/host"
	//"github.com/shirou/gopsutil/v3/mem"
	"github.com/gregdel/pushover"
	"github.com/joho/godotenv"
)

const threshold = 60.0

// Determine if script is in a timeout period
var isTimeout = false

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
		if isTimeout {
			time.Sleep(timeoutDuration)
			isTimeout = false
		}
		// Main Checking function
		timeRn := time.Now()
		fmt.Println("At " + timeRn.String())
		checkTemp()
		time.Sleep(checkInterval)
	}

}

func checkTemp() {
		t, err := host.SensorsTemperatures()
		if err != nil {
			fmt.Println(err)
		}
		var result string
		var row string
		tripped := false
		//clear results so they don't build up
		result, row = "", ""
		// Loop over all sensors
		for _, t := range t {
			row = fmt.Sprint(t.SensorKey,": ", t.Temperature, "°C \n")
			if t.Temperature > threshold {
				tripped = true
			}
			result += row
		}
		if tripped {
			//Send Push
			pushNotification(result)
			//invoke Timeout
			isTimeout = true
			fmt.Println("timeout invoked")
		}
		fmt.Println(result)

}

func pushNotification(inc string) {
	//Get envars
	pushoverAppKey := os.Getenv("PUSHOVER_APPKEY")
	pushoverRecipient := os.Getenv("PUSHOVER_RECIPIENT")
	// Pushover Application API KEY
	var app = pushover.New(pushoverAppKey)
	// Recipent ID (Check on device, or create a delivery group for multiple)
	var recipient = pushover.NewRecipient(pushoverRecipient)

	// Get Hostname of system for identification
	h, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	machine := fmt.Sprintf("%v (%v - %v)\n", h.Hostname, h.OS, h.Platform)
	// Create Push Message
	content := fmt.Sprintf(machine + inc)
	message := pushover.NewMessageWithTitle(content, "Temperature too Hot 🔥")

    // Send the message to the recipient
    response, err := app.SendMessage(message, recipient)
    if err != nil {
        fmt.Println(err)
    }

    // Print the response if you want
    fmt.Println(response)
}