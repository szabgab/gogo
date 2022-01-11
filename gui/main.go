package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	input := widget.NewEntry()
	selector := widget.NewSelect([]string{}, func(name string) { fmt.Printf("Selected %v\n", name) })
	btn := widget.NewButton("Add list", func() {
		//fmt.Println(input.Text)
		selector.Options = []string{"Apple", "Banana", "Peach", "Watermelon"}

		fmt.Println(selector.Selected)
		//hello.SetText(input.Text)
	})
	w.SetContent(container.NewVBox(
		hello,
		selector,
		input,
		btn),
	)

	w.ShowAndRun()
}
