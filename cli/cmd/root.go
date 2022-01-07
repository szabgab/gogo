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

type Language struct {
	Name string `yaml:"Name"`
	Code string `yaml:"IETF BCP 47"`
}

type License struct {
	Name      string `yaml:"Name"`
	ShortName string `yaml:"Short name"`
	Link      string `yaml:"Link"`
}

type CourseData struct {
	Language    Language `yaml:"Language"`        // TargetLanguage
	ForSpeakers Language `yaml:"For speakers of"` // SourceLanguage
	License     License  `yaml:"License"`
	Repository  string   `yaml:"Repository"`
	Characters  []string `yaml:"Special characters"`
}

type CourseFile struct {
	Course  CourseData `yaml:"Course"`
	Modules []string   `yaml:"Modules"`
}

func PrintBanner() {
	fmt.Println("Welcome!")
	fmt.Println("Answer the questions. Press x if you'd like to exit.")
}

func ReadCourseYamlFile(fullpath string) CourseFile {
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
	var data CourseFile
	err3 := yaml.Unmarshal(yfile, &data)
	if err3 != nil {
		log.Fatal(err3)
	}
	//fmt.Println(data.Course.License.Name)
	//fmt.Println(data.Course.Language)
	//fmt.Println(data.Course.Language.Name)
	//fmt.Println(data.Course.ForSpeakers.Name)
	return data
}

type ModuleFile struct {
	Module Module   `yaml:"Module"`
	Skills []string `yaml:"Skills"`
}

type Module struct {
	Name string `yaml:"Name"`
}

func ReadModuleYamlFile(fullpath string, name string) ModuleFile {
	module_yaml_file := filepath.Join(fullpath, name, "module.yaml")
	yfile, err := ioutil.ReadFile(module_yaml_file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var data ModuleFile
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)
	return data
}

type SkillMeta struct {
	Name       string   `yaml:"Name"`
	Id         int      `yaml:"Id"`
	Thumbnails []string `yaml:"Thumbnails"`
}

type Word struct {
	Word        string   `yaml:"Word"`
	Translation string   `yaml:"Translation"`
	Images      []string `yaml:"Images"`
}

type Skill struct {
	Meta  SkillMeta `yaml:"Skill"`
	Words []Word    `yaml:"New words"`
}

func ReadSkillYamlFile(fullpath string, module_name string, skill_file string) Skill {
	skill_yaml_file := filepath.Join(fullpath, module_name, "skills", skill_file)
	//fmt.Println(skill_yaml_file)
	yfile, err := ioutil.ReadFile(skill_yaml_file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	var data Skill
	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func ParseSkills(skills []Skill) [][2]string {
	fmt.Println("parse skills")
	wordPairs := [][2]string{}
	//fmt.Println(len(skills))
	//fmt.Println(skills[0])
	for _, skill := range skills {
		for _, word := range skill.Words {
			//fmt.Println(word.Word)
			//fmt.Println(word.Translation)
			wordPairs = append(wordPairs, [2]string{word.Word, word.Translation})
		}
	}

	// wordPairs := [][2]string{
	// 	{"book", "livro"},
	// 	{"apple", "manzana"},
	// }
	return wordPairs

}

// TODO finish reading the skills
// TODO: tech the first skill
// TODO also read the md files if they exist

func ReadYamlFiles(fullpath string) [][2]string {
	course := ReadCourseYamlFile(fullpath)
	skills := []Skill{}
	//words := [][2]string{}
	//fmt.Println(course.Modules)
	for _, module_name := range course.Modules {
		//fmt.Printf("name: %s\n", module_name)
		module := ReadModuleYamlFile(fullpath, module_name)
		//fmt.Println(module.Module.Name)
		for _, skill_name := range module.Skills {
			skill := ReadSkillYamlFile(fullpath, module_name, skill_name)
			//fmt.Println(skill.Meta.Name)
			skills = append(skills, skill)
		}
	}
	wordPairs := ParseSkills(skills)
	//fmt.Println(wordPairs[0])
	//fmt.Println("----------")
	//os.Exit(0)

	//fmt.Println(len(cases))
	return wordPairs
}

func RunSession(fullpath string) {
	cases := ReadYamlFiles(fullpath)
	PrintBanner()
	fmt.Println(len(cases))
	os.Exit(0)

	for {
		selected := rand.Intn(len(cases))
		input := StringPrompt(fmt.Sprintf("%v:", cases[selected][0]))
		input = strings.Trim(input, "\n")
		if input == "x" {
			fmt.Println("Bye")
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
