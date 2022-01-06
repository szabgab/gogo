/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func PrintBanner() {
	fmt.Println("Welcome!")
	fmt.Println("Answer the questions. Press x if you'd like to exit.")
}

func RunSession() {
	PrintBanner()
	cases := [2][2]string{
		{"book", "livro"},
		{"apple", "manzana"},
	}
	//fmt.Println(len(cases))

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
		RunSession()
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
