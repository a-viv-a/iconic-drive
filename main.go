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

//returns human readable name for each drive, then the path in a map
func drives() ([]string, map[string]string) {
	driveList, _ := usbdrivedetector.Detect() //shouldnt toss this error
	driveMap := make(map[string]string)
	for i, drive := range driveList {
		driveMap[filepath.Base(drive)] = driveList[i]
		driveList[i] = filepath.Base(drive)
	}
	return driveList, driveMap
}

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
				return
			}
		}
	}
	(*applyButton).Disable()
}
