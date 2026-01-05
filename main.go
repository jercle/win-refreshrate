package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ncruces/zenity"
)

// go build -ldflags -H=windowsgui -o getRefreshRate.exe .
// go build -ldflags -H=windowsgui -o switchRefreshRate.exe .

func main() {
	// switchRefreshRate()
	getRefreshRate()
}

func switchRefreshRate() {
	currentRefreshRate := getCurrentRefreshRate()
	currentRRStr := strconv.Itoa(int(currentRefreshRate))

	var newRefreshRate uint32

	if currentRefreshRate == 60 {
		newRefreshRate = 120
	} else if currentRefreshRate == 120 {
		newRefreshRate = 60
	} else {
		zenity.Error("Current refresh rate is not 60Hz or 120Hz. Please check settings. Current refresh rate is " + currentRRStr + "Hz.")
		os.Exit(1)
	}
	newRRStr := strconv.Itoa(int(newRefreshRate))

	rsp := zenity.Question(
		"Current refresh rate is "+currentRRStr+"Hz. Do you want to change to "+newRRStr+"Hz?",
		zenity.Title("Are you sure?"),
	)

	if rsp != nil {
		os.Exit(0)
	}

	err := changeRefreshRate(newRefreshRate)

	if err != nil {
		zenity.Error(fmt.Sprint(err))
	}

	time.Sleep(2 * time.Second)

	updatedRefreshRate := getCurrentRefreshRate()

	zenity.Notify(fmt.Sprintf("%d Hz\n", updatedRefreshRate),
		zenity.Title("Updated Refresh Rate"),
	)
}

func getRefreshRate() {
	currentRefreshRate := getCurrentRefreshRate()

	zenity.Notify(fmt.Sprintf("%d Hz\n", currentRefreshRate),
		zenity.Title("Current Refresh Rate"),
	)
}
