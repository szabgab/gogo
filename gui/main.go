package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func setText(obj widget.Label) {
	obj.SetText("Welcome :)")
}
func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	btn := widget.NewButton("Hi!", func() {
		setText(*hello)
	})
	w.SetContent(container.NewVBox(
		hello,
		btn),
	)

	w.ShowAndRun()
}
