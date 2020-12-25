package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("iconic drive")
	//w.SetFixedSize(true)
	//w.Resize(fyne.NewSize(300, 500))

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

	applyButton := widget.NewButton("apply", func() {
		prog := dialog.NewProgressInfinite("working...", "setting icon...", w)
		//prog.Show()
		applyIcon(iconPath.Text, driveMap[selectedDrive])
		prog.Hide()
		dialog.ShowInformation("all icons written", "remount drive to see changes", w)
	})
	applyButton.Disable()

	driveSelect := widget.NewSelect(driveList, func(s string) {
		//println(s + " <-> " + driveMap[s])
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

	preview := canvas.NewImageFromFile("data/error.svg")
	preview.FillMode = canvas.ImageFillContain
	preview.SetMinSize(fyne.NewSize(300, 300))

	iconPath.OnChanged = func(s string) {
		if iconPath.Validate() != nil {
			//https://www.iconfinder.com/icons/381599/error_icon
			s = "data/error.svg"
		}
		preview.File = s
		preview.Refresh()
		setApplyStatus(applyButton, iconPath, &selectedDrive, &driveList)
	}

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(header, applyButton, nil, nil),
			header,
			preview,
			applyButton,
		))

	w.ShowAndRun()
}
