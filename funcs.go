package main

import (
	"fmt"
	"strconv"
	"unsafe"

	"github.com/gonutz/w32/v2"
)

func getCurrentRefreshRate() uint32 {
	deviceName := (*uint16)(nil)

	var devMode w32.DEVMODE
	devMode.DmSize = uint16(unsafe.Sizeof(devMode))

	// Call EnumDisplaySettingsW to get the current settings (ENUM_CURRENT_SETTINGS = -1)
	w32.EnumDisplaySettingsEx(deviceName, w32.ENUM_CURRENT_SETTINGS, &devMode, uint32(0))

	// jsonStr, _ := json.MarshalIndent(devMode, "", "  ")
	// fmt.Println(string(jsonStr))

	// The dmDisplayFrequency field contains the refresh rate in Hz
	// fmt.Printf("Current refresh rate: %d Hz\n", devMode.DmDisplayFrequency)

	return devMode.DmDisplayFrequency
}

func changeRefreshRate(refreshRate uint32) error {

	if refreshRate != 120 && refreshRate != 60 {
		return fmt.Errorf("Refresh rate must be either 120 or 60")
	}

	deviceName := (*uint16)(nil)

	var devMode w32.DEVMODE
	var hwnd w32.HWND
	devMode.DmSize = uint16(unsafe.Sizeof(devMode))

	// Call EnumDisplaySettingsW to get the current settings (ENUM_CURRENT_SETTINGS = -1)
	w32.EnumDisplaySettingsEx(deviceName, w32.ENUM_CURRENT_SETTINGS, &devMode, uint32(0))

	// jsonStr, _ := json.MarshalIndent(devMode, "", "  ")
	// fmt.Println(string(jsonStr))

	// The dmDisplayFrequency field contains the refresh rate in Hz
	fmt.Printf("Current refresh rate: %d Hz\n", devMode.DmDisplayFrequency)

	newMode := devMode

	newMode.DmDisplayFrequency = refreshRate

	var lParam uintptr

	rsp := w32.ChangeDisplaySettingsEx(deviceName, &newMode, hwnd, uint32(0), lParam)

	if rsp != 0 {
		return fmt.Errorf("Failed with code:", strconv.Itoa(int(rsp)))
	} else {
		return nil
	}
}
