package main

import (
	"image/png"
	"os"

	ico "github.com/biessek/golang-ico"
)

/*writes the image at the icon path in the proper formats to the drive path, along with the needed files*/
func applyIcon(iconPath string, drivePath string) {
	println("writeing\n" + iconPath + "\nto\n" + drivePath)

	icon, err := os.Open(iconPath)
	if err != nil {
		println(err)
	}

	image, _ := png.Decode(icon)

	target, _ := os.Create(drivePath + "/test.ico")

	_ = ico.Encode(target, image)

	icon.Close()
	target.Close()
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
