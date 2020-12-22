package main

import (
	"errors"
	"io/ioutil"

	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	usbdrivedetector "github.com/deepakjois/gousbdrivedetector"
	"github.com/h2non/filetype"
)

func main() {
	a := app.New()
	w := a.NewWindow("iconic drive")
	//w.SetFixedSize(true)
	w.Resize(fyne.NewSize(1, 500))

	iconPath := widget.NewEntry()
	iconPath.SetPlaceHolder("paste or type image path")
	iconPath.Validator = func(s string) error {
		buf, _ := ioutil.ReadFile(s) //remember to actually test error
		if filetype.IsImage(buf) {
			return nil
		}
		return errors.New("bad") //this is so bad it hurts but ill fix it later
	}
	clearButton := widget.NewButton("clear", func() { iconPath.SetText("") })
	pathWrapper := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(nil, nil, nil, clearButton),
		container.NewHScroll(iconPath), clearButton,
	)

	driveNameList, driveList := drives()

	driveSelect := widget.NewSelect(driveNameList, func(s string) { println(s, driveList) })
	driveSelect.PlaceHolder = "select target drive"
	refreshButton := widget.NewButton("refresh",
		func() {
			driveNameList, driveList = drives()
			driveSelect.Options = driveNameList
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

	applyButton := widget.NewButton("apply", nil)
	applyButton.Disable()

	iconPath.OnChanged = func(s string) {
		if iconPath.Validate() != nil {
			//https://www.iconfinder.com/icons/381599/error_icon
			s = "error.svg"
			applyButton.Disable()
		} else {
			applyButton.Enable()
		}
		preview.File = s
		preview.Refresh()
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

func drives() ([]string, []string) {
	//returns human readable name for each drive, then the path
	driveList, _ := usbdrivedetector.Detect() //shouldnt toss this error
	driveNameList := make([]string, len(driveList))
	for i, drive := range driveList {
		driveNameList[i] = filepath.Base(drive)
	}
	return driveNameList, driveList
}
