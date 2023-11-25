/*
Copyright © 2023 NAME HERE <tuancnttbk93@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "translate",
	Short: "English to English dictionary CLI",
	Long:  `Dictionary cli is a simple tool that allows you to quickly translate a keyword in terminals without switch contexts.`,
	Run: func(cmd *cobra.Command, args []string) {
		type Definition struct {
			Fl       string   `json:"fl"`
			Shortdef []string `json:"shortdef"`
		}

		type ApiResponse struct {
			Definitions []Definition `json:"definition"`
		}

		var query = "hello"

		if len(args) >= 1 && args[0] != "" {
			query = args[0]
		}

		fmt.Println("Translating for key '" + query + "' ...")

		URL := "https://3jf1ce6vn9.execute-api.ap-southeast-1.amazonaws.com/default/DictionaryAPI?q=" + query

		response, err := http.Get(URL)

		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Read the response body
			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Println("Error reading the response body:", err)
				return
			}

			// Create an instance of the ApiResponse struct
			var apiResponse ApiResponse

			// Unmarshal the JSON data into the struct
			err = json.Unmarshal(body, &apiResponse)
			if err != nil {
				fmt.Println("Error unmarshaling JSON:", err)
				return
			}

			for index, definition := range apiResponse.Definitions {
				if len(definition.Shortdef) == 0 {
					continue
				}
				seq := index + 1
				fmt.Printf("%d. (%s):\n", seq, definition.Fl)
				for _, def := range definition.Shortdef {
					fmt.Printf(" - %s\n", def)
				}
			}
		} else {
			fmt.Println("Something went wrong")
		}
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dictionary-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
