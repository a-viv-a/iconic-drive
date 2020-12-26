package main

import (
	"bytes"
	"image"
	"io"
	"os"
	"sync"

	ico "github.com/biessek/golang-ico"
	"github.com/jackmordaunt/icns"
)

/*writes the image at the icon path in the proper formats to the drive path, along with the needed files*/
func applyIcon(iconPath string, drivePath string) {
	//times the icon writing process
	defer elapsed("icon writing")()
	/*these errors need to be delt with eventually
	this block removes existing files by the same names
	it should be concurrent now*/
	var wg sync.WaitGroup
	removals := []string{"/.autorun.ico", "/autorun.inf", "/.VolumeIcon.icns", "/._", "/._.VolumeIcon.icns"}
	for _, file := range removals {
		wg.Add(1)
		if err := asyncRemove(drivePath+file, &wg); !(err == nil || os.IsNotExist(err)) {
			handleErr(err)
		}
	}
	wg.Wait()

	//this block writes the windows icon and autorun file

	icon, err := os.Open(iconPath)
	handleErr(err)

	image, _, err := image.Decode(icon)
	handleErr(err)

	wg.Add(1)
	go func() {
		defer wg.Done()
		target, err := os.Create(drivePath + "/.autorun.ico")
		handleErr(err)
		handleErr(ico.Encode(target, image)) //write the autorun.ico image
		handleErr(target.Close())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		autorun, err := os.Create(drivePath + "/autorun.inf") //make the autorun.inf
		handleErr(err)
		_, err = autorun.WriteString("[Autorun]\nIcon=.autorun.ico")
		handleErr(err)
		handleErr(autorun.Close())
	}()

	//figure out how to use fatattr to hide these files on any system

	//MacOs block

	wg.Add(1)
	go func() {
		defer wg.Done()
		icnsTarget, err := os.Create(drivePath + "/.VolumeIcon.icns")
		handleErr(err)
		handleErr(icns.Encode(icnsTarget, image))
		handleErr(icnsTarget.Close())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		byteSource := bytes.NewReader(MustAsset("data/._"))
		byteTarget, err := os.Create(drivePath + "/._")
		handleErr(err)
		_, err = io.Copy(byteTarget, byteSource)
		handleErr(err)
		handleErr(byteTarget.Close())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		volumeSource := bytes.NewReader(MustAsset("data/._.VolumeIcon.icns"))
		volumeTarget, err := os.Create(drivePath + "/._.VolumeIcon.icns")
		handleErr(err)
		_, err = io.Copy(volumeTarget, volumeSource)
		handleErr(err)
		handleErr(volumeTarget.Close())
	}()

	wg.Wait()

	handleErr(icon.Close())

}

/*takes an []*os.File{} and closes every file
was closing enough files that i felt the need to make a function*/
func closeAll(closeList []*os.File) {
	for _, file := range closeList {
		handleErr(file.Close())
	}
}

func asyncRemove(path string, wg *sync.WaitGroup) error {
	defer wg.Done()
	return os.Remove(path)
}
