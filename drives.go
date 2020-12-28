//+build linux
//+build darwin

package main

import (
	"log"
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

	for _, drive := range driveList {
		log.Println(drive + " is mapped to " + driveMap[drive])
	}

	return driveList, driveMap
}
