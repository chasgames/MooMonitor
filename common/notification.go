package common

import (
		"fmt"
		"os"
		"github.com/shirou/gopsutil/v3/host"
		"github.com/gregdel/pushover"

)

func PushNotification(msgContents string, msgTitle string) {
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
	content := fmt.Sprintf(machine + msgContents)
	message := pushover.NewMessageWithTitle(content, msgTitle)

    // Send the message to the recipient
    response, err := app.SendMessage(message, recipient)
    if err != nil {
        fmt.Println(err)
    }

    // Print the response if you want
    fmt.Println(response)
}