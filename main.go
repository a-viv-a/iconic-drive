package main

import (
	"os"

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
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 500))

	iconPath := widget.NewEntry()
	iconPath.SetPlaceHolder("paste or type icon's path")
	iconPath.Validator = func(s string) error {
		_, err := os.Stat(s)
		print(err)
		return err
	}
	clearButton := widget.NewButton("clear", func() { iconPath.SetText("") })
	pathWrapper := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, clearButton), container.NewHScroll(iconPath), clearButton)

	preview := canvas.NewImageFromFile("/home/isaacr/Documents/go shit/bareblue/Icon.png")
	preview.FillMode = canvas.ImageFillContain
	preview.SetMinSize(fyne.NewSize(160, 160))

	iconPath.OnChanged = func(s string) {
		//println(s)
		if !os.IsNotExist(iconPath.Validate()) {
			preview.File = s
			preview.Refresh()
		}
	}

	c := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), pathWrapper, preview)
	w.SetContent(c)

	w.ShowAndRun()

}
