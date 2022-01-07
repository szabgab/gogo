package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

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

func ReadYamlFiles(fullpath string) (CourseFile, []Skill) {
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
	return course, skills
}
