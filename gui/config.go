package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func createAndGetConfigDir() (string, error) {
	var err error
	var dirname string

	dirname, err = os.UserHomeDir()
	if err != nil {
		log.Println("Error: %v", err)
		return "", err
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
				return "", err
			}
		} else {
			fmt.Printf("Error: %s\n", err)
			return "", err
		}
	}
	return gogoPath, nil
}

func readSelectedCourse(gogoPath string) (string, error) {
	var err error
	var fh *os.File
	var line string

	configPath := filepath.Join(gogoPath, "gogo.txt")
	fh, err = os.Open(configPath)
	if err != nil {
		fmt.Printf("Could not open file '%v': %v", configPath, err)
		return "", err
	}
	//fmt.Printf("%T\n", fh)
	reader := bufio.NewReader(fh)
	line, err = reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Could not read from file '%v': %v", configPath, err)
		return "", err
	}
	//fmt.Println("end readConfig")
	return line, nil
}

// Try to read the local configuration
func readConfig() (string, []string, error) {
	var gogoPath string
	var courseName string
	var err error
	var courses []string

	gogoPath, err = createAndGetConfigDir()
	if err != nil {
		fmt.Println(err)
		return courseName, courses, err
	}

	courseName, err = readSelectedCourse(gogoPath)
	if err != nil {
		fmt.Println(err)
		return "", courses, err
	}

	//readCourses(gogoPath)
	//getCourses()
	return courseName, courses, nil
}

func saveListOfCourses(courses []byte) {
	var err error
	var gogoPath string

	gogoPath, err = createAndGetConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	jsonPath := filepath.Join(gogoPath, "courses.json")
	err = ioutil.WriteFile(jsonPath, courses, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
