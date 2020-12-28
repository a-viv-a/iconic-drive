// +build windows

package main

import (
	"fmt"
	"log"
	"path/filepath"

	usbdrivedetector "github.com/deepakjois/gousbdrivedetector"
	"github.com/winlabs/gowin32"
)

//returns human readable name for each drive, then the path in a map - windows syscalls version
func drives() ([]string, map[string]string) {
	driveList, err := usbdrivedetector.Detect()
	handleErr(err)
	driveMap := make(map[string]string)
	for i, drive := range driveList {
		/*driveMap[filepath.Base(drive)] = driveList[i]
		driveList[i] = filepath.Base(drive)*/

		driveInfo, _ := gowin32.GetVolumeInfo(drive)
		driveName := fmt.Sprintf("%s (%s)", driveInfo.VolumeName, filepath.VolumeName(drive))

		driveMap[driveName] = driveList[i]
		driveList[i] = driveName
	}

	for _, drive := range driveList {
		log.Println(drive + " is mapped to " + driveMap[drive])
	}

	return driveList, driveMap
}
