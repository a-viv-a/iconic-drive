package main

import (
	"errors"
	"io/ioutil"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/h2non/filetype"
)

func main() {
	a := app.New()
	w := a.NewWindow("iconic drive")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(300, 1))

	iconPath := widget.NewEntry()
	iconPath.SetPlaceHolder("paste or type image path")
	iconPath.Validator = func(s string) error {
		buf, _ := ioutil.ReadFile(s)
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

	preview := canvas.NewImageFromFile("error.svg")
	preview.FillMode = canvas.ImageFillContain
	preview.SetMinSize(fyne.NewSize(256, 256))

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

	c := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		pathWrapper,
		widget.NewSeparator(),
		preview,
		widget.NewSeparator(),
		applyButton,
	)
	w.SetContent(c)

	w.ShowAndRun()

}
