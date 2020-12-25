package main

import (
	"path/filepath"

	usbdrivedetector "github.com/deepakjois/gousbdrivedetector"
)

//returns human readable name for each drive, then the path in a map
func drives() ([]string, map[string]string) {
	driveList, err := usbdrivedetector.Detect()
	handleErr(err)
	driveMap := make(map[string]string)
	for i, drive := range driveList {
		driveMap[filepath.Base(drive)] = driveList[i]
		driveList[i] = filepath.Base(drive)
	}
	return driveList, driveMap
}
