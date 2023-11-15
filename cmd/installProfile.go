/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ogabriel/ricer/internal"
	"github.com/ogabriel/ricer/internal/profile"
	"github.com/spf13/cobra"
)

var installProfileCmd = &cobra.Command{
	Use:   "install-profile",
	Short: "Install any profile",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "hyprland" {
			internal.RunProfile(profile.Hyprland)
		} else if args[0] == "hyprland-git" {
			internal.RunProfile(profile.HyprlandGit)
		}
	},
}

func init() {
	rootCmd.AddCommand(installProfileCmd)
}
