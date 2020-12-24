package main

import (
	"image"
	"io"
	"os"

	ico "github.com/biessek/golang-ico"
	"github.com/jackmordaunt/icns"
)

/*writes the image at the icon path in the proper formats to the drive path, along with the needed files*/
func applyIcon(iconPath string, drivePath string) {
	//these errors all need to be caught eventually
	//this block removes existing files by the same names
	//writing over the files instead of removing seems to cause issues
	os.Remove(drivePath + "/.autorun.ico")
	os.Remove(drivePath + "/autorun.inf")
	os.Remove(drivePath + "/.VolumeIcon.icns")
	os.Remove(drivePath + "/._")
	os.Remove(drivePath + "/._.VolumeIcon.icns")

	//this block writes the windows icon and autorun file

	icon, _ := os.Open(iconPath)

	image, _, _ := image.Decode(icon)

	target, _ := os.Create(drivePath + "/.autorun.ico")

	_ = ico.Encode(target, image) //write the autorun.ico image

	autorun, _ := os.Create(drivePath + "/autorun.inf") //make the autorun.inf

	autorun.WriteString("[Autorun]\nIcon=.autorun.ico")

	//figure out how to use fatattr to hide these files on any system

	//MacOs  .VolumeIcon.icns   ._
	icnsTarget, _ := os.Create(drivePath + "/.VolumeIcon.icns")

	icns.Encode(icnsTarget, image)

	byteSource, _ := os.Open("._")
	byteTarget, _ := os.Create(drivePath + "/._")
	io.Copy(byteTarget, byteSource)

	volumeSource, _ := os.Open("._.VolumeIcon.icns")
	volumeTarget, _ := os.Create(drivePath + "/._.VolumeIcon.icns")
	io.Copy(volumeTarget, volumeSource)

	closeAll([]*os.File{target, autorun, icnsTarget, icon, byteTarget, byteSource, volumeTarget, volumeSource})

}

/*takes an []*os.File{} and closes every file
was closing enough files that i felt the need to make a function
*/
func closeAll(closeList []*os.File) {
	for _, file := range closeList {
		file.Close()
	}
}

/*
@Isaac the standard lib has the ability to read and write png, jpg, and others
decode the input image into an image.Image, then encode to the output format
this package has support for ico
https://godoc.org/github.com/biessek/golang-ico
	Package ico
	Golang .ico encoder & decoder
https://play.golang.org/p/LQICDOh5qdq
that examples can read images in ico, png, gif, and jpg format, and output the image as an ico
*/
