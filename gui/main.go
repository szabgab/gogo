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
var currentCourse string
var allCourses Courses

const coursesURL = "https://github.szabgab.com/lili/courses.json"

func showError(text string) {
	label := widget.NewLabel(text)
	errorView := container.NewVBox(label)
	myWindow.SetContent(errorView)
}

func showSplashScreen() {
	fmt.Println("splash screen")
	label := widget.NewLabel("Welcome. Please wait")
	splashView := container.NewVBox(label)
	myWindow.SetContent(splashView)
}

func showMain() {
	label := widget.NewLabel("Hello")
	btn := widget.NewButton("Open", pressButton)
	mainView := container.NewVBox(
		label,
		btn)

	myWindow.SetContent(mainView)
}

func showCourseSelector() {
	var options []string
	fmt.Println("show course selector")

	label := widget.NewLabel("Select the course")
	selector := widget.NewSelect([]string{}, func(selected string) {})
	btn := widget.NewButton("Start", func() {
		fmt.Println("select button pressed")
		fmt.Println(selector.Selected)
		currentCourse = selector.Selected
		saveSelectedCourse(selector.Selected)
	})
	languageSelectorView := container.NewVBox(
		label,
		selector,
		btn)
	for key, _ := range allCourses {
		fmt.Println(key)
		options = append(options, key)
	}
	selector.Options = options

	myWindow.SetContent(languageSelectorView)
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

func showCurrentCourse() {
	fmt.Println("show current course")
	label := widget.NewLabel("Welcome to the " + currentCourse + " course")
	btn := widget.NewButton("Open", pressButton)
	mainView := container.NewVBox(
		label,
		btn)

	myWindow.SetContent(mainView)
}

func main() {
	var err error

	myApp := app.New()
	myWindow = myApp.NewWindow("GoGo")

	currentCourse, allCourses, err = readConfig()
	if err != nil {
		fmt.Println(err)
	}
	if len(allCourses) == 0 {
		go downloadListOfCourses()
	}

	cases = [][]string{
		[]string{"apple", "banana", "peach"},
		[]string{"dolphin", "ant"},
	}
	cnt = 0

	if currentCourse != "" {
		showCurrentCourse()
	} else if len(allCourses) > 0 {
		showCourseSelector()
	} else {
		showSplashScreen()
	}

	myWindow.ShowAndRun()
}
