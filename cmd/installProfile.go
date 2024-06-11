/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ogabriel/ricer/internal"
	"github.com/ogabriel/ricer/internal/profile"
	"github.com/spf13/cobra"
)

var installProfileCmd = &cobra.Command{
	Use:   "install profile",
	Short: "Install any profile",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please specify a profile to install")
		} else if args[0] == "hyprland" {
			internal.RunProfile(profile.Hyprland)
		} else if args[0] == "hyprland-git" {
			internal.RunProfile(profile.HyprlandGit)
		} else {
			fmt.Println("Profile not found")
		}
	},
}

func init() {
	rootCmd.AddCommand(installProfileCmd)
}
