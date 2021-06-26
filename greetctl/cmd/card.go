/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card <name>",
	Short: "This command creates fun greetings",
	Long: `For example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("card called")
		fmt.Println("Here are the arguments of card command : " + strings.Join(args, ","))

		name, _ := cmd.Flags().GetString("name")
		language, _ := cmd.Flags().GetString("language")
		fmt.Println("value of the flag name :" + name)
		fmt.Println("value of the flag language :" + language)
	},
}

func init() {
	createCmd.AddCommand(cardCmd)
	cardCmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear, thanksgiving, birthday")
	cardCmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")
	cardCmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	cardCmd.MarkPersistentFlagRequired("name")
	cardCmd.MarkPersistentFlagRequired("occasion")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var card2Cmd = &cobra.Command{
	Use:   "card <name>",
	Short: "This command creates fun greetings",
	Long: `For example:
	greetctl create card eva -n="Eva Green" -o=thanksgiving -l=fr `,
	Run: func(cmd *cobra.Command, args []string) {
		o, _ := cmd.Flags().GetString("ocassion")
		n, _ := cmd.Flags().GetString("name")
		switch o {
		case "birthday":
			fmt.Println("Happy Birthday !", n)
		case "newyear":
			fmt.Println("Happy newyear !", n)
		default:
			fmt.Println("Happy Thanksgiving")
		}
	},
}

func init() {
	getCmd.AddCommand(card2Cmd)
	card2Cmd.PersistentFlags().StringP("occasion", "o", "", "Possible values: newyear, thanksgiving, birthday")
	card2Cmd.PersistentFlags().StringP("language", "l", "en", "Possible values: en, fr")
	card2Cmd.PersistentFlags().StringP("name", "n", "", "Name of the user to whom you want to greet")
	card2Cmd.MarkPersistentFlagRequired("name")
	card2Cmd.MarkPersistentFlagRequired("occasion")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
