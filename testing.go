package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"

	"fyne.io/fyne/widget"
	"github.com/h2non/filetype"
)

/*enables or disables apply button based on status of selected drive, image
this is some poor code, ill fix it if it causes issues*/
func setApplyStatus(
	applyButton *widget.Button,
	iconPath *widget.Entry,
	selectedDrive *string,
	driveList *[]string) {

	if (*iconPath).Validate() == nil {
		for _, element := range *driveList { //test that device is still mounted
			if element == *selectedDrive {
				(*applyButton).Enable()
				log.Println("apply button was enabled")
				return
			}
		}
	}
	(*applyButton).Disable()
	log.Println("apply button was disabled")
}

/* returns nil if the path points to a file that is an image
intended for use in the validator of a fyne entry*/
func testImgPath(path string) error {
	buf, err := ioutil.ReadFile(path) //remember to actually test error
	if !os.IsNotExist(err) {
		handleErr(err)
	}
	if filetype.IsImage(buf) {
		log.Println("Valid image tested")
		return nil
	}
	return errors.New("bad") //this is so bad it hurts but ill fix it later
}

func elapsed(timed string) func() {
	start := time.Now()
	return func() { log.Printf("%s took %v\n", timed, time.Since(start)) }
}
