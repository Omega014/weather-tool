/*
Copyright © 2023 Mitsuki Sugiya <metallic.code@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/omega014/weather-tool/apis"
)

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Print the weather of the specified location",
	Long:  "hoge fuga piyo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vertion 0.0.1")
		cmd.Usage()
		os.Exit(0)
	},
}

var weatherGetCmd = &cobra.Command{
	Use:   "get [Name]",
	Short: "get the weather of the specified location",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city := args[0]

		data := apis.GetWeather(city)
		fmt.Println("get called: ", data)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)
	weatherCmd.AddCommand(weatherGetCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
