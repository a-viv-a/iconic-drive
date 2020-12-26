package main

import (
	"log"
	"runtime"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	log.Printf("GOARCH:%s GOOS:%s\n", runtime.GOARCH, runtime.GOOS)
	a := app.New()
	w := a.NewWindow("iconic drive")
	//w.SetFixedSize(true)
	//w.Resize(fyne.NewSize(300, 500))

	iconPath := widget.NewEntry()
	iconPath.SetPlaceHolder("paste or type image path")
	iconPath.Validator = testImgPath
	clearButton := widget.NewButtonWithIcon("", theme.DeleteIcon(), func() { iconPath.SetText("") })
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
		log.Println(s + " selected from dropdown")
		selectedDrive = s
		setApplyStatus(applyButton, iconPath, &selectedDrive, &driveList)
	})

	driveSelect.PlaceHolder = "select target drive"
	refreshButton := widget.NewButtonWithIcon("",
		theme.ViewRefreshIcon(),
		func() {
			log.Println("refreshed drive selection")
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

	errRes := theme.QuestionIcon() //fyne.NewStaticResource("error", MustAsset("data/error.png"))
	preview := canvas.NewImageFromResource(errRes)
	preview.FillMode = canvas.ImageFillContain
	preview.SetMinSize(fyne.NewSize(300, 300))

	iconPath.OnChanged = func(s string) {
		preview.Resource = nil
		preview.File = ""
		if iconPath.Validate() != nil {
			//https://www.iconfinder.com/icons/381599/error_icon
			preview.Resource = errRes
		} else {
			preview.File = s
		}
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

	log.Println("GUI shown and run")
	w.ShowAndRun()
	log.Println("GUI closed/crashed")

}
