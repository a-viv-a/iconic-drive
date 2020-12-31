//+build !windows

package main

import (
	"os"

	"golang.org/x/sys/unix"
)

//ATTR_HIDDEN is the hidden atribute for fat systems
const ATTR_HIDDEN = 2

//FAT_IOCTL_SET_ATTRIBUTES is write mode
const FAT_IOCTL_SET_ATTRIBUTES = 0x80047211

func hideFile(fileStr string, wg *WaitGroupBar) {
	defer wg.Done()

	openFile, err := os.Open(fileStr)
	handleErr(err)
	defer openFile.Close()

	handleErr(unix.IoctlSetInt(int(openFile.Fd()), FAT_IOCTL_SET_ATTRIBUTES, ATTR_HIDDEN))
	//bunch of magic I dont really understand
	//much thanks to chanbakjsd#7968 from the gophers discord
}
