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
	"pluralith/helpers"
	"pluralith/ux"

	"github.com/spf13/cobra"
)

// Defining command args/flags
var pluralithApplyArgs = []string{}

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Run terraform apply and draw diagram",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Manually parsing arg (due to cobra lacking a feature)
		parsedArgs, _ := helpers.ParseArgs(args, pluralithApplyArgs)
		parsedArgs = append(parsedArgs, "-auto-approve")

		var confirm string
		ux.PrintFormatted("?", []string{"blue", "bold"})
		fmt.Println(" Apply Plan?")
		ux.PrintFormatted("  Yes to confirm: ", []string{"bold"})
		fmt.Scanln(&confirm)

		if confirm == "yes" {
			ux.PrintFormatted("\n✔", []string{"blue", "bold"})
			fmt.Println(" Apply Confirmed")

			// Launching Pluralith
			helpers.LaunchPluralith()

			ux.PrintFormatted("⣿", []string{"blue", "bold"})
			fmt.Println(" Status:")

			if _, code := helpers.ExecuteTerraform("apply", parsedArgs, false, true, true); code == 0 {

				ux.PrintFormatted("✔ All Done!\n", []string{"blue", "bold"})
			}
		} else {
			ux.PrintFormatted("\n✖️", []string{"red", "bold"})
			fmt.Println(" Apply Aborted")
		}
		// helpers.ExecuteTerraform("apply", parsedArgs, false, false)

	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
}
