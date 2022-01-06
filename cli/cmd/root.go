/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type Course struct {
	Course  map[interface{}]interface{} `yaml:"Course"`
	Modules []string                    `yaml:"Modules"`
}

func PrintBanner() {
	fmt.Println("Welcome!")
	fmt.Println("Answer the questions. Press x if you'd like to exit.")
}

func ReadYamlFiles(fullpath string) [2][2]string {
	course_yaml_file := filepath.Join(fullpath, "course.yaml")
	//fmt.Println(course_yaml_file)
	_, err := os.Stat(course_yaml_file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("course.yaml file could not be found. Have you provided the path to the course directory?")
			os.Exit(1)
		}
		log.Fatal(err)
		os.Exit(1)
	}

	yfile, err2 := ioutil.ReadFile(course_yaml_file)
	if err2 != nil {
		log.Fatal(err2)
		os.Exit(1)
	}
	//data := make(map[interface{}]interface{})
	//data := //make(map[string]Course)
	var data Course
	err3 := yaml.Unmarshal(yfile, &data)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(data.Course)
	//fmt.Println(data.Course.Language)
	fmt.Println(data.Modules)

	// for k, v := range data["Modules"] {
	// 	fmt.Printf("%s -> %d\n", k, v)
	// }
	os.Exit(0)

	cases := [2][2]string{
		{"book", "livro"},
		{"apple", "manzana"},
	}
	//fmt.Println(len(cases))
	return cases
}

func RunSession(fullpath string) {
	cases := ReadYamlFiles(fullpath)
	PrintBanner()

	for {
		selected := rand.Intn(len(cases))
		input := StringPrompt(fmt.Sprintf("%v:", cases[selected][0]))
		input = strings.Trim(input, "\n")
		if input == "x" {
			fmt.Print("Bye")
			return
		}
		if input == cases[selected][1] {
			fmt.Println("+")
		} else {
			fmt.Println("-")
		}
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "Learn a language",
	Long:  `A longer explanation on how to learn a language`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		rand.Seed(time.Now().UnixNano())
		var coursedir = "."
		if len(args) >= 1 && args[0] != "" {
			coursedir = args[0]
		}

		//fmt.Println(coursedir)
		fullpath, err := filepath.Abs(coursedir)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}
		//fmt.Println(fullpath)
		RunSession(fullpath)
	},
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
