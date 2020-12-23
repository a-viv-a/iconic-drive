package main

import (
	"image"
	"os"

	ico "github.com/biessek/golang-ico"
)

/*writes the image at the icon path in the proper formats to the drive path, along with the needed files*/
func applyIcon(iconPath string, drivePath string) {
	println("writeing\n" + iconPath + "\nto\n" + drivePath)

	//these errors need to be caught eventually
	//this block writes the windows icon and autorun file

	icon, _ := os.Open(iconPath)

	image, _, _ := image.Decode(icon)

	target, _ := os.Create(drivePath + "/autorun.ico")

	_ = ico.Encode(target, image) //write the autorun.ico image

	autorun, _ := os.Create(drivePath + "/autorun.inf") //make the autorun.inf

	autorun.WriteString("[Autorun]\nIcon=autorun.ico")

	//not sure if i need to do this
	icon.Close()
	target.Close()
	autorun.Close()

	//figure out how to use fatattr to hide these files on any system
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
