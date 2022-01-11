package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myWindow fyne.Window

func main() {
	myApp := app.New()
	myWindow = myApp.NewWindow("Hello")

	label := widget.NewLabel("Hello")
	btn := widget.NewButton("Open", func() {
		fmt.Println("open")
		go showAnother(myApp)
	})

	myWindow.SetContent(container.NewVBox(
		label,
		btn))

	myWindow.ShowAndRun()
}

func showAnother(a fyne.App) {
	//time.Sleep(time.Second * 5)

	win := a.NewWindow("Shown later")
	//win.SetContent(widget.NewLabel("5 seconds later"))
	win.SetContent(widget.NewButton("Close", func() { win.Close() }))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()

	//time.Sleep(time.Second * 2)
	//win.Close()
}
