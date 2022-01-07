/*
Copyright © 2022 Gabor Szabo <gabor@szabgab.com>

*/
package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func PrintBanner() {
	fmt.Println("Welcome!")
	fmt.Println("Answer the questions. Press x if you'd like to exit.")
}

func getKey(wordPair map[string]interface{}) string {
	for key, _ := range wordPair {
		return fmt.Sprint(key)
	}
	os.Exit(1)
	return ""
}

func ParseSkills(course CourseFile, skills []Skill) ([][2]string, [][2]string) {
	//fmt.Println("parse skills")

	wordPairs := [][2]string{}
	phrasePairs := [][2]string{}
	for _, skill := range skills {
		for _, word := range skill.Words {
			wordPairs = append(wordPairs, [2]string{word.Word, word.Translation})
		}
		for _, phrase := range skill.Phrases {
			phrasePairs = append(phrasePairs, [2]string{phrase.Phrase, phrase.Translation})
		}
		//sourceLanguage := course.Course.ForSpeakers.Name
		targetLanguage := course.Course.Language.Name
		for _, wordPair := range skill.Dictionary[targetLanguage] {
			word := getKey(wordPair)
			wordPairs = append(wordPairs, [2]string{word, fmt.Sprint(word)})
		}
	}

	return wordPairs, phrasePairs

}

// Given a map of name => value pairs, randomly return one of the names using the values as weights.
// The values do NOT need to add up to 1 or 100 or anything special.
func selectChallenge(weights map[string]float64) string {
	var weightList []float64
	var actionList []string
	for action, weight := range weights {
		weightList = append(weightList, weight)
		actionList = append(actionList, action)
	}

	sum := 0.0
	for _, num := range weightList {
		sum += num
	}
	//fmt.Printf("sum: %v\n", sum)
	cdf := make([]float64, len(weights))
	partial := 0.0
	for ix, num := range weightList {
		partial += num
		cdf[ix] = partial / sum
	}
	//fmt.Printf("cdf %v\n", cdf)
	//fmt.Printf("actions %v\n", actionList)
	selected := rand.Float64()
	//fmt.Printf("selected %v\n", selected)
	for ix, num := range cdf {
		//fmt.Printf("ix: %v num: %v\n", ix, num)
		if selected < num {
			return actionList[ix]
		}
	}
	//fmt.Println("return default")
	return actionList[len(weights)-1]
}

func runChallenge(cases [][2]string) bool {
	selected := rand.Intn(len(cases))
	input := StringPrompt(fmt.Sprintf("%v:", cases[selected][0]))
	input = strings.Trim(input, "\n")
	if input == "x" {
		fmt.Println("Bye")
		return true
	}
	if input == cases[selected][1] {
		fmt.Println("+")
	} else {
		fmt.Println("-")
		fmt.Println(cases[selected][1])
		input = StringPrompt("try again:")
		input = strings.Trim(input, "\n")
		if input == "x" {
			fmt.Println("Bye")
			return true
		}
		if input == cases[selected][1] {
			fmt.Println("++")
		} else {
			fmt.Println("-")
		}
	}
	fmt.Println("")
	return false
}

func RunSession(fullpath string) {
	course, skills := ReadYamlFiles(fullpath)
	wordPairs, phrasePairs := ParseSkills(course, skills)
	//fmt.Println(wordPairs[0])
	//fmt.Println("----------")
	//os.Exit(0)

	PrintBanner()

	weights := map[string]float64{
		"word-source-to-target":   0.5,
		"word-target-to-source":   0.5,
		"phrase-source-to-target": 0.5,
		"phrase-target-to-source": 0.5,
	}

	for {
		challengeName := selectChallenge(weights)
		if challengeName == "word-source-to-target" {
			if runChallenge(wordPairs) {
				os.Exit(0)
			}
		}
		if challengeName == "phrase-source-to-target" {
			if runChallenge(phrasePairs) {
				os.Exit(0)
			}
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
