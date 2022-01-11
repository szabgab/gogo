package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var myWindow fyne.Window
var cnt int
var cases = [][]string{}

func showMain() {
	label := widget.NewLabel("Hello")
	btn := widget.NewButton("Open", pressButton)
	mainView := container.NewVBox(
		label,
		btn)

	myWindow.SetContent(mainView)
}

func pressButton() {
	fmt.Println("open\n----")
	cnt += 1
	cnt %= 2
	//fmt.Println(cnt)
	//fmt.Println(cases[cnt])
	buttons := []*widget.Button{}
	//buttons := []string{}
	buttonsView := container.NewVBox()
	for _, name := range cases[cnt] {
		//fmt.Println(name)
		copy := name[:]
		btn := widget.NewButton(name, func() { fmt.Println(copy) })
		buttons = append(buttons, btn)
		buttonsView.Add(btn)
		//buttons = append(buttons, name)
	}
	btn := widget.NewButton("Back", showMain)

	buttonsView.Add(btn)
	myWindow.SetContent(buttonsView)
}

func main() {
	myApp := app.New()
	myWindow = myApp.NewWindow("Hello")

	cases = [][]string{
		[]string{"apple", "banana", "peach"},
		[]string{"dolphin", "ant"},
	}
	cnt = 0

	showMain()
	myWindow.ShowAndRun()
}
