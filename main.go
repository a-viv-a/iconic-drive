package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("iconic drive")
	//w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1, 500))

	iconPath := widget.NewEntry()
	iconPath.SetPlaceHolder("paste or type image path")
	iconPath.Validator = testImgPath
	clearButton := widget.NewButton("clear", func() { iconPath.SetText("") })
	pathWrapper := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, nil, clearButton),
		container.NewHScroll(iconPath), clearButton,
	)

	driveList, driveMap := drives()
	selectedDrive := "" //nil makes an error, this is bandaid solve

	applyButton := widget.NewButton("apply", nil)
	applyButton.Disable()

	driveSelect := widget.NewSelect(driveList, func(s string) {
		println(s + " <-> " + driveMap[s])
		selectedDrive = s
		setApplyStatus(applyButton, iconPath, &selectedDrive, &driveList)
	})

	driveSelect.PlaceHolder = "select target drive"
	refreshButton := widget.NewButton("refresh",
		func() {
			driveList, driveMap = drives()
			driveSelect.Options = driveList
			driveSelect.Refresh()
		})
	driveWrapper := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, nil, refreshButton),
		driveSelect, refreshButton,
	)

	header := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		pathWrapper,
		driveWrapper)

	preview := canvas.NewImageFromFile("error.svg")
	preview.FillMode = canvas.ImageFillContain
	preview.SetMinSize(fyne.NewSize(64, 64))

	iconPath.OnChanged = func(s string) {
		if iconPath.Validate() != nil {
			//https://www.iconfinder.com/icons/381599/error_icon
			s = "error.svg"
		}
		preview.File = s
		preview.Refresh()
		setApplyStatus(applyButton, iconPath, &selectedDrive, &driveList)
	}

	c := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(header, applyButton, nil, nil),
		header,
		preview,
		applyButton,
	)
	w.SetContent(c)

	w.ShowAndRun()
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
