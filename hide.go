//+build !windows

package main

import (
	"os"

	"golang.org/x/sys/unix"
)

//ATTR_HIDDEN is the hidden atribute for fat systems
const ATTR_HIDDEN int = 2

//FAT_IOCTL_GET_ATTRIBUTES is read mode
const FAT_IOCTL_GET_ATTRIBUTES uint = 2147774992

//FAT_IOCTL_SET_ATTRIBUTES is write mode
const FAT_IOCTL_SET_ATTRIBUTES uint = 1074033169

func hideFile(fileStr string, wg *WaitGroupBar) {
	defer wg.Done()

	openFile, err := os.Open(fileStr)
	handleErr(err)
	defer openFile.Close()

	current, err := unix.IoctlGetInt(int(openFile.Fd()), FAT_IOCTL_GET_ATTRIBUTES)
	handleErr(err)

	//log.Printf("Current flags are: %d", current)

	handleErr(unix.IoctlSetPointerInt(int(openFile.Fd()), FAT_IOCTL_SET_ATTRIBUTES, current|ATTR_HIDDEN))
	//bunch of magic I dont really understand
	//much thanks to chanbakjsd#7968 and _diamondburned_#4507 from the gophers discord
}
