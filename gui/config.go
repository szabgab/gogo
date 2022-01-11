package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Course struct {
	SourceLanguageCode string `json:"source_language_code"`
	SourceLanguageName string `json:"source_language_name"`
	SourcePhrases      int    `json:"source_phrases"`
	SourceWords        int    `json:"source_words"`

	TargetLanguageCode string `json:"target_language_code"`
	TargetLanguageName string `json:"target_language_name"`
	TargetPhrases      int    `json:"target_phrases"`
	TargetWords        int    `json:"target_words"`
}

type Courses map[string]Course

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

func saveSelectedCourse(courseName string) {
	var err error
	var gogoPath string
	var fh *os.File

	gogoPath, err = createAndGetConfigDir()
	if err != nil {
		fmt.Println(err)
		return
	}

	configPath := filepath.Join(gogoPath, "gogo.txt")
	fh, err = os.Create(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fh.WriteString(courseName + "\n")
	fh.Close()
}

// Try to read the local configuration
func readConfig() (string, Courses, error) {
	var gogoPath string
	var courseName string
	var err error
	var courses Courses

	gogoPath, err = createAndGetConfigDir()
	if err != nil {
		fmt.Println(err)
		return courseName, courses, err
	}

	courses, err = readListOfCourses(gogoPath)

	courseName, err = readSelectedCourse(gogoPath)
	if err != nil {
		fmt.Println(err)
		return "", courses, err
	}

	return courseName, courses, nil
}

func readListOfCourses(gogoPath string) (Courses, error) {
	var err error
	var yfile []uint8
	var data Courses

	jsonPath := filepath.Join(gogoPath, "courses.json")
	yfile, err = ioutil.ReadFile(jsonPath)
	//fmt.Printf("%T\n", yfile)
	if err != nil {
		log.Println(err)
		return data, err
	}
	err = json.Unmarshal(yfile, &data)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(data)
	return data, nil
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
