package main

import (
	"log"
	"syscall"
)

func hideFile(file string, wg *WaitGroupBar) {
	defer wg.Done()
	log.Printf("Hiding %s with win32", file)

	fileW, err := syscall.UTF16PtrFromString(file)
	handleErr(err)
	handleErr(syscall.SetFileAttributes(fileW, syscall.FILE_ATTRIBUTE_HIDDEN))

}
