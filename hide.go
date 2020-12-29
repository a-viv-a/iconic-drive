//+build !windows

package main

import "log"

func hideFile(file string, wg *WaitGroupBar) {
	defer wg.Done()
	log.Printf("hide method not implemented for this platform. %s cannot be hidden", file)
}
