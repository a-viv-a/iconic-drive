package main

import (
	"log"
	"runtime"
)

func handleErr(err error) {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		details := runtime.FuncForPC(pc)
		if ok && details != nil {
			log.Printf("%s <- was called from %s in %s on line %d\n", err.Error(), details.Name(), file, line)
		} else if ok {
			log.Printf("%s <- was called from ~~~ in %s on line %d\n", err.Error(), file, line)
		} else {
			log.Printf("%s <- runtime.Caller was not ok.\n", err.Error())
		}
	}
}
