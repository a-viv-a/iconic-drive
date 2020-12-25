package main

import (
	"log"
	"path/filepath"
	"runtime"

	usbdrivedetector "github.com/deepakjois/gousbdrivedetector"
)

//returns human readable name for each drive, then the path in a map
func drives() ([]string, map[string]string) {
	driveList, err := usbdrivedetector.Detect()
	handleErr(err)
	driveMap := make(map[string]string)
	if runtime.GOOS != "windows" {
		for i, drive := range driveList {
			driveMap[filepath.Base(drive)] = driveList[i]
			driveList[i] = filepath.Base(drive)
		}
	} else {
		log.Println("Using windows bandaid drive name fix")
		for i, drive := range driveList {
			//this is a bandaid fix
			//shows drive letter for windows
			driveMap[drive] = driveList[i]
			driveList[i] = drive
		}
	}

	for _, drive := range driveList {
		log.Println(drive + " is mapped to " + driveMap[drive])
	}

	return driveList, driveMap
}
