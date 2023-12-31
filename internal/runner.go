package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func run(command string) {
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func RunCommand(command string) {
	run(command)
}

func RunCommands(commands []string) {
	command := strings.Join(commands, " && ")
	run(command)
}

func RunPackages(packages []string) {
	packages = append([]string{"yay -Syu --needed --noconfirm"}, packages...)
	command := strings.Join(packages, " ")
	RunCommand(command)
}

func RunSetup(setup Setup) {
	if setup.Before != nil {
		fmt.Println(setup.Name)
		setup.Before()
	}

	if setup.Packages != nil {
		fmt.Println(setup.Name)
		RunPackages(setup.Packages)
	}

	if setup.After != nil {
		fmt.Println(setup.Name)
		setup.After()
	}
}

func RunSetups(setups []Setup) {
	for _, setup := range setups {
		if setup.Before != nil {
			fmt.Println("Running before for", setup.Name)
			setup.Before()
		}
	}

	fmt.Println("Installing packages...")
	var packages []string
	for _, setup := range setups {
		if setup.Packages != nil {
			packages = append(packages, setup.Packages...)
		}
	}

	if packages != nil {
		RunPackages(packages)
	}

	for _, setup := range setups {
		if setup.After != nil {
			fmt.Println("Running after for", setup.Name)
			setup.After()
		}
	}

}

func RunProfile(profile Profile) {
	fmt.Println("Running profile", profile.Name)
	if profile.Setups != nil {
		RunSetups(profile.Setups)
	}
}
