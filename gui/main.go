package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	btn := widget.NewButton("Hi!", func() {
		hello.SetText("Welcome :)")
	})
	w.SetContent(container.NewVBox(
		hello,
		btn),
	)

	w.ShowAndRun()
}
