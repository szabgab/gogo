package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func createAndGetConfigDir() string {
	var err error
	var dirname string

	dirname, err = os.UserHomeDir()
	if err != nil {
		log.Println("Error: %v", err)
	}
	gogoPath := filepath.Join(dirname, "gogo")
	_, err = os.Stat(gogoPath)
	if err != nil {
		//fmt.Printf("Error: %s\n", err)
		if os.IsNotExist(err) {
			fmt.Printf("Creating %s\n", gogoPath)
			err = os.Mkdir(gogoPath, 0755)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		} else {
			fmt.Printf("Error: %s\n", err)
		}
	}
	return gogoPath
}

func readSelectedCourse(gogoPath string) string {
	var err error
	var fh *os.File
	line := ""

	configPath := filepath.Join(gogoPath, "gogo.txt")
	fh, err = os.Open(configPath)
	if err != nil {
		fmt.Printf("Could not open file '%v': %v", configPath, err)
		return line
	}
	//fmt.Printf("%T\n", fh)
	reader := bufio.NewReader(fh)
	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Could not read from file '%v': %v", configPath, err)
		return line
	}
	//fmt.Println("end readConfig")
	return line
}

// Try to read the local configuration
func readConfig() string {
	gogoPath := createAndGetConfigDir()
	courseName := readSelectedCourse(gogoPath)

	//getCourses()
	return courseName
}
