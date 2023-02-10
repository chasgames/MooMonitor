package checkers

import (
	"fmt"

	common "moomonitor/common"
	"github.com/shirou/gopsutil/v3/host"
)

func CheckTemp() {
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
			if t.Temperature > common.Threshold {
				tripped = true
			}
			result += row
		}
		if tripped {
			//Send Push
			common.PushNotification(result, "Temperature too Hot 🔥")
			//invoke Timeout
			common.IsTimeout = true
			fmt.Println("timeout invoked")
		}
		fmt.Println(result)

}