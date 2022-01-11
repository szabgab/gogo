package main

import (
	"io"
	"log"
	"net/http"
)

func downloadListOfCourses() {
	//resource, err := fyne.LoadResourceFromURLString(coursesURL)
	var err error
	var resp *http.Response
	var body []byte

	resp, err = http.Get(coursesURL)
	if err != nil {
		log.Println("Error: %v", err)
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: %v", err)
	}
	//fmt.Println(body)

	if err != nil {
		log.Println("Error: %v", err)
	}
	saveListOfCourses(body)
}
