/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ogabriel/ricer/internal"
	"github.com/ogabriel/ricer/internal/setup"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install setups",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a profile to install")
		} else if setup, found := findSetup(args[0]); found {
			internal.RunSetup(setup)
		} else {
			fmt.Println("Profile not found")
		}
	},
}

func findSetup(arg_setup string) (internal.Setup, bool) {
	for _, s := range setup.Setups {
		if s.Name == arg_setup {
			return s, true
		}
	}

	return internal.Setup{}, false
}

func init() {
	rootCmd.AddCommand(installCmd)
}
