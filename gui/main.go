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
	input := widget.NewEntry()
	btn := widget.NewButton("Hi!", func() {
		//fmt.Println(input.Text)
		hello.SetText(input.Text)
	})
	w.SetContent(container.NewVBox(
		hello,
		input,
		btn),
	)

	w.ShowAndRun()
}
